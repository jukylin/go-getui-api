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
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	Cid          string             `json:"cid,omitempty"`
	Alias        string             `json:"alias,omitempty"`
	RequestId    string             `json:"requestid"`
	PushInfo     *tool.PushInfo     `json:"push_info"`
}


func PushSingle(appId string, auth_token string, parmar *PushSingleParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_single"
	bodyByte, err := util.GetBody(parmar)
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
