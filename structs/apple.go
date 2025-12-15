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

func NewApples(numberOfColumns int, snake *Snake, multiAppleMode bool) *Apple {
	const appleDensity = 0.12
	totalSpace := numberOfColumns*numberOfColumns - 1
	maxApples := 1
	if multiAppleMode {
		maxApples = int(float64(totalSpace) * appleDensity)
	}
	apple := Apple{body: make(map[Coordinates]struct{}, maxApples)}

	for apple.GetLen() < maxApples {
		apple.AddNew(numberOfColumns, snake)
	}
	return &apple
}

func (apple *Apple) EatApple(coords Coordinates, numberOfColumns int, snake *Snake) {
	delete(apple.body, coords)
	if numberOfColumns*numberOfColumns-snake.GetLen() <= apple.GetLen() {
		return
	}

	apple.AddNew(numberOfColumns, snake)
}
