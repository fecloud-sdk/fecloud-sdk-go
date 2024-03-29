package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceRequest struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *AddonMetadata `json:"metadata"`

	Spec *InstanceRequestSpec `json:"spec"`
}

func (o InstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRequest struct{}"
	}

	return strings.Join([]string{"InstanceRequest", string(data)}, " ")
}
