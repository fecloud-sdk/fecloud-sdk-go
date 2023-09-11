package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVpcRoutesResponse struct {
	Routes *[]VpcRoute `json:"routes,omitempty"`

	RoutesLinks    *[]NeutronPageLink `json:"routes_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVpcRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListVpcRoutesResponse", string(data)}, " ")
}
