package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type TotalMetaData struct {
	Total *int32 `json:"total,omitempty"`
}

func (o TotalMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalMetaData struct{}"
	}

	return strings.Join([]string{"TotalMetaData", string(data)}, " ")
}
