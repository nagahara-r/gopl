# -f jpg, png, gif のいずれかを選択すると変換します。
# 入力形式は自動判別します。
echo original.png to convert.gif
go run main.go -f gif < original.png > convert.gif
file convert.gif
echo
echo convert.gif to convert.jpg
go run main.go -f jpg < convert.gif > convert.jpg
file convert.jpg
echo
echo convert.jpg to convert.png
go run main.go -f png < convert.jpg > convert.png
file convert.png
