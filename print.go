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

func (ptx *PTX) Println(args ...interface{}) (n int) {
	return done.VE(fmt.Fprintln(ptx, args...)).Done()
}

func (ptx *PTX) Print(args ...interface{}) (n int) {
	return done.VE(fmt.Fprint(ptx, args...)).Done()
}

func NewPTS() *PTS {
	return &PTS{}
}

type PTS struct{ strings.Builder }

func (pts *PTS) Println(args ...interface{}) (n int) {
	return done.VE(fmt.Fprintln(pts, args...)).Done()
}

func (pts *PTS) Print(args ...interface{}) (n int) {
	return done.VE(fmt.Fprint(pts, args...)).Done()
}
