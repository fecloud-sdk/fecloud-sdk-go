package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResizeServerRequestBody struct {
	Resize *ResizePrePaidServerOption `json:"resize"`

	DryRun *bool `json:"dry_run,omitempty"`
}

func (o ResizeServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeServerRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeServerRequestBody", string(data)}, " ")
}
