TZ=US/Eastern go run clock/worldclock.go &
TZ=Asia/Tokyo go run clock/worldclock.go -port 8010 &
TZ=Europe/London go run clock/worldclock.go -port 8020 &
go run clockwall.go NewYork=localhost:8000 Tokyo=localhost:8010 London=localhost:8020
killall worldclock
