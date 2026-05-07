package file_reader

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadWordsFromFile(path string, ch chan<- string) {
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

func ReadWordsFromPipe(r io.Reader, ch chan<- string) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
