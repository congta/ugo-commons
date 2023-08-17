package ulogids

import (
	"hash/crc32"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	machinePosition uint64 = 0xffff000000000000
	pidCodePosition uint64 = 0x0000ff0000000000
	secondPosition  uint64 = 0x000000ffffff0000
	incPosition     uint64 = 0x000000000000ffff
)

var defaultGeneratorV0 = NewLogIdV0Generator()

// GetLogIdV0 return a uint64 number
func GetLogIdV0() uint64 {
	return defaultGeneratorV0.GetId()
}

type LogIdV0 struct {
	machineCode uint64 // 16bit
	pidCode     uint64 // 8bit
	second      int64  // 24bit
	inc         uint64 // 16bit
}

func NewLogIdV0Generator() *LogIdV0 {
	ld := &LogIdV0{
		second: time.Now().Unix(),
		inc:    0,
	}
	ld.machineCode, ld.pidCode = ld.machinePidCode()
	return ld
}

func (ld *LogIdV0) machinePidCode() (uint64, uint64) {
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	m32Code := crc32.ChecksumIEEE([]byte(hostname + strconv.Itoa(pid)))
	return uint64(m32Code), uint64(pid)
}

func (ld *LogIdV0) GetId() uint64 {
	now := time.Now().Unix()
	inc := atomic.AddUint64(&ld.inc, 1)
	ts := uint64(now)
	var ret uint64
	ret = (ld.machineCode << 48) & machinePosition
	ret = ret | ((ld.pidCode << 40) & pidCodePosition)
	ret = ret | ((ts << 16) & secondPosition)
	ret = ret | (inc & incPosition)
	return ret
}
