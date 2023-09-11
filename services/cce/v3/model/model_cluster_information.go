package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ClusterInformation struct {
	Spec *ClusterInformationSpec `json:"spec"`

	Metadata *ClusterMetadataForUpdate `json:"metadata,omitempty"`
}

func (o ClusterInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInformation struct{}"
	}

	return strings.Join([]string{"ClusterInformation", string(data)}, " ")
}
