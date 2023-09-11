package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ClusterNodeInformationMetadata struct {
	Name string `json:"name"`
}

func (o ClusterNodeInformationMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeInformationMetadata struct{}"
	}

	return strings.Join([]string{"ClusterNodeInformationMetadata", string(data)}, " ")
}
