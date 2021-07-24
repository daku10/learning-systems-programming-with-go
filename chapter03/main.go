package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break;
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}

	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)

	conn, err := net.Dial("tcp", "ascii.jp:80")
	// if err != nil {
	// 	panic(err)
	// }
	// conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	// io.Copy(os.Stdout, conn)

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

	// var buffer1 bytes.Buffer
	// buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	// buffer3 := bytes.NewBufferString("初期文字列")
	
	// bReader1 := bytes.NewReader([]byte{0x10, 0x20, 0x30})
	// bReader2 := bytes.NewReader([]byte("文字列をバイト配列にキャストして設定"))

	sReader := strings.NewReader("Readerの出力内容は文字列で渡す")
	io.Copy(os.Stdout, sReader)

	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
	// 文字列の切り取りの場合は substr 的なのがあるんだろうなぁ

	// 32ビットのビッグエンディアンのデータ
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)

	pngFile, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()
	chunks := readChunks(pngFile)

	// ここでdumpするとreaderが読み込まれて、続くコードの書き込み時に読めなくなるのでコメントアウト
	// for _, chunk := range chunks {
	// 	dumpChunk(chunk)
	// }

	newFile, err := os.Create("Lenna2.png")
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))

	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	strReader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := strReader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

	fReader := strings.NewReader(fSource)
	var inum int
	var f, g float64
	var s string
	fmt.Fscan(fReader, &inum, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", inum, f, g, s)

	cReader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(cReader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}

	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	multiReader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, multiReader)

	var tBuffer bytes.Buffer
	tReader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(tReader, &tBuffer)
	// データを読み捨てる
	// _, _ = ioutil.ReadAll(teeReader)
	_, _ = teeReader.Read(make([]byte, 5))

	// けどバッファに残っている(というか読み込まれた分が同時に書き込まれている ↑で５バイトだけ読み込んだらその分だけ書き込まれている)
	fmt.Println(tBuffer.String())

	q3_1()
	q3_2()

	q3_3_pre()
	q3_3()

	// ブロッキングするからコメントアウト
	// http.HandleFunc("/", q3_4_handler)
	// http.ListenAndServe(":8080", nil)

	q3_5()

	q3_6()
}

func q3_1() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(newFile, file)
}

func q3_2() {
	// こんな感じ？ 何かもっと良い書き方ある気がするけど
	buf := make([]byte, 1024)
	rand.Reader.Read(buf)
	// ioutil.WriteFile, os.WriteFile は truncateして書き込むのでこっちのほうが簡単そう
	os.WriteFile("randomBytes", buf, os.ModeAppend)
	// file, err := os.Create("randomBytes")
	// if err != nil {
	// 	panic(err)
	// }
	// file.Write(buf)
}

func q3_3_pre() {
	// zipファイルの練習
	file, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	writer, _ := zipWriter.Create("new-zip.txt")
	writer.Write([]byte("hogehogehoge"))
}

func q3_3() {
	// そもそも問題文の意味が分からん？？？
	file, err := os.Create("new3.zip")
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	strReader := strings.NewReader("hogehoge")
	writer, err := zipWriter.Create("new3-zip.txt")
	if err != nil {
		panic(err)
	}
	strReader.WriteTo(writer)
}

func q3_4_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachiment; filename=ascii_sample.zip")
	zipWriter := zip.NewWriter(w)	
	writer, err := zipWriter.Create("archived.txt")
	if err != nil {
		panic(err)
	}
	defer zipWriter.Close()
	// この書き方楽かな？いっぱい方法があるからうーん。。。[]byte("this is zip") でもできる気がする
	writer.Write(bytes.NewBufferString("this is zip?").Bytes())
}

func q3_5() {
	q3_5_copyN(os.Stdout, strings.NewReader("hogehogehoge"), 3)
}

func q3_5_copyN(dest io.Writer, src io.Reader, length int) {
	buf := make([]byte, length)
	src.Read(buf)
	dest.Write(buf)
}

func q3_6() {
	var stream io.Reader

	// ここでASCII を作る
	// A from PROGR A MMING
	// S from S YSTEM
	// C from C OMPUTER
	// I from PROGRAMM I NG
	// I from PROGRAMM I NG
	
	// とりあえず動いたけど全然面白くない作り方な気がする
	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(system, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i1 := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(a, s, c, i1, i2)

	io.Copy(os.Stdout, stream)
}

var (
	computer = strings.NewReader("COMPUTER")
	system = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

var source = `１行目
２行目
３行目 
`

var fSource = "123 1.234 1.0e4 test"

var csvSource = 
`13101,"100  ","1000003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(1ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（１丁目）",1,0,1,0,0,0
13101,"101  ","1010003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(2ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（２丁目）",1,0,1,0,0,0
13101,"100  ","1000012","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾋﾞﾔｺｳｴﾝ","東京都","千代田区","日比谷公園",0,0,0,0,0,0
13101,"102  ","1020093","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾗｶﾜﾁｮｳ","東京都","千代田区","平河町",0,0,1,0,0,0
13101,"102  ","1020071","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾌｼﾞﾐ","東京都","千代田区","富士見",0,0,1,0,0,0
`

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	crc.Write(byteData)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)]\n", string(buffer), length)
}