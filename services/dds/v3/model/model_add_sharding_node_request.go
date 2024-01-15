package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddShardingNodeRequest struct {
	InstanceId string `json:"instance_id"`

	Body *EnlargeInstanceRequestBody `json:"body,omitempty"`
}

func (o AddShardingNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeRequest struct{}"
	}

	return strings.Join([]string{"AddShardingNodeRequest", string(data)}, " ")
}
