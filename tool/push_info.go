package tool

type PushInfo struct {
	Aps *Apns `json:"aps"` //通知标题
	Payload  string `json:"body"`  //通知内容
	Multi_media []*MultiMedia `json:"multimedia"` //（用于多语言支持）指定执行按钮所使用的Localizable.strings
}

func GetPushInfo() *PushInfo {
	pushInfo := &PushInfo{}
	return pushInfo
}

func (this *PushInfo) SetAps(aps *Apns) {
	this.Aps = aps
}

func (this *PushInfo) SetMultimedia(multimedia *MultiMedia) {
	this.Multi_media = append(this.Multi_media, multimedia)
}
