package main

import (
	"fmt"
	"time"

	. "github.com/XPLassal/simple-snake-on-go/render"
	. "github.com/XPLassal/simple-snake-on-go/structs"
	"github.com/eiannone/keyboard"
)

func main() {
	var numbersOfColumns int
	var gameAcceleration string
	GetNumberOfColumns(&numbersOfColumns)

	fmt.Print("Do you want the game speed to increase as you progress? (y/n): ")
	fmt.Scan(&gameAcceleration, "\n")

	baseMilliseconds := 200 * time.Millisecond
	ticker := time.NewTicker(baseMilliseconds)

	if gameAcceleration == "y" {
		go func() {
			for {
				time.Sleep(time.Duration(numbersOfColumns / 3))
				if baseMilliseconds > 50*time.Millisecond {
					baseMilliseconds -= 2 * time.Millisecond
				}
				ticker.Reset(baseMilliseconds)
			}
		}()
	}

	snake := NewSnake(numbersOfColumns)
	apple := NewApples(numbersOfColumns, &snake)

	if err := keyboard.Open(); err != nil {
		fmt.Println(Red + Bold + "Error: " + err.Error() + Reset)
		return
	}
	defer keyboard.Close()

	keyCh := make(chan rune)
	go func() {
		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				fmt.Println(Red + Bold + "Error: " + err.Error() + Reset)
				return
			}
			keyCh <- char
		}
	}()

	for {
		RenderField(numbersOfColumns, &apple, &snake)
		select {
		case <-ticker.C:
			err := snake.Move(&apple, numbersOfColumns)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(BrightGreen+Bold+"Your score: ", snake.GetLen(), Reset)
				return
			}
		case direction := <-keyCh:
			if direction == 'q' {
				fmt.Println("Game stopped, bye! 🙃")
				return
			}
			snake.SetDirection(direction)
		}
	}
}
