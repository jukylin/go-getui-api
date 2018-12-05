package tool

type PushInfo struct {
	Aps *Apns `json:"aps"` //通知标题
	Payload  string `json:"body"`  //通知内容
	Multimedia []MultiMedia `json:"Multimedia"` //（用于多语言支持）指定执行按钮所使用的Localizable.strings
}

func GetPushInfo() *Alert {
	alert := &PushInfo{}
	return alert
}

func (this *PushInfo) SetAps(aps *Apns) {
	this.Aps = aps
}

func (this *PushInfo) SetMultimedia(multimedia *Multimedia) {
	this.Multimedia = append(this.Multimedia, multimedia)
}
