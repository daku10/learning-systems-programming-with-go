package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	var conn net.Conn = nil
	var err error
	requests := make([]*http.Request, 0, len(sendMessages))

	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Access\n")
	defer conn.Close()

	// リクエストだけ先に送る
	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages) - 1
		request, err := http.NewRequest(
			"GET",
			"http://localhost:8888?message=" + sendMessages[i],
			nil,
		)
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessages[i])
		requests = append(requests, request)
	}

	// レスポンスをまとめて受信
	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}

	// 以下 chunk版
	// conn, err := net.Dial("tcp", "localhost:8888")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()
	// request, err := http.NewRequest(
	// 	"GET",
	// 	"http://localhost:8888",
	// 	nil)
	
	// if err != nil {
	// 	panic(err)
	// }
	// err = request.Write(conn)
	// if err != nil {
	// 	panic(err)
	// }
	// reader := bufio.NewReader(conn)
	// response, err := http.ReadResponse(reader, request)
	// if err != nil {
	// 	panic(err)
	// }
	// dump, err := httputil.DumpResponse(response, false)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(dump))
	// if len(response.TransferEncoding) < 1 ||
	//   response.TransferEncoding[0] != "chunked" {
	// 	  panic("wrong transfer encoding")
	//   }

	// for {
	// 	// サイズを取得
	// 	sizeStr, err := reader.ReadBytes('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	// 16進数のサイズをパース。サイズがゼロならクローズ
	// 	size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
	// 	if size == 0 {
	// 		break
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// サイズ数分バッファを確保して読み込み
	// 	line := make([]byte, int(size))
	// 	io.ReadFull(reader, line)
	// 	// CRLF改行を捨てているのね(ちゃんとやるなら改行かどうかを見ないといけないんだろうなぁ...)
	// 	reader.Discard(2)
	// 	fmt.Printf("    %d bytes: %s\n", size, string(line))
	// }

	// 以下keep-alive版
	// sendMessages := []string{
	// 	"ASCII",
	// 	"PROGRAMMING",
	// 	"PLUS",
	// }
	// current := 0
	// var conn net.Conn = nil

	// // リトライ用にループで全体を囲う
	// for {
	// 	var err error
	// 	// まだコネクションを張っていない / エラーでリトライ
	// 	if conn == nil {
	// 		// Dial から行って conn を初期化
	// 		conn, err = net.Dial("tcp", "localhost:8888")
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		fmt.Printf("Access: %d\n", current)
	// 	}
	// 	// POSTで文字列を送るリクエストを作成
	// 	request, err := http.NewRequest(
	// 		"POST",
	// 		"http://localhost:8888",
	// 		strings.NewReader(sendMessages[current]),
	// 	)
	// 	request.Header.Set("Accept-Encoding", "gzip")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	err = request.Write(conn)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
	// 	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	// 	if err != nil {
	// 		fmt.Println("Retry")
	// 		conn = nil
	// 		continue
	// 	}
	// 	// 結果を表示
	// 	dump, err := httputil.DumpResponse(response, false)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(dump))

	// 	defer response.Body.Close()

	// 	if response.Header.Get("Content-Encoding") == "gzip" {
	// 		reader, err := gzip.NewReader(response.Body)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		io.Copy(os.Stdout, reader)
	// 	} else {
	// 		io.Copy(os.Stdout, response.Body)
	// 	}

	// 	// 全部送信完了していれば終了
	// 	current++
	// 	if current == len(sendMessages) {
	// 		break
	// 	}
	// }
	// conn.Close()
}
