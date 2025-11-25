package structs

import "math/rand"

type Coordinates struct {
	x, y int
}

func SetCoordinates(x, y int) Coordinates {
	return Coordinates{x, y}
}

func NewCoordinates(n int) Coordinates {
	return Coordinates{rand.Intn(n), rand.Intn(n)}
}

func (coords *Coordinates) GenerateNewCoords(n int) {
	*coords = Coordinates{rand.Intn(n), rand.Intn(n)}
}
