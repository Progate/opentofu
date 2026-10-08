//go:build !wasip1

package command

import "github.com/bgentry/speakeasy"

func askSecret() (string, error) { return speakeasy.Ask("") }
