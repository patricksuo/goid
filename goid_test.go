package goid

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGoid(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			buf := make([]byte, 20)
			buf = buf[:runtime.Stack(buf, false)]
			goidRuntime := string(bytes.Split(buf, []byte(" "))[1])
			goidAsm := fmt.Sprintf("%d", Goid())
			if goidRuntime != goidAsm {
				t.Error("goid from runtime is %s, goid from asm is %s", goidRuntime, goidAsm)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkGoid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Goid()
	}
}
