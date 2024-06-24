package keybus

import (
	"fmt"
	"strings"
	"testing"
)

func TestPipe(t *testing.T) {
	s := ""
	box := NewKeyBusByLegacyFile("maptime-prod")
	fmt.Println(box.DecryptStr(s))
}

func TestInstall(t *testing.T) {
	err := HandleKeyBusCmd(strings.Split("--create --sid=qunmus-prod -n 32", " "))
	if err != nil {
		panic(err)
	}
}
