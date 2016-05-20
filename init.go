package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"
	"log"

	"github.com/google/subcommands"
	"golang.org/x/net/context"

	_ "github.com/mattn/go-sqlite3"
)

type initCmd struct {
	dbPath string
}

func (ic *initCmd) Name() string {
	return "init"
}

func (ic *initCmd) Synopsis() string {
	return "Initialize wishlist"
}

func (ic *initCmd) Usage() string {
	return `init
initiaize wishlist`
}

func (ic *initCmd) SetFlags(f *flag.FlagSet) {
	dir, err := wishlistDir()
	if err != nil {
		return
	}

	dbPath := filepath.Join(dir, "wishlist.db")
	f.StringVar(&ic.dbPath, "db", dbPath, "Database path")
}

func (ic *initCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	dir, err := wishlistDir()
	if err != nil {
		return subcommands.ExitFailure
	}

	os.MkdirAll(filepath.Join(dir), 0777)

	config := configFile{
		DBPath: ic.dbPath,
	}
	b, err := json.Marshal(&config)
	if err != nil {
		return subcommands.ExitFailure
	}

	configPath := filepath.Join(dir, "wishlist.json")
	file, err := os.Create(configPath)
	if err != nil {
		log.Printf("Failed open %s[%s]\n", configPath, err)
		return subcommands.ExitFailure
	}

	if _, err := file.Write(b); err != nil {
		log.Printf("Failed write to file[%s]\n", err)
		return subcommands.ExitFailure
	}

	dbPath := filepath.Join(dir, "wishlist.db")
	if err := initDBTable(dbPath); err != nil {
		log.Printf("Can't create DB[%s]\n", err)
		return subcommands.ExitFailure
	}

	fmt.Printf("Create wishlist Database in %s\n", dbPath)
	return subcommands.ExitSuccess
}

func initDBTable(dbPath string) error {
	os.Remove(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
CREATE TABLE if not EXISTS wishlist (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       product VARCHAR(255) NOT NULL,
       url VARCHAR(255) NOT NULL
);
`
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
