package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListEngineProductsResponse struct {
	Engine *string `json:"engine,omitempty"`

	Versions *[]string `json:"versions,omitempty"`

	Products       *[]ListEngineProductsEntity `json:"products,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListEngineProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsResponse struct{}"
	}

	return strings.Join([]string{"ListEngineProductsResponse", string(data)}, " ")
}
