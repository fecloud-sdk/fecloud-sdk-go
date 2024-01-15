package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AttachEipRequest struct {
	NodeId string `json:"node_id"`

	Body *AttachEipRequestBody `json:"body,omitempty"`
}

func (o AttachEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequest struct{}"
	}

	return strings.Join([]string{"AttachEipRequest", string(data)}, " ")
}
