package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchShowNodesInformationResponse struct {
	Count *int32 `json:"count,omitempty"`

	Instances      *[]InstanceNodesInfoResp `json:"instances,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchShowNodesInformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowNodesInformationResponse struct{}"
	}

	return strings.Join([]string{"BatchShowNodesInformationResponse", string(data)}, " ")
}
