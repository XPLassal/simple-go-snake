package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Scan(&gameAcceleration)

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

	writer := bufio.NewWriter(os.Stdout)
	var sb strings.Builder
	sb.Grow(numbersOfColumns * numbersOfColumns * 20)
	ClearConsole()

	for {
		sb.Reset()
		RenderField(numbersOfColumns, &apple, &snake, &sb)
		writer.WriteString(sb.String())
		writer.Flush()
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
