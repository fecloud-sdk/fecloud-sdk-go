package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePostPaidServersRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	Server *PostPaidServer `json:"server"`
}

func (o CreatePostPaidServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidServersRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePostPaidServersRequestBody", string(data)}, " ")
}
