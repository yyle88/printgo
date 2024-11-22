package printgo

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yyle88/done"
)

func NewPTX() *PTX {
	return &PTX{}
}

type PTX struct{ bytes.Buffer }

func (T *PTX) Println(args ...interface{}) (n int) {
	return done.VE(fmt.Fprintln(T, args...)).Done()
}

func (T *PTX) Print(args ...interface{}) (n int) {
	return done.VE(fmt.Fprint(T, args...)).Done()
}

func NewPTS() *PTS {
	return &PTS{}
}

type PTS struct{ strings.Builder }

func (T *PTS) Println(args ...interface{}) (n int) {
	return done.VE(fmt.Fprintln(T, args...)).Done()
}

func (T *PTS) Print(args ...interface{}) (n int) {
	return done.VE(fmt.Fprint(T, args...)).Done()
}
