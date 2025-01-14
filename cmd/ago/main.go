package main

import (
	"log"
	"os"
	"time"

	"github.com/karlpokus/ago"
)

func main() {
	stdout := log.New(os.Stdout, "", 0)
	stderr := log.New(os.Stderr, "", 0)
	if len(os.Args) < 2 {
		stderr.Fatal("error: missing ISO-formatted time argument")
	}
	t, err := time.Parse(time.RFC3339, os.Args[1])
	if err != nil {
		stderr.Fatal(err)
	}
	stdout.Println(ago.ParseWithContext(t))
}
