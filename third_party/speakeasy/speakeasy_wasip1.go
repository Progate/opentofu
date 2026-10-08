//go:build wasip1

package speakeasy

import (
	"bufio"
	"os"
	"strings"
)

// WASI には termios が無いので、エコーは切れない。そのまま 1 行読む。
func getPassword() (string, error) {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(line, "\r\n"), err
}
