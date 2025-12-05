package main

import (
	pathSize "code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name: "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(_ context.Context, ctx *cli.Command) error {
			if ctx.Args().Len() != 1 {
				return fmt.Errorf("usage: hexlet-path-size <path>")
			}
			path := ctx.Args().Get(0)
			size, err := pathSize.GetSize(path)
			if err != nil {
				return err
			}
			fmt.Println(size)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}