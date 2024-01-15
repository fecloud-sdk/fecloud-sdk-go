package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostStatusResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]Host `json:"data_list,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHostStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostStatusResponse struct{}"
	}

	return strings.Join([]string{"ListHostStatusResponse", string(data)}, " ")
}
