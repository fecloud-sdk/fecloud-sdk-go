package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowJobDetailRequest Request Object
type ShowJobDetailRequest struct {

	// 任务ID。
	Id string `json:"id"`
}

func (o ShowJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}
