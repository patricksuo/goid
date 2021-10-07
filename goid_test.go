package goid

import (
	"sync"
	"testing"
)

func TestGoid(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			goidRuntime := goidFallback()
			goidAsm := Goid()
			if goidAsm == 0 || goidRuntime != goidAsm {
				t.Errorf("goid from runtime is %d, goid from asm is %d", goidRuntime, goidAsm)
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
