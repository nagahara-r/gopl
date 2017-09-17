# words と main.go を並行に圧縮
go run -ldflags -s main.go /usr/share/dict/words main.go
