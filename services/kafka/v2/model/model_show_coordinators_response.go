package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowCoordinatorsResponse struct {
	Coordinators   *[]ShowCoordinatorsRespCoordinators `json:"coordinators,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowCoordinatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCoordinatorsResponse struct{}"
	}

	return strings.Join([]string{"ShowCoordinatorsResponse", string(data)}, " ")
}
