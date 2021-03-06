package push

import (
	tool "github.com/jukylin/go-getui-api/tool"
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSingleParmar struct {
	Message      *tool.Message      `json:"message"`
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission map[string]interface{} `json:"transmission,omitempty"`
	Cid          string             `json:"cid,omitempty"`
	Alias        string             `json:"alias,omitempty"`
	RequestId    string             `json:"requestid"`
	PushInfo     map[string]interface{} `json:"push_info"`
	data map[string]interface{} `json:"data"`
}

func PushSingle(appId string, auth_token string, parmar *PushSingleParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_single"
	bodyByte, err := util.GetBody(parmar.GetData())
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

func NewPushSingleParmar() *PushSingleParmar {
	return &PushSingleParmar{
		data : make(map[string]interface{}),
	}
}

func (this *PushSingleParmar) SetMessage(message *tool.Message) {
	this.Message = message
	this.data["message"] = message
}

func (this *PushSingleParmar) SetNotification(notification *tool.Notification) {
	this.Notification = notification
	this.data["notification"] = notification
}


func (this *PushSingleParmar) SetLink(link *tool.Link) {
	this.Link = link
	this.data["link"] = link
}

func (this *PushSingleParmar) SetNotypopload(notyPopload *tool.NotyPopload) {
	this.Notypopload = notyPopload
	this.data["notypopload"] = notyPopload
}


func (this *PushSingleParmar) SetTransmission(trans map[string]interface{}) {
	this.Transmission = trans
	this.data["transmission"] = trans
}

func (this *PushSingleParmar) SetCid(cid string) {
	this.Cid = cid
	this.data["cid"] = cid
}

func (this *PushSingleParmar) SetAlias(alias string) {
	this.Alias = alias
	this.data["alias"] = alias
}

func (this *PushSingleParmar) SetRequestId(reqId string) {
	this.RequestId = reqId
	this.data["requestid"] = reqId
}

func (this *PushSingleParmar) SetPushInfo(pushInfo map[string]interface{}) {
	this.PushInfo = pushInfo
	this.data["push_info"] = pushInfo
}


func (this *PushSingleParmar) GetData() map[string]interface{}{
	return this.data
}