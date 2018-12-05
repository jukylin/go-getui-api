package push

import (
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)


func StopTask(appId string, auth_token string, taskid string) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/stop_task/" + taskid

	result, err := util.Delete(url, auth_token, nil)
	if err != nil {
		return nil, err
	}

	pushResult := new(util.PushResult)
	if err := json.Unmarshal([]byte(result), &pushResult); err != nil {
		return nil, err
	}

	return pushResult, err
}
