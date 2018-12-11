package push

import (
	tool "github.com/jukylin/go-getui-api/tool"
	util "github.com/jukylin/go-getui-api/util"
	"encoding/json"
)


//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type saveListBodyParmar struct {
	Message      *tool.Message      `json:"message"` //消息内容
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	PushInfo     map[string]interface{}             `json:"push_info,omitempty"` //json串，当手机为ios，并且为离线的时候；或者简化推送的时候，使用该参数
	TaskName     string             `json:"task_name,omitempty"` //	任务名称 可以给多个任务指定相同的task_name，后面用task_name查询推送结果能得到多个任务的结果  可选
	data map[string]interface{} 	`json:"data"`
}

func SaveListBody(appId string, auth_token string, parmar *saveListBodyParmar) (*util.PushResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/save_list_body"
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

func NewSaveListBodyParmar() *saveListBodyParmar {
	return &saveListBodyParmar{
		data : make(map[string]interface{}),
	}
}

func (this *saveListBodyParmar) SetMessage(message *tool.Message) {
	this.Message = message
	this.data["message"] = message
}

func (this *saveListBodyParmar) SetNotification(notification *tool.Notification) {
	this.Notification = notification
	this.data["notification"] = notification
}


func (this *saveListBodyParmar) SetLink(link *tool.Link) {
	this.Link = link
	this.data["link"] = link
}

func (this *saveListBodyParmar) SetNotypopload(notyPopload *tool.NotyPopload) {
	this.Notypopload = notyPopload
	this.data["notypopload"] = notyPopload
}


func (this *saveListBodyParmar) SetTransmission(trans *tool.Transmission) {
	this.Transmission = trans
	this.data["transmission"] = trans
}

func (this *saveListBodyParmar) SetPushInfo(pushInfo map[string]interface{}) {
	this.PushInfo = pushInfo
	this.data["push_info"] = pushInfo
}

func (this *saveListBodyParmar) SetTaskName(TaskName string) {
	this.TaskName = TaskName
	this.data["task_name"] = TaskName
}



func (this *saveListBodyParmar) GetData() map[string]interface{}{
	return this.data
}