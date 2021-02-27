package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"time"

	mcs "github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// icmpConst is usually one, used for protocol marshalling (ipv6 support has a diff one)
const icmpConst int = 1

// timeout is how long the conn read will wait for before timing
const timeout int = 10

// readSize is the amount to read from the listener (always send less than this)
const readSize int = 1600

// calculate the sha1 sum of the data and use (almost) the last 8 chars as the checksum
func calcChecksum(data []byte) []byte {
	hasher := sha1.New()
	hasher.Write(data)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	chk := sha[len(sha)-9 : len(sha)-1]
	return []byte(chk)
}

// generate an icmp listener on the specified IP (usually 0.0.0.0)
func getListener(ip string) *icmp.PacketConn {
	conn, err := icmp.ListenPacket("ip4:icmp", ip)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// take in the required fields for an `IFTPP` message and proto-fy the data
func buildProto(sid int32, payld []byte, chksum []byte, typ mcs.CTP_Flag) []byte {
	testPro := &mcs.CTP{
		SessionId: sid,
		Payload:   payld,
		Checksum:  chksum,
		TypeFlag:  typ,
	}

	data, err := proto.Marshal(testPro)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	return data
}

// marshal the data into an `IFTPP` msg
func decodeProto(data []byte) *mcs.CTP {
	iftppMesage := &mcs.CTP{}
	if err := proto.Unmarshal(data, iftppMesage); err != nil {
		fmt.Println("Failed to unmarshal IFTPP:", err)
		return nil
	}
	return iftppMesage
}

// take a byte slice payload an turn it into an ICMP message packet
func buildICMP(payload []byte) []byte {
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.RawBody{
			Data: payload,
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// disassemble ICMP message packet and extract body
func disasICMP(msg []byte, n int) []byte {
	parsed, err := icmp.ParseMessage(icmpConst, msg[:n])
	if err != nil {
		return nil
	}
	bod, err := parsed.Body.Marshal(icmpConst)
	if err != nil {
		return nil
	}

	return bod
}

// take the required fields for an `IFTPP` message and put them into an ICMP packet
func buildPacket(sid int32, payld []byte, chksum []byte, typ mcs.CTP_Flag) []byte {
	buf := buildProto(sid, payld, chksum, typ)
	packet := buildICMP(buf)
	return packet
}

// disassemble ICMP packet, return an `IFTPP` message so we can use field references
func disasPacket(msg []byte, n int) *mcs.CTP {
	packetBody := disasICMP(msg, n)
	protoMsg := decodeProto(packetBody)
	return protoMsg
}

// read a packet from the listener, extract body from the ICMP, unmarshall the body into an `IFTPP` message
func readFromListener(conn *icmp.PacketConn) (*mcs.CTP, net.Addr) {
	msg := make([]byte, readSize)
	err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return nil, nil
	}
	n, peer, err := conn.ReadFrom(msg)
	if err != nil {
		return nil, nil
	}

	protoMsg := disasPacket(msg, n)
	if protoMsg == nil {
		return nil, nil
	}

	return protoMsg, peer
}

// marshal provided fields into an `IFTPP` message, put that into an ICMP packet body, write the packet to the conn
func writeToListener(conn *icmp.PacketConn, dst net.Addr, sid int32, payld []byte, chksum []byte, typ mcs.CTP_Flag) {
	packet := buildPacket(sid, payld, chksum, typ)
	_, err := conn.WriteTo(packet, dst)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func compareChks(myChk []byte, chk []byte) bool {
	if len(myChk) != len(chk) {
		return false
	}

	for i, j := range myChk {
		if j != chk[i] {
			return false
		}
	}
	return true
}
