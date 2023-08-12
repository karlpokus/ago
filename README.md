# ago
Parse duration into a common time string like 5s, 10m. Can be run as a cli or a go package.

- cli

````sh
$ go install github.com/karlpokus/ago/cmd/ago@latest
# pass time in RFC3339
$ ago $(date -d "-5 min" -Is)
5m
````

- package

See example in `cmd/ago`

# test

````sh
$ go test
````
