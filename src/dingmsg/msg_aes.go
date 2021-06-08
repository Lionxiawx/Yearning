package dingtalk

import (
	"bytes"
	basicAES "crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

/**
* run test
 */
func TestAes(t *testing.T) {

	plainText := "{message:{content:121致力于完善大宗商品供应链及物流生态链，打造互联网化的“供应链技术+物流服务+金融场景”的新生态;司机宝改变传统的物流承运概念,建立一站式O2O物流综合服务体系,利用移动互联网、云计算、大数据的特点、提供信息服务、供应链金融技术、数据服务、多式联运及汽车后市场等运营服务支持；司机宝由武汉物易云通网络科技有限公司自主研发，公司成立于2015年6月，总部位于中国光谷武汉东湖高新区、核心团队是国内从事供应链与物流科技市场开发与应用的团队之一。}}"
	securityKey16 := "b48e2a74925e4f6ab4cac988e63bc9c5"
	//securityKey24 := "Skesj(eE%32sLOapA9e2snEw"
	//securityKey32 := "Skesj(eE%32sLOapA9e2snEwEeopsWui"
	iv := "0000000000000000"
	aes := aesTool(securityKey16, iv)
	fmt.Println("加密前的明文：" + plainText)
	cipherText, _ := aes.encrypt(plainText)
	fmt.Println("加密后的密文：" + cipherText)
	outPlainText, _ := aes.decrypt(cipherText)
	fmt.Println("解密后明文：" + outPlainText)
}

type aes struct {
	securityKey []byte
	iv          []byte
}

/**
* constructor
 */
func aesTool(securityKey string, iv string) *aes {
	return &aes{[]byte(securityKey), []byte(iv)}
}

/**
* 加密
* @param string $plainText 明文
* @return bool|string
 */
func (a aes) encrypt(plainText string) (string, error) {
	block, err := basicAES.NewCipher(a.securityKey)
	if err != nil {
		return "", err
	}
	plainTextByte := []byte(plainText)
	blockSize := block.BlockSize()
	plainTextByte = addPKCS7Padding(plainTextByte, blockSize)
	cipherText := make([]byte, len(plainTextByte))
	mode := cipher.NewCBCEncrypter(block, a.iv)
	mode.CryptBlocks(cipherText, plainTextByte)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

/**
* 解密
* @param string $cipherText 密文
* @return bool|string
 */
func (a aes) decrypt(cipherText string) (string, error) {
	block, err := basicAES.NewCipher(a.securityKey)
	if err != nil {
		return "", err
	}
	cipherDecodeText, decodeErr := base64.StdEncoding.DecodeString(cipherText)
	if decodeErr != nil {
		return "", decodeErr
	}
	mode := cipher.NewCBCDecrypter(block, a.iv)
	originCipherText := make([]byte, len(cipherDecodeText))
	mode.CryptBlocks(originCipherText, cipherDecodeText)
	originCipherText = stripPKSC7Padding(originCipherText)
	return string(originCipherText), nil
}

/**
* 填充算法
* @param string $source
* @return string
 */
func addPKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

/**
* 移去填充算法
* @param string $source
* @return string
 */
func stripPKSC7Padding(cipherText []byte) []byte {
	length := len(cipherText)
	unpadding := int(cipherText[length-1])
	return cipherText[:(length - unpadding)]
}
