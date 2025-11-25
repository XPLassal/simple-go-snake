package structs

type Apple struct {
	body map[Coordinates]struct{}
}

func (apple *Apple) AddNew(numberOfColumns int, snake *Snake) {
	for {
		coords := NewCoordinates(numberOfColumns)

		if apple.Contains(coords) {
			continue
		}

		_, isBody := snake.Contains(coords)
		if isBody {
			continue
		}

		apple.body[coords] = struct{}{}
		break
	}
}

func (apple *Apple) Contains(coords Coordinates) bool {
	_, ok := apple.body[coords]
	return ok
}

func (apple *Apple) GetLen() int {
	return len(apple.body)
}

func NewApples(numberOfColumns int, snake *Snake) Apple {
	cap := numberOfColumns*2 - 1
	apple := Apple{body: make(map[Coordinates]struct{}, cap)}

	for apple.GetLen() < cap {
		apple.AddNew(numberOfColumns, snake)
	}
	return apple
}

func (apple *Apple) EatApple(coords Coordinates, numberOfColumns int, snake *Snake) {
	delete(apple.body, coords)
	if numberOfColumns*numberOfColumns-snake.GetLen() <= apple.GetLen() {
		return
	}

	apple.AddNew(numberOfColumns, snake)
}
