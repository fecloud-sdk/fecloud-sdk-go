package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListProtectionServerResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]ProtectionServerInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListProtectionServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionServerResponse struct{}"
	}

	return strings.Join([]string{"ListProtectionServerResponse", string(data)}, " ")
}
