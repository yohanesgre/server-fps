package game

import (
	"encoding/json"
	"log"
	"net"

	serverplugin "github.comyohanesgre/go-game-server-module"
)

var a = []*net.UDPAddr{}

// HandleMessageUDP
func HandleMessageUDP(u *serverplugin.UDPConn, s *serverplugin.Session, msg []byte) {
	var f interface{}
	err := json.Unmarshal(msg, &f)
	if err != nil {
		return
	}
	m := f.(map[string]interface{})
	t := m["dataType"]
	_d := int32(t.(float64))
	switch _d {
	case PingUDP:
		playerID := m["playerID"].(string)
		player := world.GetPlayer(playerID)
		log.Println("PlayerID Self: " + player.PlayerID)
		v := u.GetAddr()
		log.Print("IP UDP Self: " + v.IP.String())
		log.Println(v.Port)
		player.Session.SetUDPAddr(v)
		players := world.GetPlayerList()
		for _, p := range players {
			msg := p.Session.GetUDPAddr()
			log.Print("IP UDP: " + msg.IP.String())
			log.Println(msg.Port)
		}
		break
	case RequestMove:
		x := m["x"]
		y := m["y"]
		r := m["rotation"]
		playerID := m["playerID"].(string)
		player := world.GetPlayer(playerID)
		player.X = float32(x.(float64))
		player.Y = float32(y.(float64))
		player.Rotation = int(r.(float64))
		playerResponse := CreatePlayerResponse(
			BroadcastMove,
			playerID,
			player.X,
			player.Y,
			player.Rotation,
		)
		log.Println("Player Response: " + string(playerResponse.ToJSON()))
		players := world.GetPlayerList()
		for _, p := range players {
			u.SetTAddr(p.Session.GetUDPAddr())
			u.SendMessage(playerResponse.ToJSON())
		}
		break
	}
}

// HandleMessage
func HandleMessage(s *serverplugin.Session, msg *serverplugin.Message) {
	msgID := msg.GetID()

	switch msgID {
	case RequestJoin:
		name := s.GetConn().GetName()
		player := CreatePlayer(name, s)
		m := make(map[string]interface{})
		m["self"] = player
		m["list"] = world.GetPlayerList()
		data, _ := json.Marshal(m)
		response := serverplugin.NewMessage(ResponseJoin, data)
		s.GetConn().SendMessage(response)

		for _, p := range world.GetPlayerList() {
			response := serverplugin.NewMessage(BroadcastJoin, player.ToJSON())
			p.Session.GetConn().SendMessage(response)
		}

		world.AddPlayer(player)

		s.BindUserID(player.PlayerID)
		break

	case RequestMove:
		log.Println("Request Move TCP")
		var f interface{}
		err := json.Unmarshal(msg.GetData(), &f)
		if err != nil {
			return
		}
		m := f.(map[string]interface{})
		x := m["x"]
		y := m["y"]
		r := m["rotation"]
		playerID := m["playerID"].(string)
		player := world.GetPlayer(playerID)
		player.X = float32(x.(float64))
		player.Y = float32(y.(float64))
		player.Rotation = int(r.(float64))

		players := world.GetPlayerList()

		for _, p := range players {
			message := serverplugin.NewMessage(BroadcastMove, player.ToJSON())
			p.Session.GetConn().SendMessage(message)
		}
		break
	}
}

// HandleDisconnect
func HandleDisconnect(s *serverplugin.Session, err error) {
	log.Println(s.GetConn().GetName() + " lost.")
	uid := s.GetUserID()
	lostPlayer := world.GetPlayer(uid)
	if lostPlayer == nil {
		return
	}

	world.RemovePlayer(uid)
	for _, p := range world.GetPlayerList() {
		message := serverplugin.NewMessage(BroadcastLeave, lostPlayer.ToJSON())
		p.Session.GetConn().SendMessage(message)
	}
}

// HandleConnect
func HandleConnect(s *serverplugin.Session) {
	log.Println(s.GetConn().GetName() + " connected.")
}
