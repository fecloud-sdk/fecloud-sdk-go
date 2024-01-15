package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTopicResp struct {
	Id *string `json:"id,omitempty"`
}

func (o CreateTopicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicResp struct{}"
	}

	return strings.Join([]string{"CreateTopicResp", string(data)}, " ")
}
