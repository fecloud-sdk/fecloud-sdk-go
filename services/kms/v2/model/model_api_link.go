package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ApiLink struct {
	Href *string `json:"href,omitempty"`

	Rel *string `json:"rel,omitempty"`
}

func (o ApiLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiLink struct{}"
	}

	return strings.Join([]string{"ApiLink", string(data)}, " ")
}
