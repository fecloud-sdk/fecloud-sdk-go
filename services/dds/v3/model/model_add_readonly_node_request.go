package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AddReadonlyNodeRequest Request Object
type AddReadonlyNodeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *AddReadonlyNodeRequestBody `json:"body,omitempty"`
}

func (o AddReadonlyNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddReadonlyNodeRequest struct{}"
	}

	return strings.Join([]string{"AddReadonlyNodeRequest", string(data)}, " ")
}
