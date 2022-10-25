package core

import (
	"os"
	"path"
)

func GetUserFolder() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, "Documents")
}
