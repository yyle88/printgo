package printgo

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
)

func NewPTX() *PTX {
	return &PTX{}
}

type PTX struct{ bytes.Buffer }

func (ptx *PTX) Println(args ...interface{}) (n int) {
	return done.VE(fmt.Fprintln(ptx, args...)).Done()
}

func (ptx *PTX) Print(args ...interface{}) (n int) {
	n, err := fmt.Fprint(ptx, args...)
	must.Done(err)
	return n
}

func (ptx *PTX) Fprintf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(ptx, format, args...))
}

func (ptx *PTX) Printf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(ptx, format, args...))
}

func NewPTS() *PTS {
	return &PTS{}
}

type PTS struct{ strings.Builder }

func (pts *PTS) Println(args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintln(pts, args...))
}

func (pts *PTS) Print(args ...interface{}) (n int) {
	n, err := fmt.Fprint(pts, args...)
	must.Done(err)
	return n
}

func (pts *PTS) Fprintf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(pts, format, args...))
}

func (pts *PTS) Printf(format string, args ...interface{}) (n int) {
	return rese.V1(fmt.Fprintf(pts, format, args...))
}
