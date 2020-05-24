package game

import (
	"encoding/json"
	"math/rand"

	serverplugin "github.com/yohanesgre/server-fps/network"
)

// Player Player
type Player struct {
	PlayerID string                `json:"playerID"`
	X        float32               `json:"x"`
	Y        float32               `json:"y"`
	Rotation int                   `json:"rotation"`
	Name     string                `json:"name"`
	Session  *serverplugin.Session `json:"-"`
}

// CreatePlayer 创建玩家
func CreatePlayer(name string, s *serverplugin.Session) *Player {

	player := &Player{
		PlayerID: name,
		Name:     name,
		X:        float32(rand.Intn(10)),
		Y:        float32(rand.Intn(10)),
		Rotation: 0,
		Session:  s,
	}

	return player
}

// ToJSON 转成json数据
func (p *Player) ToJSON() []byte {
	b, _ := json.Marshal(p)
	return b
}
