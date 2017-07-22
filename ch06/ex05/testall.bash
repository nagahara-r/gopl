echo 32bit
GOARCH=386 go test ./intset
echo 64bit
GOARCH=amd64 go test ./intset
