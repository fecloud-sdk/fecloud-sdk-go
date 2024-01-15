package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateUserResponse struct {
	AllCounts *int32 `json:"all_counts,omitempty"`

	Results        *[]QueryUserResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchUpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUserResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateUserResponse", string(data)}, " ")
}
