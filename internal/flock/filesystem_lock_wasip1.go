//go:build wasip1

package flock

import (
	"context"
	"os"
)

// WASI には fcntl のレコードロックが無い。ブラウザーでは state を触るプロセスが 1 つしか
// 無いので、ロックは常に取れたことにする（取れないまま進むのではなく、競合が起きない）。
func Lock(*os.File) error                          { return nil }
func LockBlocking(context.Context, *os.File) error { return nil }
func Unlock(*os.File) error                        { return nil }
