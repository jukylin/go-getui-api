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
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	Condition    []*tool.Condition  `json:"condition,omitempty"` //	筛选目标用户条件，参考下面的condition说明  可选
	Requestid    string             `json:"requestid"`           //请求唯一标识
	Speed	int        `json:"speed"`
}

func ToApp(appId string, auth_token string, toAppParmar *ToAppParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_app"
	bodyByte, err := util.GetBody(toAppParmar)
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
