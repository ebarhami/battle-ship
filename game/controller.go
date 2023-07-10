package game

type Game struct {
	Players []*Player
}

func NewGame(nPlayers int, boardSize int) *Game {
	players := make([]*Player, nPlayers)
	for i := 0; i < nPlayers; i++ {
		players[i] = NewPlayer(boardSize)
	}
	return &Game{
		Players: players,
	}
}
