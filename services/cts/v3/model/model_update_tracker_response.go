package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTrackerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrackerResponse struct{}"
	}

	return strings.Join([]string{"UpdateTrackerResponse", string(data)}, " ")
}
