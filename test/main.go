package test

import (
	"fmt"
	"os"

	"github.com/congta/ugo-commons/keybus"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Please enter arguments")
		return
	}

	// args[0] is bin file

	var err error
	switch args[1] {
	case "key-bus":
		// key-bus --create -other 1 3 --hello --world=2 --w2 2 --sid=qunmus
		err = keybus.HandleKeyBusCmd(args[2:])
	default:
		err = fmt.Errorf("unknown command %s", args[1])
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
