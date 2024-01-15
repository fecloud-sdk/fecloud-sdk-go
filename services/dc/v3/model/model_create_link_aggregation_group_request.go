package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLinkAggregationGroupRequest struct {
	Body *CreateLinkAggregationGroupRequestBody `json:"body,omitempty"`
}

func (o CreateLinkAggregationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLinkAggregationGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateLinkAggregationGroupRequest", string(data)}, " ")
}
