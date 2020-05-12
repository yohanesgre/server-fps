package main

import (
	"./game"
	network "./network"
	"log"
	"time"
)

func main() {
	host := ":1887"
	hostUdp := ":1878"

	ss, err := network.NewSocketService(host, hostUdp)
	if err != nil {
		log.Println(er)
		return
	}

	ss.SetHeartBeat(5*time.Second, 30*time.Second)
	ss.RegMessageHandler(game.HandleMessage)
	ss.RegUDPMessageHandler(game.HandleMessaeUDP)
	ss.RegConnectHandler(game.HandleConnect)
	ss.RegDisconnectHandler(game.HandleDisconect)

	log.Println("server running on " + host + " and " + hostUdp)
	ss.Serv()
}
