# advent-of-code

## Run

```sh
# for the years using Go
# example: go run -C ./2019/02 .
go run -C [path] .

# for the years using Python
# example: uv run ./2015/01/solution.py
uv run [path]
```

## Test

Testing a single day

```sh
# for the years using Go
# example: go test ./2019/02 
go test [path]

# for the years using Python
# example: uv run -m unittest ./2015/01/solution_test.py
uv run -m unittest [path]
```

Testing all days

```sh
# for the years using Go
go test -timeout 30s ./...

# for the years using Python
uv run -m unittest discover -p "*_test.py"
```
