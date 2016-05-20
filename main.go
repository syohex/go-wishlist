package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/context"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	flag.Parse()
	subcmd := flag.Args()[0]

	var db *sql.DB
	if subcmd != "init" {
		configFile, err := wishlistConfig()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		if !isExisted(configFile) {
			fmt.Fprintf(os.Stderr, "Config file %s not found. Please 'wishlist init'\n", configFile)
			return 1
		}

		conf, err := readConfig(configFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		db, err = sql.Open("sqlite3", conf.DBPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		defer db.Close()
	}

	ctx := context.WithValue(context.Background(), "db", db)

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&initCmd{}, "")
	subcommands.Register(&addCmd{}, "")
	subcommands.Register(&listCmd{}, "")
	subcommands.Register(&delCmd{}, "")

	return int(subcommands.Execute(ctx))
}
