package push

import (
	tool "github.com/jukylin/go-getui-api/tool"
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type ToAppParmar struct {
	Message      *tool.Message      `json:"message"` //消息内容
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission map[string]interface{} `json:"transmission,omitempty"`
	Condition    []*tool.Condition  `json:"condition,omitempty"` //	筛选目标用户条件，参考下面的condition说明  可选
	RequestId    string             `json:"requestid"`           //请求唯一标识
	Speed	int        `json:"speed"`
	PushInfo     map[string]interface{} `json:"push_info"`
	data map[string]interface{} 	`json:"data"`
}

func ToApp(appId string, auth_token string, ToAppParmar *ToAppParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_app"
	bodyByte, err := util.GetBody(ToAppParmar.GetData())
	if err != nil {
		return nil, err
	}

	result, err := util.BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	pushResult := new(util.PushResult)
	if err := json.Unmarshal([]byte(result), &pushResult); err != nil {
		return nil, err
	}

	return pushResult, err
}


func NewToAppParmar() *ToAppParmar {
	return &ToAppParmar{
		data : make(map[string]interface{}),
	}
}

func (this *ToAppParmar) SetMessage(message *tool.Message) {
	this.Message = message
	this.data["message"] = message
}

func (this *ToAppParmar) SetNotification(notification *tool.Notification) {
	this.Notification = notification
	this.data["notification"] = notification
}


func (this *ToAppParmar) SetLink(link *tool.Link) {
	this.Link = link
	this.data["link"] = link
}

func (this *ToAppParmar) SetNotypopload(notyPopload *tool.NotyPopload) {
	this.Notypopload = notyPopload
	this.data["notypopload"] = notyPopload
}


func (this *ToAppParmar) SetTransmission(trans map[string]interface{}) {
	this.Transmission = trans
	this.data["transmission"] = trans
}

func (this *ToAppParmar) SetRequestId(reqId string) {
	this.RequestId = reqId
	this.data["requestid"] = reqId
}

func (this *ToAppParmar) SetPushInfo(pushInfo map[string]interface{}) {
	this.PushInfo = pushInfo
	this.data["push_info"] = pushInfo
}


func (this *ToAppParmar) SetSpeed(speed int) {
	this.Speed = speed
	this.data["speed"] = speed
}


func (this *ToAppParmar) SetCondition(condition []*tool.Condition) {
	this.Condition = condition
	this.data["condition"] = condition
}


func (this *ToAppParmar) GetData() map[string]interface{}{
	return this.data
}