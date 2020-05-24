package core

import(
	"bytes"
)

type Builder struct{
	payload		[]byte
	position	int
}

func InitBuilder(maxPacketSize int) *Builder{
	var bu *Builder
	bu = new(Builder)
	bu.payload = make([]byte, maxPacketSize)
	bu.position = 0
	return bu
}

func (bu *Builder) AddByte(data byte) *Builder{
	bu.payload[bu.position++] = data
	return bu
}

func (bu *Builder) AddPayload(data byte[]) *Builder{
	return AddPayload(data, 0, data.Length);
}

func (bu *Builder) AddPayload(data byte[], index int, size int) *Builder{
	copy(bu.payload[position], data[index:size])
	bu.position+=size
	return bu
}

func (bu *Builder) AddDirection(type_ Direction) *Builder{
	return AddByte(type_)
}

func (bu *Builder) AddEndPoint(type_ Endpoint) *Builder{
	return AddByte(type_)
}

func (bu *Builder) AddFloat(data float32) *Builder{
	return AddPayload(Float32ToByte(data))
}

func (bu *Builder) AddInteger(data int) *Builder{
	return AddPayload(IntToByte(data))
}

func (bu *Builder) AddPacketType(type_ PacketType)  *Builder {
	return AddByte(type)
}

func (bu *Builder) AddString(data string)  *Builder{
	var stringPayload []byte(data)
	AddByte(len(stringPayload));
	return AddPayload(stringPayload);
}

func (bu *Builder) AddBitBuffer(bb *BitBuffer) *Builder{
	//AddPayload(BitConverter.GetBytes(bb.getData()));
	//AddPayload(BitConverter.GetBytes(bb.getCount()));
	var payload []byte = bb.GetPayload();
	/* foreach (byte b in payload) {
		Debug.Log("Manda: " + Convert.ToString(b,2));
	}*/
	AddByte(payload.Length);
	return AddPayload(payload);
}

func (bu *Builder) Build() *Builder{
	
	return pa
}