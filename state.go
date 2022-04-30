package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (g *Game) Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Virtua Chess")
	fmt.Println("---------------------")

	// Player 1 chooses input
	// Player 2 chooses input
	//
	for {
		fmt.Print("Player Input-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("p", text) == 0 {
			fmt.Println("Inputted: Punch")
		}
		if strings.Compare("6p", text) == 0 {
			fmt.Println("Inputted: Elbow")
		}
		if strings.Compare("2k", text) == 0 {
			fmt.Println("Inputted: Low Kick")
		}
		if strings.Compare("6p+g", text) == 0 {
			fmt.Println("Inputted: Throw")
		}

	}
}
