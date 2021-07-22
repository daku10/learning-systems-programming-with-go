package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWrite example")
}

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()

	os.Stdout.Write([]byte("os.Stdout example\n"))

	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())

	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())

	// conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	// io.WriteString(conn, "GET / HTTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	// io.Copy(os.Stdout, conn)

	// req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// req.Write(conn)
	// io.Copy(os.Stdout, conn)
	// conn.Close()

	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	file, err = os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")

	file, err = os.Create("text.txt.gz")
	if err != nil {
		panic(err)
	}
	gzipWriter := gzip.NewWriter(file)
	gzipWriter.Header.Name = "test.txt"
	io.WriteString(gzipWriter, "gzip.Writer example\n")
	gzipWriter.Close()

	bufioBuffer := bufio.NewWriter(os.Stdout)
	bufioBuffer.WriteString("bufio.Writer ")
	bufioBuffer.Flush()
	bufioBuffer.WriteString("example\n")
	bufioBuffer.Flush()

	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})

	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)

	q1File, err := os.Create("q1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(q1File, "fomatted integer: %d, strings: %s, floating number: %f", 1, "Q2.1 answer", 5.1)

	csvWriter := csv.NewWriter(os.Stdout)
	csvWriter.Write([]string{"aaa", "bbb", "ccc"})
	csvWriter.WriteAll([][]string{{"ddd", "eee", "fff"}, {"ggg", "hhh", "iii"}})
	csvWriter.Flush()

	http.HandleFunc("/", q3Handler)
	http.ListenAndServe(":8080", nil)
}

func q3Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(os.Stdout, gzipWriter)
	jsonEncoder := json.NewEncoder(multiWriter)
	jsonEncoder.SetIndent("", "    ")
	err := jsonEncoder.Encode(source)

	if err != nil {
		panic(err)
	}

	gzipWriter.Flush()
}
