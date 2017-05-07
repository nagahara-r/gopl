go build -o issue main.go
echo 下記を実行可能です。
echo Issue 作成
echo ./issue -c
echo Issue 読み出し
echo ./issue -r -n [issue No]
echo Issue アップデート
echo ./issue -u -n [issue No]
echo Issue クローズ
echo ./issue -cl -n [issue No]
