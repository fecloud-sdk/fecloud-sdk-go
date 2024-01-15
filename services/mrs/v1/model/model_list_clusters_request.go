package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListClustersRequest struct {
	Tags *string `json:"tags,omitempty"`

	PageSize *string `json:"pageSize,omitempty"`

	CurrentPage *string `json:"currentPage,omitempty"`

	ClusterName *string `json:"clusterName,omitempty"`

	ClusterState *string `json:"clusterState,omitempty"`

	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
