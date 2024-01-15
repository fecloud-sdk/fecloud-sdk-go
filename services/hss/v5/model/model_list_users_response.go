package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListUsersResponse struct {
	TotalNum *int32 `json:"total_num,omitempty"`

	DataList       *[]UserResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
