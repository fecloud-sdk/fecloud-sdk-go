package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListResizeFlavorsResponse Response Object
type ListResizeFlavorsResponse struct {

	// 云服务器规格列表。
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
