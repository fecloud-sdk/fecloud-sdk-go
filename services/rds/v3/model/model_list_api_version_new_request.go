package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListApiVersionNewRequest struct {
}

func (o ListApiVersionNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionNewRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionNewRequest", string(data)}, " ")
}
