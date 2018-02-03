package main

import (
	"fmt"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fenrir"
	app.Usage = "Run SQL Queries on CSV"
	app.Action = func(c *cli.Context) error {
		csv := c.Args().Get(0)
		r, err := os.Open(csv)
		if err != nil {
			return err
		}
		df := dataframe.ReadCSV(r)
		fmt.Println(df)
		return nil
	}

	app.Run(os.Args)
}
