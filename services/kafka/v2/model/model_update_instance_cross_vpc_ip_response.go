package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateInstanceCrossVpcIpResponse struct {
	Success *bool `json:"success,omitempty"`

	Results        *[]UpdateInstanceCrossVpcIpRespResults `json:"results,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o UpdateInstanceCrossVpcIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceCrossVpcIpResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceCrossVpcIpResponse", string(data)}, " ")
}
