package server

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func initializeWords(minLen, maxLen int) ([]string, error) {
	var ws []string

	path, err := filepath.Abs("conf/nounlist.txt")
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		w := scanner.Text()
		if len(w) <= maxLen && len(w) >= minLen && len(w) != 0 {
			w = strings.Title(w)
			ws = append(ws, w)
		}
	}
	return ws, nil
}
