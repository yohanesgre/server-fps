package core

type Endpoint byte

const(
	JOIN 			Endpoint = 0
	MOVE 			Endpoint = 1
	SHOOT 			Endpoint = 2
	FRAG 			Endpoint = 3
	RESPAWN 		Endpoint = 4
)