package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchListJobStatusResponse Response Object
type BatchListJobStatusResponse struct {

	// 任务状态信息
	Results *[]QueryJobStatusResp `json:"results,omitempty"`

	// 返回任务数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobStatusResponse", string(data)}, " ")
}
