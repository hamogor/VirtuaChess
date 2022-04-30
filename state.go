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
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}
}
