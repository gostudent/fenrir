package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xwb1989/sqlparser"

	"github.com/kniren/gota/dataframe"
	"github.com/urfave/cli"
)

func handleSelect(stmt *sqlparser.Select) error {
	// fmt.Println("SELECT Statement")
	// fmt.Println(stmt)
	// fmt.Println("Cache :- ", stmt.Cache)
	// fmt.Println("Comments :- ", stmt.Comments)
	// fmt.Println("Distinct :- ", stmt.Distinct)
	// fmt.Println("Hints :- ", stmt.Hints)
	// fmt.Println("SelectExprs :- ", stmt.SelectExprs)
	// fmt.Println("From :- ", stmt.From)
	// fmt.Println("Where :- ", stmt.Where)
	// fmt.Println("GroupBy :- ", stmt.GroupBy)
	// fmt.Println("Having :- ", stmt.Having)
	// fmt.Println("OrderBy :- ", stmt.OrderBy)
	// fmt.Println("Limit :- ", stmt.Limit)
	// fmt.Println("Lock :- ", stmt.Lock)

	if len(stmt.From) != 1 {
		fmt.Println("ERROR :- Multiple from is not supported")
	}

	table := sqlparser.String(stmt.From)
	fmt.Println("Table Name :- ", table)

	return nil
}

func handleInsert(stmt *sqlparser.Insert) error {
	table := stmt.Table.Name
	fmt.Println("Table Name :- ", table)

	return nil
}

func fenrir(c *cli.Context) error {
	fmt.Println("fenrir v0.0.1")
	fmt.Println("")
	csv := c.Args().Get(0)
	r, err := os.Open(csv)
	if err != nil {
		return err
	}
	df := dataframe.ReadCSV(r)
	fmt.Println(df)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		sql, _ := reader.ReadString('\n')
		//fmt.Println(sql)
		if sql == "exit" {
			break
		}
		stmt, err := sqlparser.Parse(sql)
		if err != nil {
			return err
		}
		switch stmt := stmt.(type) {
		case *sqlparser.Select:
			handleSelect(stmt)
		case *sqlparser.Insert:
			handleInsert(stmt)
		}
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "fenrir"
	app.Usage = "Run SQL Queries on CSV"
	app.Version = "v0.0.1"
	app.Action = fenrir

	app.Run(os.Args)
}
