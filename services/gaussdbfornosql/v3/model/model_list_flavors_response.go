package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlavorsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Flavors        *[]ListFlavorsResult `json:"flavors,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
