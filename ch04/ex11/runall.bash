go get golang.org/x/crypto
go build -o issue main.go
echo 下記を実行可能です。
echo Issue 作成
echo ./issue -u [USERID] -r [REPOSITORY] -c
echo Issue 読み出し
echo ./issue -u [USERID] -r [REPOSITORY] -re -n [ISSUE NO]
echo Issue 編集
echo ./issue -u [USERID] -r [REPOSITORY] -e -n [ISSUE NO]
echo Issue クローズ
echo ./issue -u [USERID] -r [REPOSITORY] -cl -n [ISSUE NO]
