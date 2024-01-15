package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListResizeFlavorsResponse struct {
	Flavors        *[]ListResizeFlavorsResult `json:"flavors,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListResizeFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResizeFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListResizeFlavorsResponse", string(data)}, " ")
}
