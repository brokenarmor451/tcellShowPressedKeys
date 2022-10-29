package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"	
)
func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if err := screen.Init(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	
	screen.SetStyle(defStyle)
	
	x := 0
	y := 0
	for {
		screen.Sync()
		switch event := screen.PollEvent().(type) {
			case *tcell.EventResize:
			screen.Sync()
		
			case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
			
			ShowPressedKeys(x + 10  , (y + 10) / 2 - 4, screen, event.Rune())
			x++
		}
	}
}

func ShowPressedKeys(X, Y int, Screen tcell.Screen, Key rune) {
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	Screen.SetContent(X, Y, Key, nil, style)
}

