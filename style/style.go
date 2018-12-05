package style

//open application templates
type Style struct {
	Type        int    `json:"type"`                    //固定为6
	Text        string `json:"text"`                    //通知正文
	Title       string `json:"title"`                   //通知标题
	Logo        string `json:"logo,omitempty"`          //通知的图标名称，包含后缀名（需要在客户端开发时嵌入），如“push.png” 可选
	LogoUrl     string `json:"logourl,omitempty"`       //通知的图标logourl 可选
	BigStyle    string `json:"big_style"`               //通知展示样式,枚举值包括 1,2,3
	BigImageUrl string `json:"big_image_url,omitempty"` //通知大图URL地址  对应枚举值 1
	BigText     string `json:"big_text,omitempty"`      //通知大图URL地址  对应枚举值 2
	BannerUrl   string `json:"banner_url,omitempty"`    //通知小图URL地址
	IsRing      bool   `json:"is_ring,omitempty"`       //收到通知是否响铃：true响铃，false不响铃。默认响铃  可选
	IsVibrate   bool   `json:"is_vibrate,omitempty"`    //收到通知是否振动：true振动，false不振动。默认振动  可选
	IsClearable bool   `json:"is_clearable,omitempty"`  //通知是否可清除： true可清除，false不可清除。默认可清除  可选
}

func GetStyle(text string, title string, _type int) *Style {
	style := &Style{
		Type:        _type,
		Text:        text,
		Title:       title,
		IsRing:      true,
		IsVibrate:   true,
		IsClearable: true,
	}
	return style
}

func (this *Style) SetType(t int) {
	this.Type = t
}

func (this *Style) SetText(str string) {
	this.Text = str
}

func (this *Style) SetTitle(str string) {
	this.Title = str
}

func (this *Style) SetLogo(s string) {
	this.Logo = s
}

func (this *Style) SetLogoUrl(s string) {
	this.LogoUrl = s
}

func (this *Style) SetBigStyle(s string) {
	this.BigStyle = s
}

func (this *Style) SetBigImageUrl(s string) {
	this.BigImageUrl = s
}

func (this *Style) SetBigText(s string) {
	this.BigText = s
}

func (this *Style) SetBannerUrl(s string) {
	this.BannerUrl = s
}

func (this *Style) SetIsRing(is bool) {
	this.IsRing = is
}

func (this *Style) SetIsVibrate(is bool) {
	this.IsVibrate = is
}

func (this *Style) SetIsClearable(is bool) {
	this.IsClearable = is
}
