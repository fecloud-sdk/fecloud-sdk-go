package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePostPaidServersRequest Request Object
type CreatePostPaidServersRequest struct {
	Body *CreatePostPaidServersRequestBody `json:"body,omitempty"`
}

func (o CreatePostPaidServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidServersRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidServersRequest", string(data)}, " ")
}
