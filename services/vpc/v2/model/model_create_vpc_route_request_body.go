package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcRouteRequestBody struct {
	Route *CreateVpcRouteOption `json:"route"`
}

func (o CreateVpcRouteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteRequestBody", string(data)}, " ")
}
