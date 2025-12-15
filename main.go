package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/XPLassal/simple-go-snake/render"
	. "github.com/XPLassal/simple-go-snake/structs"
	"github.com/eiannone/keyboard"
)

func main() {
s:
	cfg, exists := LoadConfig()

	if !exists {
		cfg = CreateConfig()
	}

	numbersOfColumns := cfg.Columns

	fps := 1000 / cfg.FPS

	baseMilliseconds := time.Duration(fps) * time.Millisecond
	ticker := time.NewTicker(baseMilliseconds)

	if cfg.HardMode {
		go func() {
			for {
				time.Sleep(time.Duration(numbersOfColumns/3) * time.Second)
				if baseMilliseconds > 50*time.Millisecond {
					baseMilliseconds -= 2 * time.Millisecond
				}
				ticker.Reset(baseMilliseconds)
			}
		}()
	}

	snake := NewSnake(numbersOfColumns)
	apple := NewApples(numbersOfColumns, snake)

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
				close(keyCh)
				return
			}
			keyCh <- char
		}
	}()

	writer := bufio.NewWriter(os.Stdout)
	var sb strings.Builder
	sb.Grow(numbersOfColumns * numbersOfColumns * 50)
	ClearConsole()

	moveSnake := func() bool {
		err := snake.Move(apple, numbersOfColumns)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(BrightGreen + Bold + "Your score: " + strconv.Itoa(snake.GetLen()) + Reset)
			return false
		}
		return true
	}

	for {
		sb.Reset()
		RenderField(numbersOfColumns, apple, snake, &sb, cfg.UseEmoji)
		writer.WriteString(sb.String())
		writer.Flush()
		select {
		case <-ticker.C:
			if !moveSnake() {
				return
			}
		case direction := <-keyCh:
			switch direction {
			case 'q':
				sb.WriteString("Game stopped, bye! 🙃")
				return
			case 'c':
				keyboard.Close()
				cfg = CreateConfig()
				fmt.Println("Okay, restart the game😉")
				time.Sleep(500 * time.Millisecond)
				goto s
			case 'p':
				direction = ' '
			case 'w', 'a', 's', 'd':
				snake.SetDirection(direction)
				<-ticker.C
				if !moveSnake() {
					return
				}
			}
		}
	}
}
