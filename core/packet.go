package packet

import(
	"bytes"
)

type Packet struct{
	payload		[]byte
	position	int
}

func InitPacket(payload []byte) *Packet{
	var pa *Packet
	pa = new(Packet)
	pa.payload = payload
	pa.position = 0
	return pa
}

func InitPacket(builder *Builder) *Packet{
	var pa *Packet
	pa = new(Packet)
	pa.position = 0
	if builder.position == len(builder.payload){
		pa.payload = builder.payload
	}else{
		pa.payload = make([]byte, builder.position)
		copy(pa.payload, builder.payload[:builder.position])
	}
	return pa
}

func (pa *Packet) GetPayload() []byte{
	return pa.payload
}

func (pa *Packet) GetByte() []byte{
	return pa.payload[pa.position++]
}

func (pa *Packet) GetDirection() Direction{
	return GetByte()
}

func (pa *Packet) GetEndpoint() Endpoint{
	return GetByte()
}

func (pa *Packet) GetFloat() float32{
	data:= BytesToSingle(pa.payload, pa.position)
	pa.position += 4
	return data
}

func (pa *Packet) GetInt() int{
	data:= BytesToInt(pa.payload, pa.position)
	pa.position += 4
	return data
}

func (pa *Packet) GetPacketType() PacketType{
	return GetByte()
}

func (pa *Packet) GetString() string{
	length:=GetByte()
	data:= string(pa.payload[pa.position:length])
	pa.position += length
	return data
}

func (pa *Packet) GetBitBuffer() BitBuffer{
	length:=GetByte()
	buffer:=make([]byte, length)
	for i := 0; i < length; ++i {
		buffer[i] = GetByte()
	}
	bitbuffer:= InitBitBuffer(buffer)
	return bitbuffer
}

func (pa *Packet) Reset(position int) *Packet{
	pa.position = position
	return pa
}

func (pa *Packet) Reset() *Packet{
	return Reset(0)
}