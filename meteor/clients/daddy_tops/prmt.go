package main

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"os"
	"strings"
)

var suggestions = []prompt.Suggest{
	{"action", "Queue a new action - `action <targetHost> <mode_code> <args>`"},
	{"gaction", "Queue a new group action - `action <targetGroup> <mode> <args>`"},
	{"result", "Show the result for an action - `result <uuid>`"},
	{"list", "List the available <x> - `list <modes/hosts/host/groups/group/bots/actions> <OPT:host/group>`"},
	{"help", "Show the help prompt"},
	{"exit", "Exit prompt"},
}

func executor(in string) {
	fmt.Printf("You said: %s\n", in)
	splitArgs := strings.SplitN(in, " ", 4)
	keywrd := splitArgs[0]
	switch keywrd {
	case "action":
		handleActionKW(splitArgs)
	case "gaction":
		handleGroupActionKW(splitArgs)
	case "result":
		handleResultKW(splitArgs)
	case "list":
		handleListKW(splitArgs)
	case "help":
		handleHelpKW(splitArgs)
	case "exit":
		handleExitKW(splitArgs)
	}
	return
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

func prm() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("Daddy_Tops: Interactive Commander for Meteor C2"),
		prompt.OptionPrefix("DT>"),
	)
	p.Run()
}

func handleActionKW(splitargs []string) {
	return
}

func handleGroupActionKW(splitargs []string) {
	return
}

func handleResultKW(splitargs []string) {
	return
}

func handleListKW(splitargs []string) {
	return
}

func handleHelpKW(splitargs []string) {
	fmt.Printf(`Daddy_Tops CLI: Interactive Commander for Meteor C2
	
Current Server Config:
Server: %s
User: %s
	
CAPABILITY				SYNTAX
------------------------------------------------------------------------------------------

NEW ACTION:				action <target_hostname> <mode_code> <arguments>
NEW GROUP ACTION:		gaction <target_groupname> <mode_code> <arguments>
SHOW RESULT:			result <uuid>
LIST AVAILABLE <X>:		list <modes/hosts/host/groups/group/bots/actions> <OPT:host/group>
HELP MENU				help
QUIT PROMPT				exit
`, DTSERVER, DTUSER)

	return
}

func handleExitKW(splitargs []string) {
	fmt.Printf("Goodbye %s!\n", DTUSER)
	os.Exit(0)
	return
}
