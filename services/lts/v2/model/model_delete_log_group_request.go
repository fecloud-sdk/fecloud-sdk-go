package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLogGroupRequest struct {
	LogGroupId string `json:"log_group_id"`

	ContentType string `json:"Content-Type"`
}

func (o DeleteLogGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogGroupRequest", string(data)}, " ")
}
