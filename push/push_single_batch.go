package push

import (
	tool "github.com/jukylin/go-getui-api/tool"
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)

type PushSigleBatchListParmar struct {
	MsgList    []*PushSigleBatchParmar `json:"msg_list"`              //
	NeedDetail bool                    `json:"need_detail,omitempty"` //默认值:false，是否需要返回每个CID的状态
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSigleBatchParmar struct {
	Message      *tool.Message      `json:"message"` //消息内容
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	Cid          string             `json:"cid"`   //cid为cid list，与alias list二选一
	Alias        string             `json:"alias"` //	alias为alias list，与cid list二选一
	RequestId    string             `json:"requestid"`
	PushInfo     map[string]interface{} `json:"push_info"`
}

func PushSigleBatch(appId string, auth_token string, parmar *PushSigleBatchListParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_single_batch"
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
