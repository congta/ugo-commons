package ucommons

import (
	"fmt"
	"github.com/congta/ugo-commons/commons-logging/ulogs"
	"os"
	"runtime/debug"
)

func AssertNonErr2(err error, msg string) {
	if err != nil {
		ulogs.Error("validate error: %v, msg: %s, stack: %v", err, msg, string(debug.Stack()))
		fmt.Printf("发生非预期错误，工具退出： %v\n", err)
		os.Exit(-1)
	}
}

func AssertNonErr(err error) {
	if err != nil {
		ulogs.Error("validate error: %v, stack: %v", err, string(debug.Stack()))
		fmt.Printf("发生非预期错误，工具退出： %v\n", err)
		os.Exit(-1)
	}
}

func AssertOK(ok bool, msg string) {
	if !ok {
		ulogs.Info("assert failed. %s", string(debug.Stack()))
		fmt.Printf("[ERROR] %s\n", msg)
		os.Exit(-1)
	}
}
