package game

type Game struct {
	Players []*Player
}

func NewGame(nPlayers int, boardSize int) *Game {
	if nPlayers != 2 {
		return nil
	}

	players := make([]*Player, nPlayers)
	for i := 0; i < nPlayers; i++ {
		players[i] = NewPlayer(boardSize)
	}
	return &Game{
		Players: players,
	}
}

func (g *Game) GetWinnerSoFar() (winner *Player, isDraw bool) {
	score := -1
	for _, p := range g.Players {
		if p.Score > score {
			score = p.Score
			winner = p
		}
	}
	// check whether there are two player with the same score, it is draw
	for _, p := range g.Players {
		if p.Score == winner.Score && p.Id != winner.Id {
			winner = nil
			isDraw = true
			return
		}
	}
	return
}
