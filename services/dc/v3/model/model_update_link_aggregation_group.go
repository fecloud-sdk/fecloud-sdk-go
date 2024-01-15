package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateLinkAggregationGroup struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	MinUpNum *int32 `json:"min_up_num,omitempty"`

	AssociateDc *[]string `json:"associate_dc,omitempty"`

	DisassociateDc *[]string `json:"disassociate_dc,omitempty"`
}

func (o UpdateLinkAggregationGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLinkAggregationGroup struct{}"
	}

	return strings.Join([]string{"UpdateLinkAggregationGroup", string(data)}, " ")
}
