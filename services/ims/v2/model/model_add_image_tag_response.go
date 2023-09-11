package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AddImageTagResponse Response Object
type AddImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddImageTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageTagResponse struct{}"
	}

	return strings.Join([]string{"AddImageTagResponse", string(data)}, " ")
}
