package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateQueueRequest struct {
	Body *CreateQueueRequestBody `json:"body,omitempty"`
}

func (o CreateQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequest struct{}"
	}

	return strings.Join([]string{"CreateQueueRequest", string(data)}, " ")
}
