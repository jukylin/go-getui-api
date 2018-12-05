package push

import (
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)

type PushListParmar struct {
	TaskId     string   `json:"taskid"`                //	任务号，取save_list_body返回的taskid
	Cid        []string `json:"cid,omitempty"`         //cid为cid list，与alias list二选一
	Alias      []string `json:"alias,omitempty"`       //	alias为alias list，与cid list二选一
	NeedDetail bool     `json:"need_detail,omitempty"` //默认值:false，是否需要返回每个CID的状态
}

func PushList(appId string, auth_token string, parmar *PushListParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_list"
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
