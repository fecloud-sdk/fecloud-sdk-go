package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type InstanceRequestSpec struct {
	Version *string `json:"version,omitempty"`

	ClusterID string `json:"clusterID"`

	Values map[string]interface{} `json:"values"`

	AddonTemplateName string `json:"addonTemplateName"`
}

func (o InstanceRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRequestSpec struct{}"
	}

	return strings.Join([]string{"InstanceRequestSpec", string(data)}, " ")
}
