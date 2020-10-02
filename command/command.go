package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Commander reads user input and executes matching handler.
type Commander struct {
	Handlers       map[string]func()
	NotFoundPrompt string
}

// Input ...
func (c Commander) Input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("> ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	h, ok := c.Handlers[text]
	if !ok {
		fmt.Println(c.NotFoundPrompt)
		return
	}
	h()
}
