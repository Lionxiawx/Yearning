package dingtalk

type DingTalkHeader struct {
	AppCode     string `json:"appCode"`
	Version     string `json:"version"`
	EncryptType string `json:"encryptType"`
}

type DingTalkBody struct {
	Message    Message   `json:"message"`
	Title      string    `json:"title"`
	BusinessNo string    `json:"businessNo"`
	SendKeys   []SendKey `json:"sendKeys"`
}

type SendKey struct {
	SendKeyType string `json:"sendKeyType"`
	SendKey     string `json:"sendKey"`
}

type Message struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	MessageType string `json:"messageType"`
}

//定义属性必须大写
type DingTalkMessage struct {
	Body   string         `json:"body"`
	Header DingTalkHeader `json:"header"`
}
