package v1

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/sfsturbo/v1/model"
)

type SFSTurboClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSFSTurboClient(hcClient *http_client.HcHttpClient) *SFSTurboClient {
	return &SFSTurboClient{HcClient: hcClient}
}

func SFSTurboClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *SFSTurboClient) BatchAddSharedTags(request *model.BatchAddSharedTagsRequest) (*model.BatchAddSharedTagsResponse, error) {
	requestDef := GenReqDefForBatchAddSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddSharedTagsResponse), nil
	}
}

func (c *SFSTurboClient) BatchAddSharedTagsInvoker(request *model.BatchAddSharedTagsRequest) *BatchAddSharedTagsInvoker {
	requestDef := GenReqDefForBatchAddSharedTags()
	return &BatchAddSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ChangeSecurityGroup(request *model.ChangeSecurityGroupRequest) (*model.ChangeSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSecurityGroupResponse), nil
	}
}

func (c *SFSTurboClient) ChangeSecurityGroupInvoker(request *model.ChangeSecurityGroupRequest) *ChangeSecurityGroupInvoker {
	requestDef := GenReqDefForChangeSecurityGroup()
	return &ChangeSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ChangeShareName(request *model.ChangeShareNameRequest) (*model.ChangeShareNameResponse, error) {
	requestDef := GenReqDefForChangeShareName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeShareNameResponse), nil
	}
}

func (c *SFSTurboClient) ChangeShareNameInvoker(request *model.ChangeShareNameRequest) *ChangeShareNameInvoker {
	requestDef := GenReqDefForChangeShareName()
	return &ChangeShareNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) CreateShare(request *model.CreateShareRequest) (*model.CreateShareResponse, error) {
	requestDef := GenReqDefForCreateShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShareResponse), nil
	}
}

func (c *SFSTurboClient) CreateShareInvoker(request *model.CreateShareRequest) *CreateShareInvoker {
	requestDef := GenReqDefForCreateShare()
	return &CreateShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) CreateSharedTag(request *model.CreateSharedTagRequest) (*model.CreateSharedTagResponse, error) {
	requestDef := GenReqDefForCreateSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSharedTagResponse), nil
	}
}

func (c *SFSTurboClient) CreateSharedTagInvoker(request *model.CreateSharedTagRequest) *CreateSharedTagInvoker {
	requestDef := GenReqDefForCreateSharedTag()
	return &CreateSharedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) DeleteShare(request *model.DeleteShareRequest) (*model.DeleteShareResponse, error) {
	requestDef := GenReqDefForDeleteShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShareResponse), nil
	}
}

func (c *SFSTurboClient) DeleteShareInvoker(request *model.DeleteShareRequest) *DeleteShareInvoker {
	requestDef := GenReqDefForDeleteShare()
	return &DeleteShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) DeleteSharedTag(request *model.DeleteSharedTagRequest) (*model.DeleteSharedTagResponse, error) {
	requestDef := GenReqDefForDeleteSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSharedTagResponse), nil
	}
}

func (c *SFSTurboClient) DeleteSharedTagInvoker(request *model.DeleteSharedTagRequest) *DeleteSharedTagInvoker {
	requestDef := GenReqDefForDeleteSharedTag()
	return &DeleteSharedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ExpandShare(request *model.ExpandShareRequest) (*model.ExpandShareResponse, error) {
	requestDef := GenReqDefForExpandShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandShareResponse), nil
	}
}

func (c *SFSTurboClient) ExpandShareInvoker(request *model.ExpandShareRequest) *ExpandShareInvoker {
	requestDef := GenReqDefForExpandShare()
	return &ExpandShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ListSharedTags(request *model.ListSharedTagsRequest) (*model.ListSharedTagsResponse, error) {
	requestDef := GenReqDefForListSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedTagsResponse), nil
	}
}

func (c *SFSTurboClient) ListSharedTagsInvoker(request *model.ListSharedTagsRequest) *ListSharedTagsInvoker {
	requestDef := GenReqDefForListSharedTags()
	return &ListSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ListShares(request *model.ListSharesRequest) (*model.ListSharesResponse, error) {
	requestDef := GenReqDefForListShares()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharesResponse), nil
	}
}

func (c *SFSTurboClient) ListSharesInvoker(request *model.ListSharesRequest) *ListSharesInvoker {
	requestDef := GenReqDefForListShares()
	return &ListSharesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ShowShare(request *model.ShowShareRequest) (*model.ShowShareResponse, error) {
	requestDef := GenReqDefForShowShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShareResponse), nil
	}
}

func (c *SFSTurboClient) ShowShareInvoker(request *model.ShowShareRequest) *ShowShareInvoker {
	requestDef := GenReqDefForShowShare()
	return &ShowShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *SFSTurboClient) ShowSharedTags(request *model.ShowSharedTagsRequest) (*model.ShowSharedTagsResponse, error) {
	requestDef := GenReqDefForShowSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSharedTagsResponse), nil
	}
}

func (c *SFSTurboClient) ShowSharedTagsInvoker(request *model.ShowSharedTagsRequest) *ShowSharedTagsInvoker {
	requestDef := GenReqDefForShowSharedTags()
	return &ShowSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
