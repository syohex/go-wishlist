package main

import (
	"database/sql"
	"flag"

	"github.com/google/subcommands"
	"golang.org/x/net/context"
	"log"
	"fmt"
)

type listCmd struct {
	short bool
	verbose bool
}

func (l *listCmd) Name() string {
	return "list"
}

func (l *listCmd) Synopsis() string {
	return "List wishlist"
}

func (l *listCmd) Usage() string {
	return `list [-s] [-v]
List show list
-s List only product name
-v Verbose`
}

func (l *listCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&l.short, "s", false, "short output")
	f.BoolVar(&l.verbose, "v", false, "verbose output")
}

func (l *listCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return subcommands.ExitFailure
	}

	rows, err := db.Query(`SELECT id, product, url FROM wishlist ORDER BY id`)
	if err != nil {
		log.Printf("Failed: SELECT FROM wishlist[%s]", err)
		return subcommands.ExitFailure
	}
	defer rows.Close()

	var id int
	var product, url string
	for rows.Next() {
		if err := rows.Scan(&id, &product, &url); err != nil {
			log.Print("Failed: binding variable")
			return subcommands.ExitFailure
		}

		if l.verbose {
			fmt.Printf("%2d: %s [%s]\n", id, product, url)
		} else {
			fmt.Printf("%2d: %s\n", id, product)
		}

	}

	return subcommands.ExitSuccess
}
