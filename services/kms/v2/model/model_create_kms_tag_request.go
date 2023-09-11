package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateKmsTagRequest Request Object
type CreateKmsTagRequest struct {

	// 密钥ID
	KeyId string `json:"key_id"`

	Body *CreateKmsTagRequestBody `json:"body,omitempty"`
}

func (o CreateKmsTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagRequest struct{}"
	}

	return strings.Join([]string{"CreateKmsTagRequest", string(data)}, " ")
}
