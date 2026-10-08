//go:build wasip1

package main

// go-userdirs は GOOS で分岐し、知らない OS では panic する。ブラウザーには XDG も
// Library も無く、プラグインの置き場は CLI 設定（provider_installation）で 1 か所に
// 決まるので、ここから足すものは無い。
func sysSpecificPluginDirs() []string { return nil }
