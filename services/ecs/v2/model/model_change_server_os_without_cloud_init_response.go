package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ChangeServerOsWithoutCloudInitResponse Response Object
type ChangeServerOsWithoutCloudInitResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeServerOsWithoutCloudInitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithoutCloudInitResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithoutCloudInitResponse", string(data)}, " ")
}