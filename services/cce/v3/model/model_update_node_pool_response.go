package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateNodePoolResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodePoolMetadata `json:"metadata,omitempty"`

	Spec *NodePoolSpec `json:"spec,omitempty"`

	Status         *NodePoolStatus `json:"status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodePoolResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodePoolResponse", string(data)}, " ")
}
