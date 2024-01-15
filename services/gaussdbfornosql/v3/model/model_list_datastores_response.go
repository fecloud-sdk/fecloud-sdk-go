package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListDatastoresResponse struct {
	Versions       *[]string `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatastoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastoresResponse struct{}"
	}

	return strings.Join([]string{"ListDatastoresResponse", string(data)}, " ")
}
