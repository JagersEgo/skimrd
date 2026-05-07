package main

import (
	"context"
	"errors"
	"jager/fast_reader/file_reader"
	"jager/fast_reader/processor"
	"jager/fast_reader/renderer"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Fast reader",
		Usage: "...",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "Verbose logging"},
			&cli.Int64Flag{Name: "wpm", Usage: "Display WPM speed", Value: 450},
		},
		Action: func(_ctx context.Context, cmd *cli.Command) error {
			var wpm = cmd.Int64("wpm")
			var verbose = cmd.Bool("verbose")

			// check for missing args
			if cmd.Args().Len() == 0 {
				log.Fatal(errors.New("No args"))
			}

			if verbose {
				log.Println("WPM: ", wpm)
			}

			setupSignalHandler()

			words := make(chan string)
			tokens := make(chan processor.Token, 32)

			go readWords(cmd.Args().First(), words, verbose)
			go tokeniseWords(words, tokens, verbose)

			var wg sync.WaitGroup
			wg.Go(func() { printTokens(wpm, tokens) })
			wg.Wait()

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func setupSignalHandler() {
	// channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		renderer.CleanUp()
		os.Exit(0)
	}()
}

func readWords(path string, words chan<- string, verbose bool) {
	defer close(words)
	file_reader.ReadWords(path, words)
	if verbose {
		log.Println("Finished reading inputs")
	}
}

func tokeniseWords(words <-chan string, tokens chan<- processor.Token, verbose bool) {
	defer close(tokens)
	for word := range words {
		tokens <- processor.Tokenise(word)
	}
	if verbose {
		log.Println("Finished tokenising words")
	}
}

func printTokens(wpm int64, tokens chan processor.Token) {
	renderer.PrintTokenStream(tokens, wpm)
}
