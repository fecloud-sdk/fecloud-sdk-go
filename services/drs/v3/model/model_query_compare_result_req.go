package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryCompareResultReq struct {
	JobId string `json:"job_id"`

	ObjectLevelCompareId *string `json:"object_level_compare_id,omitempty"`

	LineCompareId *string `json:"line_compare_id,omitempty"`

	ContentCompareId *string `json:"content_compare_id,omitempty"`

	CurrentPage int32 `json:"current_page"`

	PerPage int32 `json:"per_page"`
}

func (o QueryCompareResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCompareResultReq struct{}"
	}

	return strings.Join([]string{"QueryCompareResultReq", string(data)}, " ")
}
