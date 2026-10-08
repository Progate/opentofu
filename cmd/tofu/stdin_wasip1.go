//go:build wasip1

package main

import (
	"os"
	"syscall"
)

// 端末からの入力が、届いてから読まれるようにする（→ scripts/build-tofu.mjs の (e)）。
// os.Stdin は os の init で非ブロッキングとして作られているので、フラグを戻したうえで作り直す
func init() {
	if err := syscall.SetNonblock(0, false); err == nil {
		os.Stdin = os.NewFile(0, "/dev/stdin")
	}
}
