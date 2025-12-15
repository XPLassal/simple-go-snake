package structs

import (
	"errors"
	"fmt"
	"math"
)

type Snake struct {
	head      Coordinates
	tail      Coordinates
	body      map[Coordinates]Coordinates
	direction string
}

func (snake *Snake) SetDirection(r rune) {
	currentDir := snake.direction
	newDir := string(r)

	if (currentDir == "w" && newDir == "s") ||
		(currentDir == "s" && newDir == "w") ||
		(currentDir == "a" && newDir == "d") ||
		(currentDir == "d" && newDir == "a") {
		return
	}
	snake.direction = newDir
}

func (snake *Snake) GetDirection() string {
	return snake.direction
}

func (snake *Snake) GetLen() int {
	return len(snake.body)
}

func NewSnake(n int) *Snake {
	center := Coordinates{n / 2, n / 2}
	s := Snake{body: make(map[Coordinates]Coordinates, (n * n))}
	s.head, s.tail = center, center
	s.body[center] = center
	return &s
}

func (snake *Snake) Contains(coords Coordinates) (isHead bool, isBodyOfSnake bool) {
	if coords == snake.head {
		return true, true
	}

	_, ok := snake.body[coords]
	return false, ok
}

func (snake *Snake) DebugSnake(numberOfColumns int) string {
	return fmt.Sprintf("\nHead: %v\nTail: %v\nBody: %v\nNumberOfCOlumns: %d", snake.head, snake.tail, snake.body, numberOfColumns)
}

func (snake *Snake) Move(apple *Apple, numbersOfColumns int, allowWallPass bool) error {
	var dx, dy int
	switch snake.GetDirection() {
	case "w":
		dy = -1
	case "s":
		dy = 1
	case "a":
		dx = -1
	case "d":
		dx = 1
	default:
		return nil
	}

	head := snake.head
	tail := snake.tail
	newHead := Coordinates{head.x + dx, head.y + dy}

	if allowWallPass {
		if newHead.x < 0 || newHead.x >= numbersOfColumns {
			newHead.x = numbersOfColumns - int(math.Abs(float64(newHead.x)))
		}

		if newHead.y < 0 || newHead.y >= numbersOfColumns {
			newHead.y = numbersOfColumns - int(math.Abs(float64(newHead.y)))
		}
	} else {
		if newHead.x < 0 || newHead.y < 0 || newHead.x >= numbersOfColumns || newHead.y >= numbersOfColumns {
			return errors.New("The snake went out of bounds :(")
		}
	}

	if _, exists := snake.Contains(newHead); exists {
		if tail != newHead {
			return errors.New("The snake crashed into itself :(")
		}
	}

	if apple.Contains(newHead) {
		apple.EatApple(newHead, numbersOfColumns, snake)
	} else {
		snake.tail = snake.body[tail]
		delete(snake.body, tail)
	}

	snake.body[head] = newHead
	snake.head = newHead
	snake.body[newHead] = newHead

	return nil
}
