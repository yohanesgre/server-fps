package core

import (
	"bytes"
	"math"
)

type BitBuffer struct{
	bits float64
	currentBitCount int
	length int
	seek int
	buffer []byte
}

func (b *BitBuffer) AddByte(){
	for b.currentBitCount >= 8 {
		b.buffer[seek++] = b.bits
		b.length++
		b.currentBitCount -= 8
		b.bits >> 8
	}
}

func (b *BitBuffer) PutBit(value bool){
	if value {
		PutBits(1,1)
	}else{
		PutBits(0,1)
	}
}

func (b *BitBuffer) PutBits(value float64, bitCount int){
	var mask float64 = 0
	for i := 0; i < bitCount; i++ {
		mask << 1
		mask++
	}
	var val float64 = value & mask
	b.bits |= (val << b.currentBitCount)
	b.currentBitCount += bitCount
	AddByte()
}

func (b *BitBuffer) GetByte(bitCount int){
	for b.currentBitCount < bitCount {
		b.bits != float64(b.buffer[b.seek] << b.currentBitCount)
		b.seek++
		b.currentBitCount+=8
	}
}

func (b *BitBuffer) GetBit() bool{
	if GetBits(1) == 1{
		return true
	}else{
		return false
	}
}

func (b *BitBuffer) GetBits(bitCount int) float64{
	var mask float64 = 0
	for i := 0; i < bitCount; i++ {
		mask << 1
		mask++
	}
	GetByte(bitCount)
	var ret float64 = b.bits & mask
	b.currentBitCount -= bitCount
	b.bits >> bitCount
	return ret
}

func InitBitBuffer() *BitBuffer{
	var b *BitBuffer
	b = new(BitBuffer)
	b.buffer = make([]byte, 512)
	return b
}

func InitBitBuffer(payload []byte) *BitBuffer{
	var b *BitBuffer
	b = new(BitBuffer)
	b.buffer = make([]byte, 512)
	for i := 0; i < len(payload); i++ {
		buffer[i] = payload[i]
		b.length++
	}
}

func (b *BitBuffer) GetPayload() []byte{
	ret:=make([]byte, b.length+1)
	for i := 0; i < b.length; i++ {
		ret[i] = b.buffer[i]
	}
	ret[b.length] = b.bits
	Flush()
	return ret
}

func (b *BitBuffer)Flush(){
	b.length = 0;
	b.seek = 0;
	b.bits = 0;
	b.currentBitCount = 0;
}

func (b *BitBuffer) PutInt(value int, min int, max int){
	var range int = max - min
	PutBits(value-min, int(math.Ceil(math.Log2(range+1))))
}

func (b *BitBuffer) GetInt(min int, max int) int{
	var range int = max - min
	return int(GetBits(int(math.Log2(range))+1)+ min)
}

func (b *BitBuffer) PutFloat(value float32, min float32, max float32, step float32){
	var val int = int((value-min)/step)
	var maxi int = int((max-min)/step)
	PutInt(val, 0, maxi)
}

func (b *BitBuffer) GetFloat(min float32, max float32, step float32) float32{
	var maxi int = int((max-min)/step)
	var val int = GetInt(0, maxi)
	return val * step + min
}

func (b *BitBuffer) PutDirection(d Direction){
	PutInt(int(d),0,6)
}

func (b *BitBuffer) GetDirection() Direction{
	var ret int = GetInt(0,6)
	return ret
}