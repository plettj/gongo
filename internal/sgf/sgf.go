package sgf

import (
	"os"
)

func FileToString(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

// TODO: Implement a function that converts an sgf string (or file) to a gng string
func ParseSgf(sgf string) string {
	return sgf
}
