package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelEipRequest struct {
	NodeId string `json:"node_id"`
}

func (o CancelEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelEipRequest struct{}"
	}

	return strings.Join([]string{"CancelEipRequest", string(data)}, " ")
}
