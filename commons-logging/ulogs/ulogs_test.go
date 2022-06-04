package ulogs

import (
	"testing"
)

func TestWarn(t *testing.T) {
	Debug("hello word")
	var fileName = "../../output/test.log.out"
	SetLogger(LoggerOptions{
		Path:    fileName,
		Rolling: false,
		Level:   LevelWarn,
	})

	Info("you can't see this log")
	Warn("you can see this log")
}
