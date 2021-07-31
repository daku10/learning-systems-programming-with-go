package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 順番に従ってconnに書き出しをする(goroutineで実行される)
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}

// セッション内のリクエストを処理する
func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("HOAAA")
	fmt.Println(string(dump))
	message := request.URL.Query().Get("message")
	content := message
	// レスポンスを書き込む
	// セッションを維持するためにKeep-Aliveでないといけない
	response := &http.Response {
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		ContentLength: int64(len(content)),
		Body: ioutil.NopCloser(strings.NewReader(content)),
	}
	// 処理が終わっていたらチャネルに書き込み
	// ブロックされていたwriteToConnの処理を再始動する
	resultReceiver <- response	
}


// クライアントはgzipを受け入れ可能か？
func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

var contents = []string{
	"これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。",
	"むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、",
	"中山さまというおとのさまが、おられたそうです。",
	"その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。",
	"ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。",
	"そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。",
}

// 1セッションの処理をする
func processSession(conn net.Conn) {

	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// セッション内のリクエストを順に処理するためのチャネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	// レスポンスを直列化してソケットに書き出す専用のgoroutine
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)
	for {
		// レスポンスを受け取ってセッションのキューに入れる
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読み込む
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse
		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResponse)
	}

	// fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// defer conn.Close()
	// // このforいる？ keep-alive版の名残かな
	// // 動かしてみた感じ、requestとEOFで終わるのばかりだったので
	// //　クライアントからの EOF で切るのが行儀がいいからかな？と思った
	// for {
	// 	// リクエストを読み込む
	// 	request, err := http.ReadRequest(bufio.NewReader(conn))
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		panic(err)
	// 	}
	// 	dump, err := httputil.DumpRequest(request, true)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(dump))

	// 	// レスポンスを書き込む
	// 	fmt.Fprintf(conn, strings.Join([]string{
	// 		"HTTP/1.1 200 OK",
	// 		"Content-Type: text/plain",
	// 		"Transfer-Encoding: chunked",
	// 		"", "",
	// 	}, "\r\n"))
	// 	for _, content := range contents {
	// 		bytes := []byte(content)
	// 		fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
	// 	}
	// 	fmt.Fprintf(conn, "0\r\n\r\n")
	// }

	// 以下 Keep-Alive版

	// fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// defer conn.Close()

	// for {
	// 	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	// 	// リクエストを読み込む
	// 	request, err := http.ReadRequest(bufio.NewReader(conn))
	// 	if err != nil {
	// 		neterr, ok := err.(net.Error)
	// 		if ok && neterr.Timeout() {
	// 			fmt.Println("Timeout")
	// 			break
	// 		} else if err == io.EOF {
	// 			break
	// 		}
	// 		panic(err)
	// 	}
	// 	dump, err := httputil.DumpRequest(request, true)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(dump))
	// 	// レスポンスを書き込む
	// 	response := http.Response {
	// 		StatusCode: 200,
	// 		ProtoMajor: 1,
	// 		ProtoMinor: 1,
	// 		Header: make(http.Header),
	// 	}
	// 	if isGZipAcceptable(request) {
	// 		content := "Hello World (gzipped)\n"
	// 		// コンテンツをgzip化して転送
	// 		var buffer bytes.Buffer
	// 		writer := gzip.NewWriter(&buffer)
	// 		io.WriteString(writer, content)
	// 		writer.Close()
	// 		response.Body = ioutil.NopCloser(&buffer)
	// 		response.ContentLength = int64(buffer.Len())
	// 		response.Header.Set("Content-Encoding", "gzip")
	// 	} else {
	// 		content := "Hello World\n"
	// 		response.Body = ioutil.NopCloser(strings.NewReader(content))
	// 		response.ContentLength = int64(len(content))
	// 	}
	// 	response.Write(conn)
	// }
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}

	// tcp client
	// conn, err := net.Dial("tcp", "localhost:8080")
	// if err != nil {
	// 	panic(err)
	// }
	// conn.Write([]byte("hogehoge"))

	// // tcp server
	// ln, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	panic(err)
	// }
	// serverConn, err := ln.Accept()
	// if err != nil {
	// 	panic(err)
	// }
	// serverConn.Write([]byte("fugafuga"))

	// ln, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		// handle error
	// 	}
	// 	go func() {
	// 		conn.Write([]byte("from server!"))
	// 		io.Copy(os.Stdout, conn)
	// 		conn.Close()
	// 	}()
	// }

	// listener, err := net.Listen("tcp", "localhost:8888")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Server is running at localhost:8888")
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	go func() {
	// 		defer conn.Close()
	// 		fmt.Printf("Accept %v]n", conn.RemoteAddr())
	// 		for {
	// 			conn.SetDeadline(time.Now().Add(5 * time.Second))
	// 			request, err := http.ReadRequest(bufio.NewReader(conn))
	// 			if err != nil {
	// 				// タイムアウトもしくはソケットクローズ時は終了
	// 				// それ以外はエラーにする
	// 				neterr, ok := err.(net.Error)
	// 				if ok && neterr.Timeout() {
	// 					fmt.Println("Timeout")
	// 					break
	// 				} else if err == io.EOF {
	// 					break
	// 				}
	// 				panic(err)
	// 			}

	// 			// リクエスト表示
	// 			dump, err := httputil.DumpRequest(request, true)
	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 			fmt.Println(string(dump))

	// 			content := "Hello World\n"

	// 			response := http.Response{
	// 				StatusCode: 200,
	// 				ProtoMajor: 1,
	// 				ProtoMinor: 1,
	// 				ContentLength: int64(len(content)),
	// 				Body: ioutil.NopCloser(strings.NewReader(content)),
	// 			}
	// 			response.Write(conn)
	// 		}
	// 	}()
	// }

}