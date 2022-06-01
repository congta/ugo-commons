package test

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/congta/ugo-commons/commons-io/ufiles"
	"github.com/congta/ugo-commons/commons-lang/unumbers"
	"github.com/congta/ugo-commons/commons-u/ucommons"
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
		err = handleKeyBus(args[2:])
	default:
		err = fmt.Errorf("unknown command %s", args[1])
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func handleKeyBus(args []string) error {
	if args[0] == "--create" {
		// e.g. --create --sid=xxx -n 16
		argsMap := ucommons.GetArgsMap(args)
		sid, _ := argsMap["--sid"]
		if sid == "" {
			fmt.Println("--sid is required")
			return nil
		}
		num := unumbers.ToInt(argsMap["-n"], 16)
		fmt.Printf("make key box secret file, sid=%s, num=%d\n", sid, num)

		fileName := keybus.GetKeyBusFileName(sid)
		if ok, _ := ufiles.Exists(fileName); ok {
			return fmt.Errorf("file is already exist %s", fileName)
		}

		iv := make([]byte, 0)
		result := make([]string, 0, num+1)
		result = append(result, "#! MODE=CBC")
		for i := 0; i < num; i++ {
			key := make([]byte, 32)

			if _, err := rand.Read(key); err != nil {
				return err
			}

			holder := &keybus.KeyHolder{
				Key: key,
				Iv:  iv,
				Id:  i + 1,
			}
			result = append(result, holder.ToString())
		}
		if err := ufiles.WriteLines(fileName, result); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Unknown option " + args[0])
}
