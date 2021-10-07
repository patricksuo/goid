package goid

import (
	"bytes"
	"runtime"
	"strconv"
)

func goidFallback() int64 {
	buf := make([]byte, 20)
	buf = buf[:runtime.Stack(buf, false)]
	goidStr := string(bytes.Split(buf, []byte(" "))[1])
	result, _ := strconv.ParseInt(goidStr, 10, 64)
	return result
}
