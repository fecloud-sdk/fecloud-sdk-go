package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRocketInstanceTopicsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Max *int32 `json:"max,omitempty"`

	Remaining *int32 `json:"remaining,omitempty"`

	NextOffset *int32 `json:"next_offset,omitempty"`

	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	Topics         *[]Topic `json:"topics,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListRocketInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListRocketInstanceTopicsResponse", string(data)}, " ")
}
