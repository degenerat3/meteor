package utils

import (
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	goUUID "github.com/google/uuid"
)

func genRegister(interval int, delta int, regFile string, obfText string) []byte {
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

// GenRegister is the exported version of the registration payload generator
func GenRegister(interval int, delta int, regFile string, obfText string) []byte {
	return genRegister(interval, delta, regFile, obfText)
}

func genCheckin(regFile string, obfText string) []byte {
	uuid := fetchUUID(regFile, obfText)
	pro := &mcs.MCS{
		Uuid: uuid,
		Mode: "checkin",
	}
	data, _ := proto.Marshal(pro)
	return data
}

// GenCheckin is the exported version of the checkIn payload generator
func GenCheckin(regFile string, obfText string) []byte {
	return genCheckin(regFile, obfText)
}

func addResult(uuid string, result string) []byte {
	pro := &mcs.MCS{
		Uuid:   uuid,
		Result: result,
		Mode:   "addResult",
	}
	data, _ := proto.Marshal(pro)
	return data
}
