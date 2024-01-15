package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListWeakPasswordUsersResponse struct {
	TotalNum *int64 `json:"total_num,omitempty"`

	DataList       *[]WeakPwdListInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListWeakPasswordUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWeakPasswordUsersResponse struct{}"
	}

	return strings.Join([]string{"ListWeakPasswordUsersResponse", string(data)}, " ")
}
