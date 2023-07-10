package game

type Board struct {
	Grid        [][]rune
	Ships       []*Ship
	GridToShips map[Coordinate]*Ship
}
