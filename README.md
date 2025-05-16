# query-compensation-data
A read-only API (REST or GraphAPI) that returns one or more records from a static set of compensation data.

Instructions for setup
1. `make compose-up`
2. `go mod tidy`
3. I have used csv3. To read the records from the csv file and to write to sql table: run
    `go run cmd/ingest/main.go`
