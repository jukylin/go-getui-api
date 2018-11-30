package query

import (
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)

type QueryAppPushResult struct {
	Result          string `json:"result"`          //操作结果 无效用户返回no_user
	MsgTotal        int64  `json:"msg_total"`       //百日内活跃用户数
	OnLineNum       string `json:"on_line_num"`     //消息实际下发数
	MsgProcess      int64 `json:"msg_process"`      //消息接收数
	ShowNum         int64 `json:"show_num"`         //消息展示数
	ClickNum       int64  `json:"click_num"`       //消息点击数
	Desc           srting `json:"desc"`           //错误详情
}

func GetPushResultByGroupName(appId string, auth_token string, group_name string) (*QueryAppPushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/get_push_result_by_group_name/" + group_name

	result, err := util.Get(url, auth_token)
	if err != nil {
		return nil, err
	}

	queryAppPushResult := new(QueryAppPushResult)
	if err := json.Unmarshal([]byte(result), &queryAppPushResult); err != nil {
		return nil, err
	}

	return queryAppPushResult, err
}
