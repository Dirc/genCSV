# List all files in a csv file

## Use-case

We have a directory where we put all my receipts from the renovation of our house.
We want to have an overview of each file and it's value.

GenCSV is a CLI tool which creates or updates a csv file which:

- Contains a list of all files in the directory
- Manually add/update the values to the list
- Count the total of all known values

## Usage

```shell
./csvGen --dir "./example" --csv "files.csv"

./csvGen --help
```

## Build

```shell
go test
golangci-lint run # Golang lint tool: https://github.com/golangci/golangci-lint 
go build
go install

```

## ToDo

- [x] unit tests
- [ ] loglevel
- [x] CLI
- [x] golangci-lint
- [ ] csv content in alphabetic order
- [ ] print files with "0" value at end?
- [ ] remove files?


