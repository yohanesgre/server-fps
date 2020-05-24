package state

import ("fmt")

type ServerState string

const(
	Idle = "Idle"
	Loading = "Loading"
	Active = "Active"
)

type StateMachine interface{
	Add()
	CurrentState()
	Update()
	Shutdown()
	SwitchTo()
}

type State interface{
	id		ServerState
	
}