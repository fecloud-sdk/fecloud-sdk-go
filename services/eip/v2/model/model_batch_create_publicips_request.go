package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchCreatePublicipsRequest Request Object
type BatchCreatePublicipsRequest struct {
	Body *BatchCreatePublicipsV2RequestBody `json:"body,omitempty"`
}

func (o BatchCreatePublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipsRequest", string(data)}, " ")
}
