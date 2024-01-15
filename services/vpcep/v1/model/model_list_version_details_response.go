package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListVersionDetailsResponse struct {
	Versions       *[]VersionObject `json:"versions,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListVersionDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListVersionDetailsResponse", string(data)}, " ")
}
