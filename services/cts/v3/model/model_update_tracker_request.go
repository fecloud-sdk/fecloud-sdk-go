package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateTrackerRequest struct {
	Body *UpdateTrackerRequestBody `json:"body,omitempty"`
}

func (o UpdateTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrackerRequest struct{}"
	}

	return strings.Join([]string{"UpdateTrackerRequest", string(data)}, " ")
}
