package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteBigkeyScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBigkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBigkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteBigkeyScanTaskResponse", string(data)}, " ")
}
