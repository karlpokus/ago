# ago
Parse duration into a common time string like 5s, 10m.

Can be run as a cli or included as a go package (example in `cmd/ago`).

````sh
# install cli
$ go install ./cmd/ago
# pass time in RFC3339
$ ago $(date -d "-5 min" -Is)
5m ago
# or by stdin
$ date -Is -d "+5 min" | ago -
in 5m
````

# test

````sh
$ make test
````
