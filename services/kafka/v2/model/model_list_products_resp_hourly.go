package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProductsRespHourly struct {
	Name *string `json:"name,omitempty"`

	Version *string `json:"version,omitempty"`

	Values *[]ListProductsRespValues `json:"values,omitempty"`
}

func (o ListProductsRespHourly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespHourly struct{}"
	}

	return strings.Join([]string{"ListProductsRespHourly", string(data)}, " ")
}
