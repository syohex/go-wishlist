package main

import (
	"database/sql"
	"flag"
	"fmt"

	"strconv"

	"github.com/google/subcommands"
	"golang.org/x/net/context"
)

type delCmd struct {
}

func (d *delCmd) Name() string {
	return "delete"
}

func (d *delCmd) Synopsis() string {
	return "Delete item from wishlist"
}

func (d *delCmd) Usage() string {
	return `delete ID
`
}

func (*delCmd) SetFlags(*flag.FlagSet) {
}

func (d *delCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return subcommands.ExitFailure
	}

	stmt, err := db.Prepare(`DELETE FROM wishlist WHERE id = ?`)
	if err != nil {
		return subcommands.ExitFailure
	}

	for _, idStr := range f.Args() {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return subcommands.ExitFailure
		}

		if _, err := stmt.Exec(id); err != nil {
			return subcommands.ExitFailure
		}

		fmt.Printf("Delete item %d\n", id)
	}

	return subcommands.ExitSuccess
}
