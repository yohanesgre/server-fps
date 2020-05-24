package core

type PacketType byte

const(
	ACK 			Endpoint = 0
	SNAPSHOT 		Endpoint = 1
	EVENT 			Endpoint = 2
	FLOODING 		Endpoint = 3
)