go run main.go &
go build gopl.io/ch1/fetch
./fetch "http://localhost:12345/search?zcode=1234567&mail=test@example.com&credit=1234567890123456"
./fetch "http://localhost:12345/search?zcode=123456"
./fetch "http://localhost:12345/search?zcode=1234567&mail=test@example.com"
./fetch "http://localhost:12345/search?zcode=1234567&mail=testexample"
./fetch "http://localhost:12345/search?zcode=1234567&mail=test@example.com&credit=123456789012345a"
