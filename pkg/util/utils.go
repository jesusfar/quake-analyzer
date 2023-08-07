package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func PrettyPrint(out any) {
	prettyOut, _ := json.MarshalIndent(out, "", "    ")
	fmt.Println(string(prettyOut))
}

func GetFilePath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(absPath); err != nil {
		panic(err)
	}

	return absPath
}
