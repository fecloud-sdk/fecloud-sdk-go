package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteDatabaseUserResponse Response Object
type DeleteDatabaseUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserResponse", string(data)}, " ")
}
