package main

import (
	"log"
	"os"
	"time"

	"github.com/karlpokus/ago"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	if len(os.Args) < 2 {
		log.Fatal("missing time argument")
	}
	t, err := time.Parse(time.RFC3339, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ago.ParseWithContext(t))
}
