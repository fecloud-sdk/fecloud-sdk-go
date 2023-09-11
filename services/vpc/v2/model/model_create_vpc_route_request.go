package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcRouteRequest struct {
	Body *CreateVpcRouteRequestBody `json:"body,omitempty"`
}

func (o CreateVpcRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteRequest", string(data)}, " ")
}
