package v3

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/vpc/v3/model"
)

type VpcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVpcClient(hcClient *http_client.HcHttpClient) *VpcClient {
	return &VpcClient{HcClient: hcClient}
}

func VpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *VpcClient) AddVpcExtendCidr(request *model.AddVpcExtendCidrRequest) (*model.AddVpcExtendCidrResponse, error) {
	requestDef := GenReqDefForAddVpcExtendCidr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVpcExtendCidrResponse), nil
	}
}

func (c *VpcClient) AddVpcExtendCidrInvoker(request *model.AddVpcExtendCidrRequest) *AddVpcExtendCidrInvoker {
	requestDef := GenReqDefForAddVpcExtendCidr()
	return &AddVpcExtendCidrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResponse), nil
	}
}

func (c *VpcClient) CreateVpcInvoker(request *model.CreateVpcRequest) *CreateVpcInvoker {
	requestDef := GenReqDefForCreateVpc()
	return &CreateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcResponse), nil
	}
}

func (c *VpcClient) DeleteVpcInvoker(request *model.DeleteVpcRequest) *DeleteVpcInvoker {
	requestDef := GenReqDefForDeleteVpc()
	return &DeleteVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsResponse), nil
	}
}

func (c *VpcClient) ListVpcsInvoker(request *model.ListVpcsRequest) *ListVpcsInvoker {
	requestDef := GenReqDefForListVpcs()
	return &ListVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) RemoveVpcExtendCidr(request *model.RemoveVpcExtendCidrRequest) (*model.RemoveVpcExtendCidrResponse, error) {
	requestDef := GenReqDefForRemoveVpcExtendCidr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveVpcExtendCidrResponse), nil
	}
}

func (c *VpcClient) RemoveVpcExtendCidrInvoker(request *model.RemoveVpcExtendCidrRequest) *RemoveVpcExtendCidrInvoker {
	requestDef := GenReqDefForRemoveVpcExtendCidr()
	return &RemoveVpcExtendCidrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcResponse), nil
	}
}

func (c *VpcClient) ShowVpcInvoker(request *model.ShowVpcRequest) *ShowVpcInvoker {
	requestDef := GenReqDefForShowVpc()
	return &ShowVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcResponse), nil
	}
}

func (c *VpcClient) UpdateVpcInvoker(request *model.UpdateVpcRequest) *UpdateVpcInvoker {
	requestDef := GenReqDefForUpdateVpc()
	return &UpdateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
