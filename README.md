# Go directory tree / Приложение для вывода дерева файлов
****
### RUN
1. `go run main.go` `path (form main.go)` `with files (-f)`
****
### BUILD and RUN
1. `go build main.go`
2. `main.extenstion` `path (form main.go)` `with files (-f)`
****
### TESTS:
###### MAIN
1. `go test -v -coverprofile` `cover_cover.out` `main.go` `main_test.go`
2. `go tool cover --html=cover_cover.out`
