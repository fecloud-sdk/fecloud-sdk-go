package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProductsResponse struct {
	Hourly *[]ListProductsRespHourly `json:"Hourly,omitempty"`

	Monthly        *[]ListProductsRespHourly `json:"Monthly,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsResponse struct{}"
	}

	return strings.Join([]string{"ListProductsResponse", string(data)}, " ")
}
