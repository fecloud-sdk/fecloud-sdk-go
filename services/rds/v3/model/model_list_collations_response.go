package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListCollationsResponse struct {
	CharSets       *[]string `json:"charSets,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCollationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollationsResponse struct{}"
	}

	return strings.Join([]string{"ListCollationsResponse", string(data)}, " ")
}
