package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateClusterScalingResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClusterScalingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterScalingResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterScalingResponse", string(data)}, " ")
}
