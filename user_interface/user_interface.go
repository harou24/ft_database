package userinterface

import (
	"bufio"
	"fmt"
	"os"
)

func InteractivePrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("ft_database :> ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Error: %v", err)
		}
		if userInput == "exit\n" {
			break
		}
		fmt.Print(userInput)
	}
}
