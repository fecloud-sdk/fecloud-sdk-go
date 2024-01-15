package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type WeakPasswordCheckRequestBody struct {
	Password string `json:"password"`
}

func (o WeakPasswordCheckRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPasswordCheckRequestBody struct{}"
	}

	return strings.Join([]string{"WeakPasswordCheckRequestBody", string(data)}, " ")
}
