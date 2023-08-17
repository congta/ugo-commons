package ulogids

import (
	"encoding/hex"
	"github.com/bytedance/gopkg/lang/fastrand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	version    = "02"
	length     = 53
	maxRandNum = 1<<24 - 1<<20
	defaultIp6 = "00000000000000000000000000000000" // strings.Repeat("0", 32)
)

var defaultLogIdV2Generator = NewLogIdV2Generator()
var localIP6 = findIP6()

func GetLogId() string {
	return GetLogIdV2()
}

func GetLogIdV2() string {
	return defaultLogIdV2Generator.GetLogId()
}

// LogIdV2 represents a logId generator
type LogIdV2 struct{}

// NewLogIdV2Generator create a new LogId instance
func NewLogIdV2Generator() *LogIdV2 {
	return &LogIdV2{}
}

func (l LogIdV2) GetLogId() string {
	r := fastrand.Uint32n(maxRandNum) + 1<<20
	sb := strings.Builder{}
	sb.Grow(length)
	sb.WriteString(version)
	sb.WriteString(strconv.FormatUint(uint64(getTimeInMs()), 10))
	sb.Write(localIP6)
	sb.WriteString(strconv.FormatUint(uint64(r), 16))
	return sb.String()
}

func getTimeInMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func findIP6() []byte {
	interfaces, err := net.Interfaces()
	if err != nil {
		return []byte(defaultIp6)
	}

	backup := []byte(defaultIp6)
	for _, i := range interfaces {
		addresses, _ := i.Addrs()
		for _, addr := range addresses {
			switch v := addr.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() {
					dst := make([]byte, 32)
					hex.Encode(dst, v.IP.To16())
					if len(v.IP) == net.IPv6len {
						return dst
					} else if len(v.IP) == net.IPv4len {
						backup = dst
					}
				}
			}
		}
	}
	return backup
}
