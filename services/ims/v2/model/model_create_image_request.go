package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateImageRequest struct {
	Body *CreateImageRequestBody `json:"body,omitempty"`
}

func (o CreateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageRequest struct{}"
	}

	return strings.Join([]string{"CreateImageRequest", string(data)}, " ")
}
