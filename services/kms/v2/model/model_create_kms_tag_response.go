package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Response Object
type CreateKmsTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateKmsTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagResponse struct{}"
	}

	return strings.Join([]string{"CreateKmsTagResponse", string(data)}, " ")
}
