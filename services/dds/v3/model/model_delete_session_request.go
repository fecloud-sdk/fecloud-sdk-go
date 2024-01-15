package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteSessionRequest struct {
	NodeId string `json:"node_id"`

	Body *DeleteSessionRequestBody `json:"body,omitempty"`
}

func (o DeleteSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSessionRequest", string(data)}, " ")
}
