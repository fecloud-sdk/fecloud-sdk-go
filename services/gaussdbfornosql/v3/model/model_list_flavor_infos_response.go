package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlavorInfosResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Flavors        *[]ListFlavorsResult `json:"flavors,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListFlavorInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfosResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorInfosResponse", string(data)}, " ")
}
