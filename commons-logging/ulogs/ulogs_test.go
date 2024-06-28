package ulogs

import (
	"testing"
)

func TestWarn(t *testing.T) {
	DebugV1("hello word")
	var fileName = "../../output/test.log.out"
	SetLoggerV1(LoggerOptions{
		Path:    fileName,
		Rolling: false,
		Level:   LevelWarn,
	})

	InfoV1("you can't see this log")
	WarnV1("you can see this log")
}
