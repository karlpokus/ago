package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/karlpokus/ago"
)

const usage = "pass date in ISO-8601 or - to read from stdin"

func main() {
	stdout := log.New(os.Stdout, "", 0)
	stderr := log.New(os.Stderr, "", 0)
	if len(os.Args) < 2 {
		stderr.Fatalf("error: missing argument: %s", usage)
	}
	date := os.Args[1]
	switch date {
	case "-h":
		stdout.Println(usage)
		return
	case "-":
		b, err := readFile(os.Stdin, 2*time.Second)
		if err != nil {
			stderr.Fatal(err)
		}
		date = string(b)
	}
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		stderr.Fatal(err)
	}
	stdout.Println(ago.ParseWithContext(t))
}

// readFile reads from f and returns the result unless d expires.
func readFile(f *os.File, d time.Duration) ([]byte, error) {
	res := make(chan []byte)
	errc := make(chan error)
	go func() {
		buf := make([]byte, 1024)
		n, err := f.Read(buf)
		if err != nil {
			errc <- err
			return
		}
		res <- bytes.TrimSpace(buf[:n])
	}()
	select {
	case b := <-res:
		return b, nil
	case err := <-errc:
		return nil, err
	case <-time.After(d):
		return nil, fmt.Errorf("read timeout on %s", f.Name())
	}
}
