package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchTagActionDelRequestBody struct {
	Action string `json:"action"`

	Tags []TagDelWithKeyValue `json:"tags"`
}

func (o BatchTagActionDelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionDelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionDelRequestBody", string(data)}, " ")
}
