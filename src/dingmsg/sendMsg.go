package dingtalk

import (
	"Yearning-go/src/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
发送审核消息
*/
func SendApprovalMsg(order model.CoreSqlOrder) {
	//构造body
	sendKey := []SendKey{
		{"LDAP",
			order.Assigned},
	}

	content := "SQL审计通知 \n\n> 工单编号:" + order.WorkId + " \n \n> 申&nbsp;&nbsp;请&nbsp;&nbsp;人:" + order.Username + " \n \n>地&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;址: " + model.DingMsgWebUrl + " \n\n <font color=#FF0000 >请审批!</font>"
	message := Message{
		"SQL审计",
		content,
		"MARKDOWN",
	}
	dingTalkBody := &DingTalkBody{
		Message:    message,
		Title:      "SQL审计",
		BusinessNo: order.WorkId,
		SendKeys:   sendKey,
	}
	body, _ := json.Marshal(dingTalkBody)
	//println(string(body))

	//加密body
	iv := "0000000000000000"
	aes := aesTool(model.DingMsgAppSecret, iv)
	//fmt.Println("加密前的明文：" + string(body))
	cipherText, _ := aes.encrypt(string(body))

	//构造请求体
	dingTalkMessage := &DingTalkMessage{
		Body: cipherText,
		Header: DingTalkHeader{
			AppCode:     model.DingMsgAppCode,
			Version:     "1.0.0",
			EncryptType: "AES",
		},
	}
	bs, _ := json.Marshal(dingTalkMessage)

	data := bytes.NewReader(bs)
	response, err := http.Post(model.DingMsgUrl+"/api/sendNotice",
		"application/json", data)
	if err != nil {
		println(err)
	}
	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err)
	}
	println(string(resp))

}

/**
发送执行消息
*/
func SendExecMsg(order model.CoreSqlOrder) {
	//构造body
	sendKey := []SendKey{
		{"LDAP",
			order.Username},
	}
	content := "SQL审计通知 \n\n> 工单编号:" + order.WorkId + "   \n\n <font color=#008000 >已执行!</font>"
	message := Message{
		"SQL审计",
		content,
		"MARKDOWN",
	}
	dingTalkBody := &DingTalkBody{
		Message:    message,
		Title:      "SQL审计",
		BusinessNo: order.WorkId,
		SendKeys:   sendKey,
	}
	body, _ := json.Marshal(dingTalkBody)
	//println(string(body))

	//加密body
	iv := "0000000000000000"
	aes := aesTool(model.DingMsgAppSecret, iv)
	//fmt.Println("加密前的明文：" + string(body))
	cipherText, _ := aes.encrypt(string(body))

	//构造请求体
	dingTalkMessage := &DingTalkMessage{
		Body: cipherText,
		Header: DingTalkHeader{
			AppCode:     model.DingMsgAppCode,
			Version:     "1.0.0",
			EncryptType: "AES",
		},
	}
	bs, _ := json.Marshal(dingTalkMessage)

	data := bytes.NewReader(bs)
	response, err := http.Post(model.DingMsgUrl+"/api/sendNotice",
		"application/json", data)
	if err != nil {
		println(err)
	}
	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err)
	}
	println(string(resp))

}
