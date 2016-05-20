package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"golang.org/x/net/context"
)

type addCmd struct {
}

func (*addCmd) Name() string {
	return "add"
}

func (*addCmd) Synopsis() string {
	return "Add to wishlist"
}

func (*addCmd) Usage() string {
	return `add product URL
Add product and URL into database
`
}

func (*addCmd) SetFlags(*flag.FlagSet) {
}

func (a *addCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) < 2 {
		return subcommands.ExitUsageError
	}

	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return subcommands.ExitFailure
	}

	stmt, err := db.Prepare(`INSERT INTO wishlist (product, url) VALUES (?,?)`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	product := args[0]
	url := args[1]

	if _, err := stmt.Exec(product, url); err != nil {
		return subcommands.ExitFailure
	}

	fmt.Printf("Add '%s' into wishlist\n", product)
	return subcommands.ExitSuccess
}
