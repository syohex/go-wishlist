package main

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"os"
)

func wishlistDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "wishlist"), nil
}

func wishlistConfig() (string, error) {
	dir, err := wishlistDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "wishlist.json"), nil
}

func isExisted(conf string) bool {
	if _, err := os.Stat(conf); os.IsNotExist(err) {
		return false
	}

	return true
}
