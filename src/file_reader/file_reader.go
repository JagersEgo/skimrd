package file_reader

import (
	"bufio"
	"log"
	"os"
)

func ReadWords(path string, ch chan<- string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error() + ": " + path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // key improvement

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
