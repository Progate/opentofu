//go:build wasip1

package command

import (
	"github.com/opentofu/opentofu/internal/command/views"
	"github.com/opentofu/opentofu/internal/repl"
	"github.com/opentofu/opentofu/internal/tfdiags"
)

// 端末を termios で握れないので、対話モードは持たない。式をパイプで渡す形（modePiped）は動く。
func (c *ConsoleCommand) modeInteractive(_ *repl.Session, view views.Console) int {
	view.Diagnostics(tfdiags.Diagnostics{}.Append(tfdiags.Sourceless(
		tfdiags.Error,
		"Interactive console is unavailable",
		"This build has no interactive console. Pipe expressions into `tofu console` instead.",
	)))
	return 1
}
