package network

import (
	"C"
	"net"
	"encoding/binary"
	"bytes"
)

type Data struct{
	CurrentPlayers 	uint16
	MaxPlayers 		uint16
	ServerName 		string
	GameType		string
	BuildId			string
	Map				string
	Port			uint16
}

// func (d *Data) ToStream(c *net.UDPConn)  {
// 	buf := new(bytes.Buffer)
// 	WriteNetworkByteOrder(d.CurrentPlayers, buf)
// 	WriteNetworkByteOrder(d.MaxPlayers, buf)
// 	WriteString(d.ServerName, buf)
// 	WriteString(d.GameType, buf)
// 	WriteString(d.BuildId, buf)
// 	WriteString(d.Map, buf)
// 	WriteNetworkByteOrder(d.Port, buf)
// 	c.Write(buf)
// }

// func WriteNetworkByteOrder(value uint16, writer *bytes.Buffer)  {
// 	buf := new(bytes.Buffer)
// 	err := binary.Write(buf, binary.LittleEndian, value)
// 	if err != nil {
// 		fmt.Println("binary.Write failed:", err)
// 	}
// 	writer.Write(buf.Bytes())
// }

// func WriteString(value string, writer *bytes.Buffer){
// 	buf := new(bytes.Buffer)
// 	err := binary.Write(buf, binary.LittleEndian, value)
// 	if err != nil {
// 		fmt.Println("binary.Write failed:", err)
// 	}
// 	writer.Write(buf.Bytes())
// }

