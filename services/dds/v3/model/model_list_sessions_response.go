package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSessionsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Sessions       *[]QuerySessionResponse `json:"sessions,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionsResponse", string(data)}, " ")
}
