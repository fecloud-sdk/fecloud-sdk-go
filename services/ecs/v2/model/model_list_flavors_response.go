package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListFlavorsResponse Response Object
type ListFlavorsResponse struct {

	// 云服务器规格列表。
	Flavors        *[]Flavor `json:"flavors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
