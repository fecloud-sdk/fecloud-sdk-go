package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddReadonlyNodeRequest struct {
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
