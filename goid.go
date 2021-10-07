// +build !amd64

package goid

func Goid() int64 {
	return goidFallback()
}