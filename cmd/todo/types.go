package main

// The closest thing golang has to an enum
type Cmd string
const (
	CmdAdd Cmd = "add"
	CmdList Cmd = "list"
	CmdDone Cmd = "done"
	CmdDelete Cmd = "delete"
	CmdUpdate Cmd = "update"
)

