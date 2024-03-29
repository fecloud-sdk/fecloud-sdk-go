package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListTrackersResponse struct {
	Trackers       *[]TrackerResponseBody `json:"trackers,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListTrackersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrackersResponse struct{}"
	}

	return strings.Join([]string{"ListTrackersResponse", string(data)}, " ")
}
