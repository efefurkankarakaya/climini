package utility

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(prefix string) string {
	fmt.Print(prefix)
	in := bufio.NewReader(os.Stdin)
	text, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return text
}
