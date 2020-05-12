package game

import (
	"encoding/json"
)

// Player Player
type PlayerResponse struct {
	DataType int32   `json:"dataType"`
	PlayerID string  `json:"playerID"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	Rotation int     `json:"rotation"`
	Name     string  `json:"name"`
}

// CreatePlayer 创建玩家
func CreatePlayerResponse(dataType int32, name string, x float32, y float32, rotation int) *PlayerResponse {

	player := &PlayerResponse{
		DataType: dataType,
		PlayerID: name,
		Name:     name,
		X:        x,
		Y:        y,
		Rotation: rotation,
	}

	return player
}

// ToJSON 转成json数据
func (p *PlayerResponse) ToJSON() []byte {
	b, _ := json.Marshal(p)
	return b
}
