package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListClustersResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *[]Cluster `json:"items,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersResponse struct{}"
	}

	return strings.Join([]string{"ListClustersResponse", string(data)}, " ")
}
