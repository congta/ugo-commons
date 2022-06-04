package keybus

import (
	"strings"
	"testing"
)

func TestPipe(t *testing.T) {

}

func TestInstall(t *testing.T) {
	err := HandleKeyBusCmd(strings.Split("--create --sid=qunmus-prod -n 32", " "))
	if err != nil {
		panic(err)
	}
}
