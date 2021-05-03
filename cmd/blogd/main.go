package main

import (
	"os"

	"github.com/JIeeiroSst/blog/app"
	"github.com/JIeeiroSst/blog/cmd/blogd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
