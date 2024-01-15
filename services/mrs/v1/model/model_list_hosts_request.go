package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostsRequest struct {
	ClusterId string `json:"cluster_id"`

	PageSize *string `json:"pageSize,omitempty"`

	CurrentPage *string `json:"currentPage,omitempty"`
}

func (o ListHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRequest struct{}"
	}

	return strings.Join([]string{"ListHostsRequest", string(data)}, " ")
}
