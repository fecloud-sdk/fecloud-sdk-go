package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CertInfo struct {
	Category CertInfoCategory `json:"category"`

	DownloadLink string `json:"download_link"`
}

func (o CertInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertInfo struct{}"
	}

	return strings.Join([]string{"CertInfo", string(data)}, " ")
}

type CertInfoCategory struct {
	value string
}

type CertInfoCategoryEnum struct {
	INTERNATIONAL CertInfoCategory
	NATIONAL      CertInfoCategory
}

func GetCertInfoCategoryEnum() CertInfoCategoryEnum {
	return CertInfoCategoryEnum{
		INTERNATIONAL: CertInfoCategory{
			value: "international",
		},
		NATIONAL: CertInfoCategory{
			value: "national",
		},
	}
}

func (c CertInfoCategory) Value() string {
	return c.value
}

func (c CertInfoCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertInfoCategory) UnmarshalJSON(b []byte) error {
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
