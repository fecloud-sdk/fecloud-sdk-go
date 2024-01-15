package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListRetirableGrantsRequest struct {
	Body *ListRetirableGrantsRequestBody `json:"body,omitempty"`
}

func (o ListRetirableGrantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetirableGrantsRequest struct{}"
	}

	return strings.Join([]string{"ListRetirableGrantsRequest", string(data)}, " ")
}
