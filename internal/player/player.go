package player

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	Ping   = "ping"
	Pong   = "pong"
	Failed = "failed"
)

type Player struct {
	Id   uuid.UUID
	Name string
}

func (p *Player) Play(received string) string {
	time.Sleep(1 * time.Second) // so that we can see the volley
	if randInt := rand.Intn(8); randInt == 0 {
		return Failed
	}

	var nextMove string
	if received == Ping {
		nextMove = Pong
	} else if received == Pong {
		nextMove = Ping
	} else {
		log.Printf("%s: received invalid move: %s", p.Name, received)
		return Failed
	}

	log.Printf("%s: plays %s", p.Name, nextMove)
	return nextMove
}

func New(name string) *Player {
	return &Player{
		Id:   uuid.New(),
		Name: name,
	}
}
