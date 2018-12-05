package push

import (
	tool "github.com/jukylin/go-getui-api/tool"
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)


//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type SaveListBodyParmar struct {
	Message      *tool.Message      `json:"message"` //消息内容
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	PushInfo     string             `json:"push_info,omitempty"` //json串，当手机为ios，并且为离线的时候；或者简化推送的时候，使用该参数
	TaskName     string             `json:"task_name,omitempty"` //	任务名称 可以给多个任务指定相同的task_name，后面用task_name查询推送结果能得到多个任务的结果  可选
}

func SaveListBody(appId string, auth_token string, parmar *SaveListBodyParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/save_list_body"
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
