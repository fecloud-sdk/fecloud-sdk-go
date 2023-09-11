package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListApiVersionRequest Request Object
type ListApiVersionRequest struct {
}

func (o ListApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionRequest", string(data)}, " ")
}
