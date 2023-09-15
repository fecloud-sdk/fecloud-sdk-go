package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteHotkeyScanTaskResponse Response Object
type DeleteHotkeyScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHotkeyScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHotkeyScanTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteHotkeyScanTaskResponse", string(data)}, " ")
}
