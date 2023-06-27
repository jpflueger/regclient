//go:build wasi || wasm
// +build wasi wasm

package conffile

import (
	"os"
)

func homedir() string {
	home := os.Getenv(homeEnv)
	if home == "" {
		// wasi doesn't have a concept of home directory
		// so we just return the current directory because
		// this won't be an interactive session anyway
		home = "."
	}
	return home
}
