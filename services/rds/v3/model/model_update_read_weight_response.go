package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateReadWeightResponse struct {
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReadWeightResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReadWeightResponse struct{}"
	}

	return strings.Join([]string{"UpdateReadWeightResponse", string(data)}, " ")
}
