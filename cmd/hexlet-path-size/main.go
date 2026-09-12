package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	path := ""
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
		    &cli.BoolFlag{
		        Name: "human",
			Aliases: []string{"H"},
			Value: false,
			Usage: "human-readable sizes (auto-select unit)",
		    },
		    &cli.BoolFlag{
		        Name: "all",
			Aliases: []string{"a"},
			Value: false,
			Usage: "include hidden files and directories",
		    },
		    &cli.BoolFlag{
			Name:"recursive",
			Aliases: []string{"r"},
			Value: false,
			Usage: "recursive size of directories (default: false)",
		    },
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
		    
		    path = cmd.Args().Get(0)
		    return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	inf, err := code.GetPathSize(path, cmd.Bool("human"), cmd.Bool("all"), cmd.Bool("recursive"))
	if err != nil {
	        fmt.Println(err)
	} else {
		fmt.Printf("%s\t%s\n", inf, path)
	}

}
