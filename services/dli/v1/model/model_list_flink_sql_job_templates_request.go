package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListFlinkSqlJobTemplatesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int64 `json:"offset,omitempty"`

	Order *string `json:"order,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListFlinkSqlJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkSqlJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkSqlJobTemplatesRequest", string(data)}, " ")
}
