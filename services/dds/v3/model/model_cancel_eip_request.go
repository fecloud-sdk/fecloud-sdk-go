package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CancelEipRequest Request Object
type CancelEipRequest struct {

	// 节点ID。
	NodeId string `json:"node_id"`
}

func (o CancelEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelEipRequest struct{}"
	}

	return strings.Join([]string{"CancelEipRequest", string(data)}, " ")
}
