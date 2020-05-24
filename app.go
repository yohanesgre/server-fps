package main

import (
	"github.com/yohanesgre/server-fps/game"
	network "github.com/yohanesgre/server-fps/network"
	"log"
	"time"
)

func main() {
	host := ":1887"
	hostUdp := ":1878"

	ss, err := network.NewSocketService(host, hostUdp)
	if err != nil {
		log.Println(err)
		return
	}

	ss.SetHeartBeat(5*time.Second, 30*time.Second)
	ss.RegMessageHandler(game.HandleMessage)
	ss.RegUDPMessageHandler(game.HandleMessageUDP)
	ss.RegConnectHandler(game.HandleConnect)
	ss.RegDisconnectHandler(game.HandleDisconnect)

	log.Println("server running on " + host + " and " + hostUdp)
	ss.Serv()
}
