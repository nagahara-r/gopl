// Copyright © 2017 Yuki Nagahara
// 練習8-2: FTPサーバの実装

package ftpd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// User はFTPユーザ名を表す構造体です
type User struct {
	User     string
	Password string
}

type client struct {
	user         User
	message      string
	writer       io.Writer
	reader       io.Reader
	isPut        bool
	currentDir   string
	rootDir      string
	filename     string
	dataserverch chan bool
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

		331: "331 Enter Password.\r\n",
		350: "350 File exists. Ready to Move.\r\n",

		// Errors
		500: "500 Command not understood.\r\n",
		501: "501 Parameters or Auguments Parse Error.\r\n",
		502: "502 Command not Implemented.\r\n",
		504: "504 Command not Implemented for that Parameter.\r\n",
		530: "530 Failed to Login.\r\n",
		550: "550 Error Due to File Access\r\n",
	}

	commands = map[string]func(*client, net.Conn){
		"USER": user,
		"PASS": pass,
		"SYST": syst,
		"FEAT": nonimpl,
		"TYPE": typef,
		"MODE": mode,
		"STRU": stru,
		"CWD":  cwd,
		"CDUP": cdup,
		"PWD":  pwd,
		"EPSV": nonimpl,
		"PASV": pasv,
		"PORT": port,
		"LIST": list,
		"NLST": nlst,
		"SIZE": size,
		"RETR": retr,
		"STOR": stor,
		"RNFR": rnfr,
		"RNTO": rnto,
		"DELE": dele,
		"RMD":  dele,
		"MKD":  mkd,
		"MDTM": mdtm,
		"SITE": chmod,
		"NOOP": noop,
		"QUIT": quit,
	}

	dataservermutex = new(sync.Mutex)
)

// HandleConn は mainのServerSocketを受け取り、FTPサーバ処理を開始します。
func HandleConn(conn net.Conn, user User) {
	cli := new(client)

	// currentDirは自由に決められるようにする
	cli.user = user
	cli.currentDir, _ = os.Getwd()
	cli.currentDir = cli.currentDir + "/ftpdir"
	cli.rootDir = cli.currentDir
	cli.dataserverch = make(chan bool)

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
		cm := regexp.MustCompile("([A-Za-z]+)").FindString(messagesp[0])
		f, ok := commands[strings.ToUpper(cm)]
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

func user(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")
	if messages[1] != cli.user.User {
		sendString(statuses[530], conn)
		conn.Close()
		return
	}
	sendString(statuses[331], conn)
}

func pass(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")
	if messages[1] != cli.user.Password {
		sendString(statuses[530], conn)
		conn.Close()
		return
	}
	sendString(statuses[230], conn)
}

func syst(message *client, conn net.Conn) {
	sendString(statuses[215], conn)
}

func typef(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")
	// ASCIIモードの場合、サーバ側の改行コードに変換する必要があるが、
	// OS側でたいてい対応しているので、とくに変換対応していない。
	if messages[1] == "A" || messages[1] == "I" {
		sendString(statuses[200], conn)
		return
	}

	sendString(statuses[504], conn)
}

func mode(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// Streamモードのみサポート
	if messages[1] == "S" {
		sendString(statuses[200], conn)
		return
	}

	sendString(statuses[504], conn)
}

func stru(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// Fileモードのみサポート
	if messages[1] == "F" {
		sendString(statuses[200], conn)
		return
	}

	sendString(statuses[504], conn)
}

func nonimpl(cli *client, conn net.Conn) {
	sendString(statuses[502], conn)
}

func cwd(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を読めるか？
	dpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	_, err = os.Stat(dpath)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// 読めたら移動して正常系
	cli.currentDir = dpath
	sendString(statuses[250], conn)
}

func cdup(cli *client, conn net.Conn) {
	cli.message = "CWD .."
	cwd(cli, conn)
}

func pwd(cli *client, conn net.Conn) {
	dir := strings.TrimPrefix(cli.currentDir, cli.rootDir)
	if dir == "" {
		dir = "/"
	}
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

	go dataTransferServer(port, cli)
	<-cli.dataserverch

	status := fmt.Sprintf("227 Entering Passive Mode (%v,%v,%v)\r\n", strings.Replace(ftpaddrs[0], ".", ",", -1), strconv.Itoa(port/256), strconv.Itoa(port%256))
	sendString(status, conn)
}

func port(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")
	ipport := strings.Split(messages[1], ",")
	ip := strings.Join(ipport[:4], ".")
	port1, err := strconv.Atoi(ipport[4])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[501], conn)
	}
	port2, err := strconv.Atoi(ipport[5])
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[501], conn)
	}
	port := (port1 * 256) + port2

	go dataTransferActive(fmt.Sprintf("%v:%v", ip, port), cli)
	<-cli.dataserverch
	sendString(statuses[200], conn)
}

func list(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")
	messages = append(messages, "")

	patharg := 1
	all := false
	for ; strings.Index(messages[patharg], "-") == 0; patharg++ {
		// All かどうか。
		if strings.Contains(messages[patharg], "a") {
			all = true
		}
	}

	sendString(statuses[150], conn)

	// ファイルの一覧を作成
	dpath, err := getPath(messages[patharg], cli)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに失敗送信
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}

	fileinfos, err := ioutil.ReadDir(dpath)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに失敗送信
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}

	message := ""
	if all {
		curdir, err := os.Stat(cli.currentDir)
		if err != nil {
			log.Printf("%v", err)
			// データサーバに失敗送信
			cli.dataserverch <- false
			sendString(statuses[550], conn)
			return
		}
		message += fmt.Sprintf("%v 0 owner %v %v %v\r\n", curdir.Mode().String(), curdir.Size(), curdir.ModTime().Format("Jan 2 15:04"), ".")

		if cli.currentDir != cli.rootDir {
			rootdir, err := os.Stat(cli.rootDir)
			if err != nil {
				log.Printf("%v", err)
				// データサーバに失敗送信
				cli.dataserverch <- false
				sendString(statuses[550], conn)
				return
			}
			message += fmt.Sprintf("%v 0 owner %v %v %v\r\n", rootdir.Mode().String(), rootdir.Size(), rootdir.ModTime().Format("Jan 2 15:04"), "..")
		}
	}
	for _, fileinfo := range fileinfos {
		message += fmt.Sprintf("%v 0 owner %v %v %v\r\n", fileinfo.Mode().String(), fileinfo.Size(), fileinfo.ModTime().Format("Jan 2 15:04"), fileinfo.Name())
	}
	cli.reader = strings.NewReader(message)

	// データサーバに送信
	cli.dataserverch <- true

	// データサーバから受信待ち
	<-cli.dataserverch

	sendString(statuses[226], conn)
}

func nlst(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")
	messages = append(messages, "")
	sendString(statuses[150], conn)

	// ファイルの一覧を作成
	dpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに準備送信通知
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}

	fileinfos, err := ioutil.ReadDir(dpath)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに準備送信通知
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}

	message := ""
	for _, fileinfo := range fileinfos {
		message += fmt.Sprintf("%v\r\n", fileinfo.Name())
	}
	cli.reader = strings.NewReader(message)
	// データサーバに送信
	cli.dataserverch <- true

	// データサーバから受信待ち
	<-cli.dataserverch

	sendString(statuses[226], conn)
}

func size(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	// ファイルの一覧を作成
	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
	}

	fstat, err := os.Stat(fpath)
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

	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに準備送信通知
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}

	file, err := os.Open(fpath)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに準備送信通知
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}
	cli.reader = file

	cli.dataserverch <- true
	// データサーバから受信待ち
	<-cli.dataserverch

	file.Close()
	sendString(statuses[226], conn)
}

func stor(cli *client, conn net.Conn) {
	sendString(statuses[150], conn)
	messages := strings.Split(cli.message, " ")

	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに準備送信通知
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}

	file, err := os.Create(fpath)
	if err != nil {
		log.Printf("%v", err)
		// データサーバに準備送信通知
		cli.dataserverch <- false
		sendString(statuses[550], conn)
		return
	}
	cli.writer = file
	cli.isPut = true
	cli.dataserverch <- true

	// データサーバから受信待ち
	<-cli.dataserverch
	cli.isPut = false

	file.Close()
	sendString(statuses[226], conn)
}

func rnfr(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// ファイルが存在するか？
	_, err = os.Stat(fpath)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	cli.filename = fpath
	sendString(statuses[350], conn)
}

func rnto(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// ファイルリネーム
	err = os.Rename(cli.filename, fpath)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	sendString(statuses[250], conn)
}

func dele(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// ファイル削除
	err = os.Remove(fpath)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	sendString(statuses[250], conn)
}

func mkd(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	dpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// ディレクトリ作成
	err = os.Mkdir(dpath, 0755)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	sendString(statuses[250], conn)
}

func mdtm(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	fpath, err := getPath(messages[1], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	fstat, err := os.Stat(fpath)
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

func chmod(cli *client, conn net.Conn) {
	messages := strings.Split(cli.message, " ")

	mod, err := strconv.ParseInt(messages[2], 8, 32)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	fpath, err := getPath(messages[3], cli)
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}

	// ファイルの一覧を作成
	err = os.Chmod(fpath, os.FileMode(mod))
	if err != nil {
		log.Printf("%v", err)
		sendString(statuses[550], conn)
		return
	}
	sendString(statuses[250], conn)
}

func noop(cli *client, conn net.Conn) {
	sendString(statuses[200], conn)
}

func quit(cli *client, conn net.Conn) {
	sendString(statuses[221], conn)
	conn.Close()
}

func dataTransferServer(port int, cli *client) {
	dataservermutex.Lock()
	defer dataservermutex.Unlock()
	log.Printf("[DataTransferServer(%v)]", port)
	listener, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	// pasv() <-
	cli.dataserverch <- true

	conn, err := listener.Accept()
	if err != nil {
		log.Print(err) // e.g., connection aborted
	}
	listener.Close()

	go dataHandleConn(conn, cli) // handle connections concurrently
}

func dataTransferActive(dst string, cli *client) {
	log.Printf("[DataTransferActive(%v)]", dst)
	conn, err := net.Dial("tcp", dst)
	if err != nil {
		log.Fatal(err)
	}
	cli.dataserverch <- true

	go dataHandleConn(conn, cli) // handle connections concurrently
}

func dataHandleConn(conn net.Conn, cli *client) {
	// list() などでデータの用意待ち
	ready := <-cli.dataserverch

	if !ready {
		// 転送準備に失敗してるのでCloseする。
		conn.Close()
		return
	}

	log.Println("Start Data Transfer")
	if cli.isPut {
		// Caution: コピーのエラーを無視
		io.Copy(cli.writer, conn)
	} else {
		io.Copy(conn, cli.reader)
	}
	conn.Close()
	// 送信終わったので待っているメソッドを終了
	cli.dataserverch <- true
}

func getPath(path string, cli *client) (fpath string, err error) {
	if strings.Index(path, "/") == 0 {
		fpath, err = filepath.Abs(cli.rootDir + path)
	} else {
		fpath, err = filepath.Abs(cli.currentDir + "/" + path)
	}

	if err != nil {
		return
	}

	if len(fpath) < len(cli.rootDir) {
		return fpath, fmt.Errorf("Access Denied")
	}
	return
}
