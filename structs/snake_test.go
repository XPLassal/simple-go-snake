package structs

import (
	"testing"
)

func growSnake(s *Snake, length int) {
	current := s.head
	for i := 0; i < length; i++ {
		next := Coordinates{x: i + 1}
		s.body[current] = next
		current = next
	}
	s.tail = current
}

func BenchmarkMove_Short(b *testing.B) {
	n := 100
	snake := NewSnake(n)
	growSnake(snake, 10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		head := snake.head
		newHead := Coordinates{head.x + 1, head.y}

		snake.body[head] = newHead
		snake.body[newHead] = newHead
		snake.head = newHead

		oldTail := snake.tail
		newTail := snake.body[oldTail]
		delete(snake.body, oldTail)
		snake.tail = newTail
	}
}

func BenchmarkMove_Medium(b *testing.B) {
	n := 2000
	snake := NewSnake(n)
	growSnake(snake, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := snake.head
		newHead := Coordinates{head.x + 1, head.y}

		snake.body[head] = newHead
		snake.body[newHead] = newHead
		snake.head = newHead

		oldTail := snake.tail
		newTail := snake.body[oldTail]
		delete(snake.body, oldTail)
		snake.tail = newTail
	}
}

func BenchmarkMove_Huge(b *testing.B) {
	n := 10000
	snake := NewSnake(n)
	growSnake(snake, 10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := snake.head
		newHead := Coordinates{head.x + 1, head.y}

		snake.body[head] = newHead
		snake.body[newHead] = newHead
		snake.head = newHead

		oldTail := snake.tail
		newTail := snake.body[oldTail]
		delete(snake.body, oldTail)
		snake.tail = newTail
	}
}
