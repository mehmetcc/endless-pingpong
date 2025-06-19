package game

import (
	"log"

	"github.com/mehmetcc/endless-pingpong/internal/player"
)

type Game struct {
	fp      *player.Player
	sp      *player.Player
	fpCh    chan string
	spCh    chan string
	doneCh  chan struct{}
	volleys int
}

func New() *Game {
	return &Game{
		fp:     player.New("Ron"),
		sp:     player.New("Dunn"),
		fpCh:   make(chan string),
		spCh:   make(chan string),
		doneCh: make(chan struct{}),
	}
}

func (g *Game) Start() {
	go g.play(g.fp, g.fpCh, g.spCh)
	go g.play(g.sp, g.spCh, g.fpCh)

	g.fpCh <- player.Ping // Start the match

	<-g.doneCh // Wait until someone fails
	log.Printf("Game over. Total volleys: %d", g.volleys)
}

func (g *Game) play(p *player.Player, in <-chan string, out chan<- string) {
	for move := range in {
		result := p.Play(move)
		if result == player.Failed {
			log.Printf("%s failed. Game ends.", p.Name)
			close(g.doneCh)
			close(out) // stop the other goroutine
			return
		}
		g.volleys++
		out <- result
	}
}
