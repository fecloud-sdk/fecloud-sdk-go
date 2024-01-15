package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenersByTagsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Resources      *[]ResourcesByTag `json:"resources,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListListenersByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListListenersByTagsResponse", string(data)}, " ")
}
