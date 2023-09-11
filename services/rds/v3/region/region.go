package region

import (
	"fmt"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/region"
)

var (
	EU_WEST_0 = region.NewRegion("eu-west-0",
		"https://rds.eu-west-0.prod-cloud-ocb.orange-business.com")
)

var staticFields = map[string]*region.Region{
	"eu-west-0": EU_WEST_0,
}

var provider = region.DefaultProviderChain("RDS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
