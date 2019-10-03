// +build debug

package zip

import (
	"fmt"
	"os"
	"runtime"
)

func dbg(msg string, args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	prefix := fmt.Sprintf("DBG[%d]:%s: ", line, runtime.FuncForPC(pc).Name())
	fmt.Fprintf(os.Stderr, prefix+msg+"\n", args...)
}
