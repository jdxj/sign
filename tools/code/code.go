// Note: 该文件仅用于测试

package main

type code int

const (
	// Err1 500 Reserve
	Err1 code = iota + 100000
	// Err2 400 abc
	Err2
	// Err5 567 hah a
	Err5
)

const Err3 int = 56
const Err4 int = 89

type Code struct {
	Name string
	HTTP int
	Desc string
}
