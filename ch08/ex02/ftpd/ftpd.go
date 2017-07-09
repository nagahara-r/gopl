// Copyright © 2017 Yuki Nagahara
// 練習8-2: FTPサーバの実装

// 現状の制約リスト：
// ログイン機能はありません（何を入力しても匿名ユーザとしてログイン）
// ディレクトリ指定できません
// パーミッション設定できません
// ディレクトリ作成できません

package ftpd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type client struct {
	message      string
	writer       io.Writer
	reader       io.Reader
	isPut        bool
	currentDir   string
	rootDir      string
	dataserverch chan struct{}
}

var (
	statuses = map[int]string{
		// Standards
		150: "150 Opening data channel.\r\n",
		200: "200 Command Success.\r\n",
		// 213 は　数値 なので各種を求めて作る
		215: "215 Golang\r\n",
		220: "220 FTP Service is ready.\r\n",
		221: "221 Good Bye.\r\n",
		226: "226 Transfer Complete.\r\n",
		// 227はIPアドレスを含むので pasv() で用意する
		//227: "227 Entering Passive Mode (127,0,0,1,31,63)\r\n",
		230: "230 User Logged in.\r\n",
		250: "250 File Action is OK.\r\n",
		// 257 は カレントディレクトリを表すのでPWDで作る
		//257: "257 \"/\" is current directory.\r\n",

		// Errors
		500: "500 Command not understood.\r\n",
		501: "501 Parameters or Auguments Parse Error.\r\n",
		502: "502 Command Not Implemented.\r\n",
		550: "550 Error Due to File Access\r\n",
	}

	commands = map[string]func(*client, net.Conn){
		"USER": user,
		"SYST": syst,
		"FEAT": nonimpl,
		"CWD":  cwd,
		"PWD":  pwd,
		"EPSV": nonimpl,
		"PASV": pasv,
		"PORT": port,
		"LIST": list,
		"NLST": nlst,
		"SIZE": size,
		"RETR": retr,
		"STOR": stor,
		"MDTM": mdtm,
		"QUIT": quit,
	}

	dataservermutex = new(sync.Mutex)
)

// HandleConn は mainのServerSocketを受け取り、FTPサーバ処理を開始します。
func HandleConn(conn net.Conn) {
	cli := new(client)

	// currentDirは自由に決められるようにする
	cli.currentDir, _ = os.Getwd()
	cli.currentDir = cli.currentDir + "/ftpdir"
	cli.rootDir = cli.currentDir

	sendString(statuses[220], conn)

	for {
		message, err := readCommand(conn)
		if err != nil {
			log.Printf("read() %v, disconnected.\n", err)
			conn.Close()
			return
		}

		messagesp := strings.Split(message, "\r\n")
		log.Printf("[read()] %v", messagesp[0])
		cm := regexp.MustCompile("([A-Z]+)").FindString(messagesp[0])
		f, ok := commands[cm]
		if !ok {
			// Non Implemented Command
			nonimpl(cli, conn)
			continue
		}
		cli.message = messagesp[0]
		f(cli, conn)
	}
}

func sendString(message string, conn net.Conn) {
	log.Printf("[send()] %v", message)
	_, err := io.WriteString(conn, message)
	if err != nil {
		log.Printf("err = %v", err)
		return // e.g., client disconnected
	}
}

func readCommand(conn net.Conn) (str string, err error) {
	buf := make([]byte, 128)

	for n := 0; n == 0; {
		n, err = conn.Read(buf)
	}

	if err != nil {
		return // e.g., client disconnected
	}
	str = string(buf)
	return
}

func user(message *client, conn net.Conn) {
	// TODO: ユーザログイン
	sendString(statuses[230], conn)
}

func syst(message *client, conn net.Conn) {
	sendString(statuses[215], conn)
}

func nonimpl(cli *client, conn net.Conn) {
	sendString(statuses[502], conn)
}

func cwd(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を読めるか？
	_, err := ioutil.ReadDir(cli.currentDir + "/" + messages[1])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	// 読めたら移動して正常系
	cli.currentDir = cli.currentDir + "/" + messages[1]
	sendString(statuses[250], conn)
}

func pwd(cli *client, conn net.Conn) {
	dir := strings.TrimPrefix(cli.currentDir, cli.rootDir) + "/"
	message := "257 \"" + dir + "\" is current directory.\r\n"
	sendString(message, conn)
}

func pasv(cli *client, conn net.Conn) {
	// サーバ準備
	ftpaddr := conn.LocalAddr().String()
	ftpaddrs := strings.Split(ftpaddr, ":")

	port, err := strconv.Atoi(ftpaddrs[1])
	port--
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[501], conn)
		return
	}

	cli.dataserverch = make(chan struct{})
	go dataTransferServer(port, cli)
	<-cli.dataserverch

	status := fmt.Sprintf("227 Entering Passive Mode (%v,%v,%v)\r\n", strings.Replace(ftpaddrs[0], ".", ",", -1), strconv.Itoa(port/256), strconv.Itoa(port%256))
	sendString(status, conn)
}

func port(cli *client, conn net.Conn) {
	sendString(statuses[200], conn)
}

func list(cli *client, conn net.Conn) {
	sendString(statuses[150], conn)

	// ファイルの一覧を作成
	fileinfos, err := ioutil.ReadDir(cli.currentDir)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
	}

	message := ""
	for _, fileinfo := range fileinfos {
		message += fmt.Sprintf("%v\t%v\t%v\t%v\r\n", fileinfo.Mode().String(), fileinfo.Size(), fileinfo.ModTime(), fileinfo.Name())
	}
	cli.reader = strings.NewReader(message)

	// データサーバに送信
	cli.dataserverch <- struct{}{}

	// データサーバから受信待ち
	<-cli.dataserverch

	sendString(statuses[226], conn)
}

func nlst(cli *client, conn net.Conn) {
	sendString(statuses[150], conn)

	// ファイルの一覧を作成
	fileinfos, err := ioutil.ReadDir(cli.currentDir)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
	}

	message := ""
	for _, fileinfo := range fileinfos {
		message += fmt.Sprintf("%v\r\n", fileinfo.Name())
	}
	cli.reader = strings.NewReader(message)
	// データサーバに送信
	cli.dataserverch <- struct{}{}

	// データサーバから受信待ち
	<-cli.dataserverch

	sendString(statuses[226], conn)
}

func size(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を作成
	fstat, err := os.Stat(cli.currentDir + "/" + messages[1])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
	}

	message := fmt.Sprintf("213 %v\r\n", fstat.Size())
	sendString(message, conn)
}

func retr(cli *client, conn net.Conn) {
	sendString(statuses[150], conn)
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を作成
	file, err := os.Open(cli.currentDir + "/" + messages[1])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	cli.reader = file

	cli.dataserverch <- struct{}{}
	// データサーバから受信待ち
	<-cli.dataserverch

	file.Close()
	sendString(statuses[226], conn)
}

func stor(cli *client, conn net.Conn) {
	sendString(statuses[150], conn)
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を作成
	file, err := os.Create(cli.currentDir + "/" + messages[1])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	cli.writer = file
	cli.isPut = true
	cli.dataserverch <- struct{}{}

	// データサーバから受信待ち
	<-cli.dataserverch
	cli.isPut = false

	file.Close()
	sendString(statuses[226], conn)
}

func mdtm(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を作成
	fstat, err := os.Stat(cli.currentDir + "/" + messages[1])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// そのままだと日本時間が取れて向こうでずれる。UTCにする。
	format := fstat.ModTime().In(time.UTC).Format("20060102150405")

	message := fmt.Sprintf("213 %v\r\n", format)
	sendString(message, conn)
}

func quit(cli *client, conn net.Conn) {
	sendString(statuses[221], conn)
	conn.Close()
}

func dataTransferServer(port int, cli *client) {
	dataservermutex.Lock()
	defer dataservermutex.Unlock()
	log.Printf("[DataTransferServer(%v)]", port)
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	// pasv() <-
	cli.dataserverch <- struct{}{}

	conn, err := listener.Accept()
	if err != nil {
		log.Print(err) // e.g., connection aborted
	}
	listener.Close()

	go dataHandleConn(conn, cli) // handle connections concurrently
}

func dataHandleConn(conn net.Conn, cli *client) {
	// list() などでデータの用意待ち
	<-cli.dataserverch
	if cli.isPut {
		// Caution: コピーのエラーを無視
		io.Copy(cli.writer, conn)
	} else {
		io.Copy(conn, cli.reader)
	}
	conn.Close()
	// 送信終わったので待っているメソッドを終了
	cli.dataserverch <- struct{}{}
}
