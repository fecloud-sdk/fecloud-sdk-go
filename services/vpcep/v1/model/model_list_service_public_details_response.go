package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListServicePublicDetailsResponse struct {
	EndpointServices *[]EndpointService `json:"endpoint_services,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServicePublicDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePublicDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServicePublicDetailsResponse", string(data)}, " ")
}
