package tool

type Alert struct {
	Title string `json:"title"` //通知标题
	Body  string `json:"body"`  //通知内容
	ActionLocKey string `json:"action-loc-key"` //（用于多语言支持）指定执行按钮所使用的Localizable.strings
	LocKey string `json:"loc-key"` //（用于多语言支持）指定Localizable.strings文件中相应的key
	LocArgs []string `json:"loc-args"` //如果loc-key中使用了占位符，则在loc-args中指定各参数
	LaunchImage string `json:"launch-image"` //指定启动界面图片名
	TitileLocKey string `json:"titile-loc-key"` //(用于多语言支持）对于标题指定执行按钮所使用的Localizable.strings,仅支持iOS8.2以上版本
	TitleLocArgs []string `json:"title-loc-args"` //对于标题,如果loc-key中使用的占位符，则在loc-args中指定各参数,仅支持iOS8.2以上版本
	Subtitle string `json:"subtitle"`  //通知子标题,仅支持iOS8.2以上版本
	SubtitleLocKey string `json:"subtitle-loc-key"` //当前本地化文件中的子标题字符串的关键字,仅支持iOS8.2以上版本
	SubtitleLocArgs []string `json:"subtitle-loc-args"` //当前本地化子标题内容中需要置换的变量参数 ,仅支持iOS8.2以上版本
}

func GetAlert(title string, body string) *Alert {
	alert := &Alert{
		Title: title,
		Body:  body,
	}
	return alert
}

func (this *Alert) SetTitle(str string) {
	this.Title = str
}

func (this *Alert) SetBody(str string) {
	this.Body = str
}

func (this *Alert) SetSubtitle(str string) {
	this.Subtitle = str
}
