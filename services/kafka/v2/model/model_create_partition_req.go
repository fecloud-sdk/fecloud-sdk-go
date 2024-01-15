package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePartitionReq struct {
	Partition *int32 `json:"partition,omitempty"`
}

func (o CreatePartitionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionReq struct{}"
	}

	return strings.Join([]string{"CreatePartitionReq", string(data)}, " ")
}
