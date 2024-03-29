package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddImageTagRequest struct {
	ImageId string `json:"image_id"`

	Body *AddImageTagRequestBody `json:"body,omitempty"`
}

func (o AddImageTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagRequest struct{}"
	}

	return strings.Join([]string{"AddImageTagRequest", string(data)}, " ")
}
