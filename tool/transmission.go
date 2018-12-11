package tool

//传递消息内容给客户端，客户端自行定义通知栏展现形式
type Transmission struct {
	TransmissionType    bool                   `json:"transmission_type"`        //	收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否  可选
	TransmissionContent string                 `json:"transmission_content"`     //透传内容
	DurationBegin       string                 `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd         string                 `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Notify *Notify `json:"notify"`
	data map[string]interface{} `json:"data"`
}

func GetTransmission(content string) *Transmission {
	transmission := &Transmission{
		TransmissionContent: content,
		TransmissionType:    false,
		data: make(map[string]interface{}),
	}
	transmission.data["transmission_content"] = content
	transmission.data["transmission_type"] = false
	return transmission
}

func (this *Transmission) SetTransmissionContent(str string) {
	this.TransmissionContent = str
	this.data["transmission_content"] = str
}

func (this *Transmission) SetTransmissionType(is bool) {
	this.TransmissionType = is
	this.data["transmission_type"] = is
}

func (this *Transmission) SetDurationBegin(str string) {
	this.DurationBegin = str
	this.data["duration_begin"] = str
}

func (this *Transmission) SetDurationEnd(str string) {
	this.DurationEnd = str
	this.data["duration_end"] = str
}

func (this *Transmission) SetNotify(notify *Notify) {
	this.Notify = notify
	this.data["notify"] = notify
}


func (this *Transmission) GetData() map[string]interface{} {

	return this.data
}