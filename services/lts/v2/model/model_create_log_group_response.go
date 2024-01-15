package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLogGroupResponse struct {
	LogGroupId     *string `json:"log_group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateLogGroupResponse", string(data)}, " ")
}
