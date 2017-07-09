go run main.go

# $ go run main.go -r 100
# Processing Time: 42.669µs
# $ go run main.go -r 1000
# Processing Time: 339.471µs
# $ go run main.go -r 10000
# Processing Time: 4.666366ms
# $ go run main.go -r 100000
# Processing Time: 44.991284ms
# $ go run main.go -r 1000000
# Processing Time: 596.351369ms
# $ go run main.go -r 10000000
# Processing Time: 1m9.766839674s
# $ go run main.go -r 20000000
# Processing Time: 12m3.118440383s
# $ go run main.go -r 25000000
# signal: killed
# 2000万のGoroutineが作れてしまった。
# それ以上になると、物理メモリの限界までゴルーチンを作ってその後はOSに止められているようです。
