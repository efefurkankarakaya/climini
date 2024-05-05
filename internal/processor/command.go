package processor

import (
	"fmt"
	"os"
)

func ReadArguments() {
	fmt.Println(len(os.Args[1:]), os.Args[1:])
}
