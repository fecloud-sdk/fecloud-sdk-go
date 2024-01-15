package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ListServicePublicDetailsRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	Id *string `json:"id,omitempty"`

	SortKey *ListServicePublicDetailsRequestSortKey `json:"sort_key,omitempty"`

	SortDir *ListServicePublicDetailsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListServicePublicDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePublicDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServicePublicDetailsRequest", string(data)}, " ")
}

type ListServicePublicDetailsRequestSortKey struct {
	value string
}

type ListServicePublicDetailsRequestSortKeyEnum struct {
	CREATE_AT ListServicePublicDetailsRequestSortKey
	UPDATE_AT ListServicePublicDetailsRequestSortKey
}

func GetListServicePublicDetailsRequestSortKeyEnum() ListServicePublicDetailsRequestSortKeyEnum {
	return ListServicePublicDetailsRequestSortKeyEnum{
		CREATE_AT: ListServicePublicDetailsRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListServicePublicDetailsRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListServicePublicDetailsRequestSortKey) Value() string {
	return c.value
}

func (c ListServicePublicDetailsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePublicDetailsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListServicePublicDetailsRequestSortDir struct {
	value string
}

type ListServicePublicDetailsRequestSortDirEnum struct {
	ASC  ListServicePublicDetailsRequestSortDir
	DESC ListServicePublicDetailsRequestSortDir
}

func GetListServicePublicDetailsRequestSortDirEnum() ListServicePublicDetailsRequestSortDirEnum {
	return ListServicePublicDetailsRequestSortDirEnum{
		ASC: ListServicePublicDetailsRequestSortDir{
			value: "asc",
		},
		DESC: ListServicePublicDetailsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListServicePublicDetailsRequestSortDir) Value() string {
	return c.value
}

func (c ListServicePublicDetailsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePublicDetailsRequestSortDir) UnmarshalJSON(b []byte) error {
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
