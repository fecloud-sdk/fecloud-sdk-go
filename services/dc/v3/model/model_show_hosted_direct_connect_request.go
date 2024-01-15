package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowHostedDirectConnectRequest struct {
	HostedConnectId string `json:"hosted_connect_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Fields *[]string `json:"fields,omitempty"`

	SortDir *[]ShowHostedDirectConnectRequestSortDir `json:"sort_dir,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	HostingId *[]string `json:"hosting_id,omitempty"`
}

func (o ShowHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"ShowHostedDirectConnectRequest", string(data)}, " ")
}

type ShowHostedDirectConnectRequestSortDir struct {
	value string
}

type ShowHostedDirectConnectRequestSortDirEnum struct {
	ASC  ShowHostedDirectConnectRequestSortDir
	DESC ShowHostedDirectConnectRequestSortDir
}

func GetShowHostedDirectConnectRequestSortDirEnum() ShowHostedDirectConnectRequestSortDirEnum {
	return ShowHostedDirectConnectRequestSortDirEnum{
		ASC: ShowHostedDirectConnectRequestSortDir{
			value: "asc",
		},
		DESC: ShowHostedDirectConnectRequestSortDir{
			value: "desc",
		},
	}
}

func (c ShowHostedDirectConnectRequestSortDir) Value() string {
	return c.value
}

func (c ShowHostedDirectConnectRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHostedDirectConnectRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
