package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSharedTagsResponse struct {
	Tags *[]ResourceTag `json:"tags,omitempty"`

	SysTags        *[]ResourceTag `json:"sys_tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSharedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSharedTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSharedTagsResponse", string(data)}, " ")
}
