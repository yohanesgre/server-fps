package core

type Direction byte

const(
	FORWARD 			Direction = 0
	BACKWARD 			Direction = 1
	STRAFING_LEFT 		Direction = 2
	STRAFING_RIGHT 		Direction = 3
	ROTATE_LEFT 		Direction = 4
	ROTATE_RIGHT 		Direction = 5
)