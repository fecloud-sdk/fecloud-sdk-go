package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type ListGrantsRequest struct {
	Body *ListGrantsRequestBody `json:"body,omitempty"`
}

func (o ListGrantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGrantsRequest struct{}"
	}

	return strings.Join([]string{"ListGrantsRequest", string(data)}, " ")
}
