package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListPrivateNatsByTagsResponse struct {
	Resources *[]Resource `json:"resources,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPrivateNatsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsByTagsResponse", string(data)}, " ")
}
