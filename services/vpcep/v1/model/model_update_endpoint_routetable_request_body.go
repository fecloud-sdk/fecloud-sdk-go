package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointRoutetableRequestBody struct {
	Routetables []string `json:"routetables"`
}

func (o UpdateEndpointRoutetableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRoutetableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRoutetableRequestBody", string(data)}, " ")
}
