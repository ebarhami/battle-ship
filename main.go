package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"app/game"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var n, nShip, nMissile int
	var s1, s2 string // ship
	var m1, m2 string // missile

	scanner := bufio.NewScanner(file)

	// Scan ships
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	nShip, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	s1 = scanner.Text()
	scanner.Scan()
	s2 = scanner.Text()

	// Scan missiles
	scanner.Scan()
	nMissile, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	m1 = scanner.Text()
	scanner.Scan()
	m2 = scanner.Text()

	ship1String := parseInputToArray(s1)
	ship2String := parseInputToArray(s2)

	m1String := parseInputToArray(m1)
	m2String := parseInputToArray(m2)

	g := game.NewGame(2, n)
	p1 := g.Players[0]
	p2 := g.Players[1]

	coordinates := toCoordinate(ship1String)
	for i := 0; i < nShip; i++ {
		p1.Board.PlaceShip(game.NewShip(coordinates[i]))
	}
	coordinates = toCoordinate(ship2String)
	for i := 0; i < nShip; i++ {
		p2.Board.PlaceShip(game.NewShip(coordinates[i]))
	}

	coordinates = toCoordinate(m1String)
	for i := 0; i < nMissile; i++ {
		p1.Hit(p2, coordinates[i])
	}
	coordinates = toCoordinate(m2String)
	for i := 0; i < nMissile; i++ {
		p2.Hit(p1, coordinates[i])
	}

	fmt.Printf("Player 1\n")
	p1.Board.Print()
	fmt.Printf("Player 2\n")
	p2.Board.Print()

	fmt.Printf("P1:%v\nP2:%v\n", p1.Score, p2.Score)
	winner, isDraw := g.GetWinnerSoFar()
	if isDraw {
		fmt.Printf("It is a draw\n")

	} else {
		if winner == p1 {
			fmt.Printf("Player 1 wins\n")
		} else {
			fmt.Printf("Player 2 wins\n")
		}
	}
}

func parseInputToArray(s string) []string {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	s = strings.ReplaceAll(s, " ", "")
	parsed := strings.Split(s, ":")
	return parsed
}

func toCoordinate(s []string) []game.Coordinate {
	coord := make([]game.Coordinate, 0)

	for _, c := range s {
		cSplit := strings.Split(c, ",")
		x, err := strconv.Atoi(cSplit[0])
		if err != nil {
			fmt.Printf("%v\n", cSplit[0])
			panic("Invalid")
		}
		y, err := strconv.Atoi(cSplit[1])
		if err != nil {
			fmt.Printf("%v\n", cSplit[1])
			panic("Invalid")
		}
		coord = append(coord, game.Coordinate{
			X: x,
			Y: y,
		})
	}
	return coord
}
