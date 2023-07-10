package game

type Board struct {
	Grid        [][]rune
	Ships       []*Ship
	GridToShips map[Coordinate]*Ship
}

func NewBoard(size int) *Board {
	grid := make([][]rune, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			grid[i][j] = '_'
		}
	}
	return &Board{
		Grid:        grid,
		Ships:       make([]*Ship, 0),
		GridToShips: make(map[Coordinate]*Ship),
	}
}
