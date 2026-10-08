//go:build wasip1

package command

import (
	"bufio"
	"os"
	"strings"
)

// エコーを切る手段が無い。入力を隠す責任はブラウザー側の端末が持つ。
func askSecret() (string, error) {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(line, "\r\n"), err
}
