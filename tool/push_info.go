package tool

type PushInfo struct {
	Aps *Apns `json:"aps"` //通知标题
	Payload  string `json:"payload"`  //通知内容
	Multi_media []*MultiMedia `json:"multimedia"` //（用于多语言支持）指定执行按钮所使用的Localizable.strings
	data map[string]interface{} `json:"data"`
}

func GetPushInfo() *PushInfo {
	pushInfo := &PushInfo{
		data: make(map[string]interface{}),
	}
	return pushInfo
}

func (this *PushInfo) SetAps(aps *Apns) {
	this.Aps = aps
	this.data["aps"] = aps
}

func (this *PushInfo) SetMultimedia(multimedia *MultiMedia) {
	this.Multi_media = append(this.Multi_media, multimedia)
	this.data["multimedia"] = this.Multi_media
}

func (this *PushInfo) SetPayload(payload string) {
	this.Payload = payload
	this.data["payload"] = payload
}


func (this *PushInfo) GetData() map[string]interface{} {
	return this.data
}
