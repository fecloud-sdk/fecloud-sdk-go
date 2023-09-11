package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListDbUsersResponse Response Object
type ListDbUsersResponse struct {

	// 用户信息。
	Users *[]UserForList `json:"users,omitempty"`

	// 总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDbUsersResponse", string(data)}, " ")
}
