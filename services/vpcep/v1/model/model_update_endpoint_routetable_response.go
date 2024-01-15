package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateEndpointRoutetableResponse struct {
	Routetables *[]string `json:"routetables,omitempty"`

	Error          *[]RoutetableInfoError `json:"error,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateEndpointRoutetableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRoutetableResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRoutetableResponse", string(data)}, " ")
}
