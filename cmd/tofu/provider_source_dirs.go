//go:build !wasip1

package main

import "github.com/apparentlymart/go-userdirs/userdirs"

func sysSpecificPluginDirs() []string {
	return userdirs.ForApp("Terraform", "HashiCorp", "io.terraform").DataSearchPaths("plugins")
}
