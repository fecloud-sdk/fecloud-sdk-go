package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListServicePermissionsDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	Permission *string `json:"permission,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	SortKey *ListServicePermissionsDetailsRequestSortKey `json:"sort_key,omitempty"`

	SortDir *ListServicePermissionsDetailsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListServicePermissionsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePermissionsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServicePermissionsDetailsRequest", string(data)}, " ")
}

type ListServicePermissionsDetailsRequestSortKey struct {
	value string
}

type ListServicePermissionsDetailsRequestSortKeyEnum struct {
	CREATE_AT ListServicePermissionsDetailsRequestSortKey
	UPDATE_AT ListServicePermissionsDetailsRequestSortKey
}

func GetListServicePermissionsDetailsRequestSortKeyEnum() ListServicePermissionsDetailsRequestSortKeyEnum {
	return ListServicePermissionsDetailsRequestSortKeyEnum{
		CREATE_AT: ListServicePermissionsDetailsRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListServicePermissionsDetailsRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListServicePermissionsDetailsRequestSortKey) Value() string {
	return c.value
}

func (c ListServicePermissionsDetailsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePermissionsDetailsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListServicePermissionsDetailsRequestSortDir struct {
	value string
}

type ListServicePermissionsDetailsRequestSortDirEnum struct {
	ASC  ListServicePermissionsDetailsRequestSortDir
	DESC ListServicePermissionsDetailsRequestSortDir
}

func GetListServicePermissionsDetailsRequestSortDirEnum() ListServicePermissionsDetailsRequestSortDirEnum {
	return ListServicePermissionsDetailsRequestSortDirEnum{
		ASC: ListServicePermissionsDetailsRequestSortDir{
			value: "asc",
		},
		DESC: ListServicePermissionsDetailsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListServicePermissionsDetailsRequestSortDir) Value() string {
	return c.value
}

func (c ListServicePermissionsDetailsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePermissionsDetailsRequestSortDir) UnmarshalJSON(b []byte) error {
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
