package cipher

import (
	"NKNshellgo/config"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/tjfoc/gmsm/v2/sm4"
)

//decode脚本待补充
func CipherEncode(plainText string, encryptType string) []byte {
	switch encryptType {
	case "aes":
		result := AES_CBC_Encrypt(plainText, []byte(config.AesKey))
		return result
	default:
		result := Sm4encode(plainText, []byte(config.Sm4Key))
		return result
	}
}

//解密脚本
func CipherDecode(cipherText []byte, decryptType string) string {
	switch decryptType {
	case "aes":
		result := AES_CBC_Decrypt(cipherText, []byte(config.AesKey))
		return string(result)
	default:
		result := Sm4decode(cipherText, []byte(config.Sm4Key))
		return result
	}
}

//国密sm4加密
func Sm4encode(plainText string, key []byte) []byte {
	cipherText, _ := sm4.Sm4Cbc(key, []byte(plainText), true)
	return cipherText
}

//国密sm4解密
func Sm4decode(cipherText []byte, key []byte) string {
	plainText, _ := sm4.Sm4Cbc(key, cipherText, false)
	result := string(plainText)
	return result
}

//填充长度
func Padding(plainText []byte, blockSize int) []byte {
	//计算要填充的长度
	n := blockSize - len(plainText)%blockSize
	//对原来的明文填充n个n
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

//对密文删除填充
func UnPadding(cipherText []byte) []byte {
	//取出密文最后一个字节end
	end := cipherText[len(cipherText)-1]
	//删除填充
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

//AEC加密（CBC模式）
func AES_CBC_Encrypt(plain string, key []byte) []byte {
	//指定加密算法，返回一个AES算法的Block接口对象
	plainText := []byte(plain)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//进行填充
	plainText = Padding(plainText, block.BlockSize())
	//指定初始向量vi,长度和block的块尺寸一致
	iv := []byte("12345678abcdefgh")
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密连续数据库
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	//返回密文
	return cipherText
}

//AEC解密（CBC模式）
func AES_CBC_Decrypt(cipherText []byte, key []byte) []byte {
	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//指定初始化向量IV,和加密的一致
	iv := []byte("12345678abcdefgh")
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//解密
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	//删除填充
	plainText = UnPadding(plainText)
	return plainText
}

//func main() {
//	cipher := CipherEncode("A1andNS", "default")
//	plain := CipherDecode(cipher, "default")
//	fmt.Println(plain)
//}
