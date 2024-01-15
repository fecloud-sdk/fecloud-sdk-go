package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type ShowHdfsFileListRequest struct {
	ClusterId string `json:"cluster_id"`

	Path string `json:"path"`

	Offset *string `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`

	SortKey *ShowHdfsFileListRequestSortKey `json:"sort_key,omitempty"`

	Order *ShowHdfsFileListRequestOrder `json:"order,omitempty"`
}

func (o ShowHdfsFileListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdfsFileListRequest struct{}"
	}

	return strings.Join([]string{"ShowHdfsFileListRequest", string(data)}, " ")
}

type ShowHdfsFileListRequestSortKey struct {
	value string
}

type ShowHdfsFileListRequestSortKeyEnum struct {
	PATH_SUFFIX       ShowHdfsFileListRequestSortKey
	LENGTH            ShowHdfsFileListRequestSortKey
	MODIFICATION_TIME ShowHdfsFileListRequestSortKey
}

func GetShowHdfsFileListRequestSortKeyEnum() ShowHdfsFileListRequestSortKeyEnum {
	return ShowHdfsFileListRequestSortKeyEnum{
		PATH_SUFFIX: ShowHdfsFileListRequestSortKey{
			value: "path_suffix",
		},
		LENGTH: ShowHdfsFileListRequestSortKey{
			value: "length",
		},
		MODIFICATION_TIME: ShowHdfsFileListRequestSortKey{
			value: "modification_time",
		},
	}
}

func (c ShowHdfsFileListRequestSortKey) Value() string {
	return c.value
}

func (c ShowHdfsFileListRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHdfsFileListRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ShowHdfsFileListRequestOrder struct {
	value string
}

type ShowHdfsFileListRequestOrderEnum struct {
	DESC ShowHdfsFileListRequestOrder
	ASC  ShowHdfsFileListRequestOrder
}

func GetShowHdfsFileListRequestOrderEnum() ShowHdfsFileListRequestOrderEnum {
	return ShowHdfsFileListRequestOrderEnum{
		DESC: ShowHdfsFileListRequestOrder{
			value: "desc",
		},
		ASC: ShowHdfsFileListRequestOrder{
			value: "asc",
		},
	}
}

func (c ShowHdfsFileListRequestOrder) Value() string {
	return c.value
}

func (c ShowHdfsFileListRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHdfsFileListRequestOrder) UnmarshalJSON(b []byte) error {
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
