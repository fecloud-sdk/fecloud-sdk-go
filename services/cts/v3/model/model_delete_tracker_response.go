package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteTrackerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrackerResponse", string(data)}, " ")
}
