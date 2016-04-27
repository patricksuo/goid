obtain goroutine id outside go runtime.
linux/amd64 go1.6

#benchmark
go test -bench=.
PASS
BenchmarkGoid-2	2000000000	         2.04 ns/op
