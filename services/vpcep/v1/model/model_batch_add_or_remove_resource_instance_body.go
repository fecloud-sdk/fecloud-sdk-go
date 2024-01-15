package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchAddOrRemoveResourceInstanceBody struct {
	Tags *[]ResourceTag `json:"tags,omitempty"`

	Action string `json:"action"`
}

func (o BatchAddOrRemoveResourceInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceBody struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceBody", string(data)}, " ")
}
