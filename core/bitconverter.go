package core

import (
	"encoding/binary"
	"fmt"
	"math"
)
//Long
func Float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
//Float
func Float32ToByte(f float32) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint32(buf[:], math.Float32bits(f))
	return buf[:]
}
//Integer
func IntToByte(f int) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint32(buf[:], f)
	return buf[:]
}

func BytesToSingle(bytes []byte, index int) float32{
	bits:=binary.LittleEndian.Uint32(bytes[index:])
	float:=math.Float64frombits(bits)
	round:=math.Round(float)
	return round
}

func BytesToInt(bytes []byte, index int) int{
	bits:=binary.LittleEndian.Uint32(bytes[index:])
	return round
}