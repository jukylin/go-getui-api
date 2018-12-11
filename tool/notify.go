package tool

//传递消息内容给客户端，客户端自行定义通知栏展现形式
type Notify struct {
	Title    string    `json:"title"`        //通知栏标题,如果使用通知功能则必填
	Content string     `json:"content"`     //通知栏内容，如果使用通知功能则必填
	Intent  string     `json:"intent,omitempty"` //	长度小于 1000 字节，通知带 intent 传递参数，推荐使用， intent 最好由 Android 开发工程师生成，生成方式见附录
	Url  string                 `json:"url,omitempty"`   //点击通知，打开链接 可选
	Payload string `json:"payload"` //透传内容，类似 APNPayload，使用此方式的透传消息可 能会被华为拦截，不推荐使用
	Type string `json:"type"` // 取值为 0 代表 payload（默认），1 代表 intent，2 代表 url。,如果设置了 payload、url、intent，需要指定 type
}

func GetNotify(title, content string) *Notify {
	notify := &Notify{
		Title: title,
		Content: content,
	}
	return notify
}

func (this *Notify) SetTitle(title string) {
	this.Title = title
}

func (this *Notify) SetContent(content string) {
	this.Content = content
}

func (this *Notify) SetIntent(str string) {
	this.Intent = str
}

func (this *Notify) SetUrl(str string) {
	this.Url = str
}

func (this *Notify) SetPayload(str string) {
	this.Payload = str
}

func (this *Notify) SetType(t string) {
	this.Type = t
}