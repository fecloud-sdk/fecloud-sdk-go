package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateAgencyMappingResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAgencyMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyMappingResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyMappingResponse", string(data)}, " ")
}
