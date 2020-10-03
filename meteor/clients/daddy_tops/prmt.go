package main

import (
	"github.com/abiosoft/ishell"
)

func prm() {
	shell := ishell.New()
	shell.SetHomeHistoryPath(".dt_history")

	shell.AddCmd(&ishell.Cmd{
		Name: "action",
		Func: func(c *ishell.Context) {
			retval := handleActionKW(c.Args)
			c.Println(retval)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "gaction",
		Func: func(c *ishell.Context) {
			retval := handleGroupActionKW(c.Args)
			c.Println(retval)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "result",
		Func: func(c *ishell.Context) {
			retval := handleResultKW(c.Args)
			c.Println(retval)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "list",
		Func: func(c *ishell.Context) {
			retval := handleListKW(c.Args)
			c.Println(retval)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "help",
		Func: func(c *ishell.Context) {
			retval := handleHelpKW()
			c.Println(retval)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "exit",
		Func: func(c *ishell.Context) {
			retval := handleExitKW()
			c.Println(retval)
			c.Stop()
		},
	})

	shell.Run()
}
