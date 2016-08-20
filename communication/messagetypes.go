package communication

import (
	"../common"
)

//inbound messages
const GameEnded string = "se.cygni.snake.api.event.GameEndedEvent"
const MapUpdated string = "se.cygni.snake.api.event.MapUpdateEvent"
const SnakeDead string = "se.cygni.snake.api.event.SnakeDeadEvent"
const GameStarting string = "se.cygni.snake.api.event.GameStartingEvent"
const PlayerRegistered string = "se.cygni.snake.api.response.PlayerRegistered"
const InvalidPlayerName string = "se.cygni.snake.api.exception.InvalidPlayerName"

//outbound messages
const RegisterPlayerMessageType string = "se.cygni.snake.api.request.RegisterPlayer"
const StartGame string = "se.cygni.snake.api.request.StartGame"
const RegisterMove string = "se.cygni.snake.api.request.RegisterMove"

const (
	small  = iota
	medium = iota
	large  = iota
)

//Outbound messages
type gameSettings struct {
	Width                    int  `json:"width"`
	Height                   int  `json:"height"`
	MaxNoofPlayers           int  `json:"maxNoofPlayers"`
	StartSnakeLength         int  `json:"startSnakeLength"`
	TimeInMsPerTick          int  `json:"timeInMsPerTick"`
	ObstaclesEnabled         bool `json:"obstaclesEnabled"`
	FoodEnabled              bool `json:"foodEnabled"`
	EdgeWrapsAround          bool `json:"edgeWrapsAround"`
	HeadToTailConsumes       bool `json:"headToTailConsumes"`
	TailConsumeGrows         bool `json:"tailConsumeGrows"`
	AddFoodLikelihood        int  `json:"addFoodLikelihood"`
	RemoveFoodLikelihood     int  `json:"removeFoodLikelihood"`
	AddObstacleLikelihood    int  `json:"addObstacleLikelihood"`
	RemoveObstacleLikelihood int  `json:"removeObstacleLikelihood"`
}

type gameMessage struct {
	Type string `json:"type"`
}

type playerRegistrationMessage struct {
	gameMessage
	PlayerName   string       `json:"playerName"`
	GameSettings gameSettings `json:"gameSettings"`
}

type registerMoveMessage struct {
	gameMessage
	Direction string `json:"direction"`
	GameId    string `json:"gameId"`
	GameTick  int    `json:"gameTick"`
}

type startGameMessage struct {
	gameMessage
}

type pingMessage struct {
	gameMessage
}

type ClientInfoMessage struct {
	gameMessage
	Language      string `json:"language"`
	OS            string `json:"operatingSystem"`
	Ip            string `json:"ipAddress"`
	ClientVersion string `json:"clientVersion"`
}

//Inbound messages
type PlayerRegisteredMessage struct {
	gameMessage
	Name             string       `json:"name"`
	GameId           string       `json:"gameId"`
	GameSettings     gameSettings `json:"gameSettings"`
	GameMode         string       `json:"gameMode"`
	RecivingPlayerId string       `json:"RecivingPlayerId"`
}

type MapUpdatedMessage struct {
	gameMessage
	GameTick         int    `json:"gameTick"`
	GameId           string `json:"gameId"`
	Map              Map    `json:"map"`
	RecivingPlayerId string `json:"receivingPlayerId"`
}

type GameEndedMessage struct {
	gameMessage
	PlayerWinnerId string `json:"playerWinnerId"`
	GameId         string `json:"gameId"`
	GameTick       int    `json:"gameTick"`
	Map            Map    `json:"map"`
}

type SnakeDeadMessage struct {
	gameMessage
	PlayerId    string `json:"playerId"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	GameId      string `json:"gameId"`
	GameTick    int    `json:"gameTick"`
	DeathReason string `json:"deathReason"`
}

type GameStartingMessage struct {
	gameMessage
	NoOfPlayers int `json:"noofPlayers"`
	Width       int `json:"width"`
	Height      int `json:"height"`
}

type InvalidPlayerNameMessage struct {
	gameMessage
	ReasonCode int `json:"reasonCode"`
}

type Map struct {
	Width             int              `json:"width"`
	Height            int              `json:"height"`
	WorldTick         int              `json:"worldTick"`
	SnakeInfos        []SnakeInfo      `json:"snakeInfos"`
	FoodPositions     common.Positions `json:"foodPositions"`
	ObstaclePositions common.Positions `json:"obstaclePositions"`
}

type SnakeInfo struct {
	Name                      string           `json:"name"`
	Points                    int              `json:"points"`
	Positions                 common.Positions `json:"positions"`
	TailProtectedForGameTicks int              `json:"tailProtectedForGameTicks"`
	Id                        common.Id        `json:"id"`
}
