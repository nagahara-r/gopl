go build gopl.io/ch1/fetch
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run main.go div div h2 sec-intro
rm fetch
