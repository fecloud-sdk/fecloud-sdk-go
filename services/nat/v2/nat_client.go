package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/nat/v2/model"
)

type NatClient struct {
	HcClient *http_client.HcHttpClient
}

func NewNatClient(hcClient *http_client.HcHttpClient) *NatClient {
	return &NatClient{HcClient: hcClient}
}

func NatClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *NatClient) BatchCreateNatGatewayDnatRules(request *model.BatchCreateNatGatewayDnatRulesRequest) (*model.BatchCreateNatGatewayDnatRulesResponse, error) {
	requestDef := GenReqDefForBatchCreateNatGatewayDnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateNatGatewayDnatRulesResponse), nil
	}
}

func (c *NatClient) BatchCreateNatGatewayDnatRulesInvoker(request *model.BatchCreateNatGatewayDnatRulesRequest) *BatchCreateNatGatewayDnatRulesInvoker {
	requestDef := GenReqDefForBatchCreateNatGatewayDnatRules()
	return &BatchCreateNatGatewayDnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateNatGatewayDnatRule(request *model.CreateNatGatewayDnatRuleRequest) (*model.CreateNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForCreateNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) CreateNatGatewayDnatRuleInvoker(request *model.CreateNatGatewayDnatRuleRequest) *CreateNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForCreateNatGatewayDnatRule()
	return &CreateNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreatePrivateDnat(request *model.CreatePrivateDnatRequest) (*model.CreatePrivateDnatResponse, error) {
	requestDef := GenReqDefForCreatePrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateDnatResponse), nil
	}
}

func (c *NatClient) CreatePrivateDnatInvoker(request *model.CreatePrivateDnatRequest) *CreatePrivateDnatInvoker {
	requestDef := GenReqDefForCreatePrivateDnat()
	return &CreatePrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteNatGatewayDnatRule(request *model.DeleteNatGatewayDnatRuleRequest) (*model.DeleteNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) DeleteNatGatewayDnatRuleInvoker(request *model.DeleteNatGatewayDnatRuleRequest) *DeleteNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForDeleteNatGatewayDnatRule()
	return &DeleteNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeletePrivateDnat(request *model.DeletePrivateDnatRequest) (*model.DeletePrivateDnatResponse, error) {
	requestDef := GenReqDefForDeletePrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateDnatResponse), nil
	}
}

func (c *NatClient) DeletePrivateDnatInvoker(request *model.DeletePrivateDnatRequest) *DeletePrivateDnatInvoker {
	requestDef := GenReqDefForDeletePrivateDnat()
	return &DeletePrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListNatGatewayDnatRules(request *model.ListNatGatewayDnatRulesRequest) (*model.ListNatGatewayDnatRulesResponse, error) {
	requestDef := GenReqDefForListNatGatewayDnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewayDnatRulesResponse), nil
	}
}

func (c *NatClient) ListNatGatewayDnatRulesInvoker(request *model.ListNatGatewayDnatRulesRequest) *ListNatGatewayDnatRulesInvoker {
	requestDef := GenReqDefForListNatGatewayDnatRules()
	return &ListNatGatewayDnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListPrivateDnats(request *model.ListPrivateDnatsRequest) (*model.ListPrivateDnatsResponse, error) {
	requestDef := GenReqDefForListPrivateDnats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateDnatsResponse), nil
	}
}

func (c *NatClient) ListPrivateDnatsInvoker(request *model.ListPrivateDnatsRequest) *ListPrivateDnatsInvoker {
	requestDef := GenReqDefForListPrivateDnats()
	return &ListPrivateDnatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowNatGatewayDnatRule(request *model.ShowNatGatewayDnatRuleRequest) (*model.ShowNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForShowNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) ShowNatGatewayDnatRuleInvoker(request *model.ShowNatGatewayDnatRuleRequest) *ShowNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForShowNatGatewayDnatRule()
	return &ShowNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowPrivateDnat(request *model.ShowPrivateDnatRequest) (*model.ShowPrivateDnatResponse, error) {
	requestDef := GenReqDefForShowPrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateDnatResponse), nil
	}
}

func (c *NatClient) ShowPrivateDnatInvoker(request *model.ShowPrivateDnatRequest) *ShowPrivateDnatInvoker {
	requestDef := GenReqDefForShowPrivateDnat()
	return &ShowPrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdateNatGatewayDnatRule(request *model.UpdateNatGatewayDnatRuleRequest) (*model.UpdateNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayDnatRuleResponse), nil
	}
}

func (c *NatClient) UpdateNatGatewayDnatRuleInvoker(request *model.UpdateNatGatewayDnatRuleRequest) *UpdateNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForUpdateNatGatewayDnatRule()
	return &UpdateNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdatePrivateDnat(request *model.UpdatePrivateDnatRequest) (*model.UpdatePrivateDnatResponse, error) {
	requestDef := GenReqDefForUpdatePrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateDnatResponse), nil
	}
}

func (c *NatClient) UpdatePrivateDnatInvoker(request *model.UpdatePrivateDnatRequest) *UpdatePrivateDnatInvoker {
	requestDef := GenReqDefForUpdatePrivateDnat()
	return &UpdatePrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) BatchCreateDeleteTransitIpTags(request *model.BatchCreateDeleteTransitIpTagsRequest) (*model.BatchCreateDeleteTransitIpTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeleteTransitIpTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeleteTransitIpTagsResponse), nil
	}
}

func (c *NatClient) BatchCreateDeleteTransitIpTagsInvoker(request *model.BatchCreateDeleteTransitIpTagsRequest) *BatchCreateDeleteTransitIpTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeleteTransitIpTags()
	return &BatchCreateDeleteTransitIpTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateTransitIpTag(request *model.CreateTransitIpTagRequest) (*model.CreateTransitIpTagResponse, error) {
	requestDef := GenReqDefForCreateTransitIpTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransitIpTagResponse), nil
	}
}

func (c *NatClient) CreateTransitIpTagInvoker(request *model.CreateTransitIpTagRequest) *CreateTransitIpTagInvoker {
	requestDef := GenReqDefForCreateTransitIpTag()
	return &CreateTransitIpTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteTransitIpTag(request *model.DeleteTransitIpTagRequest) (*model.DeleteTransitIpTagResponse, error) {
	requestDef := GenReqDefForDeleteTransitIpTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransitIpTagResponse), nil
	}
}

func (c *NatClient) DeleteTransitIpTagInvoker(request *model.DeleteTransitIpTagRequest) *DeleteTransitIpTagInvoker {
	requestDef := GenReqDefForDeleteTransitIpTag()
	return &DeleteTransitIpTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListTransitIpTags(request *model.ListTransitIpTagsRequest) (*model.ListTransitIpTagsResponse, error) {
	requestDef := GenReqDefForListTransitIpTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitIpTagsResponse), nil
	}
}

func (c *NatClient) ListTransitIpTagsInvoker(request *model.ListTransitIpTagsRequest) *ListTransitIpTagsInvoker {
	requestDef := GenReqDefForListTransitIpTags()
	return &ListTransitIpTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListTransitIpsByTags(request *model.ListTransitIpsByTagsRequest) (*model.ListTransitIpsByTagsResponse, error) {
	requestDef := GenReqDefForListTransitIpsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitIpsByTagsResponse), nil
	}
}

func (c *NatClient) ListTransitIpsByTagsInvoker(request *model.ListTransitIpsByTagsRequest) *ListTransitIpsByTagsInvoker {
	requestDef := GenReqDefForListTransitIpsByTags()
	return &ListTransitIpsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowTransitIpTags(request *model.ShowTransitIpTagsRequest) (*model.ShowTransitIpTagsResponse, error) {
	requestDef := GenReqDefForShowTransitIpTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransitIpTagsResponse), nil
	}
}

func (c *NatClient) ShowTransitIpTagsInvoker(request *model.ShowTransitIpTagsRequest) *ShowTransitIpTagsInvoker {
	requestDef := GenReqDefForShowTransitIpTags()
	return &ShowTransitIpTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) BatchCreateDeletePrivateNatTags(request *model.BatchCreateDeletePrivateNatTagsRequest) (*model.BatchCreateDeletePrivateNatTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeletePrivateNatTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeletePrivateNatTagsResponse), nil
	}
}

func (c *NatClient) BatchCreateDeletePrivateNatTagsInvoker(request *model.BatchCreateDeletePrivateNatTagsRequest) *BatchCreateDeletePrivateNatTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeletePrivateNatTags()
	return &BatchCreateDeletePrivateNatTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateNatGateway(request *model.CreateNatGatewayRequest) (*model.CreateNatGatewayResponse, error) {
	requestDef := GenReqDefForCreateNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayResponse), nil
	}
}

func (c *NatClient) CreateNatGatewayInvoker(request *model.CreateNatGatewayRequest) *CreateNatGatewayInvoker {
	requestDef := GenReqDefForCreateNatGateway()
	return &CreateNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreatePrivateNat(request *model.CreatePrivateNatRequest) (*model.CreatePrivateNatResponse, error) {
	requestDef := GenReqDefForCreatePrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateNatResponse), nil
	}
}

func (c *NatClient) CreatePrivateNatInvoker(request *model.CreatePrivateNatRequest) *CreatePrivateNatInvoker {
	requestDef := GenReqDefForCreatePrivateNat()
	return &CreatePrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreatePrivateNatTag(request *model.CreatePrivateNatTagRequest) (*model.CreatePrivateNatTagResponse, error) {
	requestDef := GenReqDefForCreatePrivateNatTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateNatTagResponse), nil
	}
}

func (c *NatClient) CreatePrivateNatTagInvoker(request *model.CreatePrivateNatTagRequest) *CreatePrivateNatTagInvoker {
	requestDef := GenReqDefForCreatePrivateNatTag()
	return &CreatePrivateNatTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteNatGateway(request *model.DeleteNatGatewayRequest) (*model.DeleteNatGatewayResponse, error) {
	requestDef := GenReqDefForDeleteNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayResponse), nil
	}
}

func (c *NatClient) DeleteNatGatewayInvoker(request *model.DeleteNatGatewayRequest) *DeleteNatGatewayInvoker {
	requestDef := GenReqDefForDeleteNatGateway()
	return &DeleteNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeletePrivateNat(request *model.DeletePrivateNatRequest) (*model.DeletePrivateNatResponse, error) {
	requestDef := GenReqDefForDeletePrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateNatResponse), nil
	}
}

func (c *NatClient) DeletePrivateNatInvoker(request *model.DeletePrivateNatRequest) *DeletePrivateNatInvoker {
	requestDef := GenReqDefForDeletePrivateNat()
	return &DeletePrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeletePrivateNatTag(request *model.DeletePrivateNatTagRequest) (*model.DeletePrivateNatTagResponse, error) {
	requestDef := GenReqDefForDeletePrivateNatTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateNatTagResponse), nil
	}
}

func (c *NatClient) DeletePrivateNatTagInvoker(request *model.DeletePrivateNatTagRequest) *DeletePrivateNatTagInvoker {
	requestDef := GenReqDefForDeletePrivateNatTag()
	return &DeletePrivateNatTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListNatGateways(request *model.ListNatGatewaysRequest) (*model.ListNatGatewaysResponse, error) {
	requestDef := GenReqDefForListNatGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaysResponse), nil
	}
}

func (c *NatClient) ListNatGatewaysInvoker(request *model.ListNatGatewaysRequest) *ListNatGatewaysInvoker {
	requestDef := GenReqDefForListNatGateways()
	return &ListNatGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListPrivateNatTags(request *model.ListPrivateNatTagsRequest) (*model.ListPrivateNatTagsResponse, error) {
	requestDef := GenReqDefForListPrivateNatTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNatTagsResponse), nil
	}
}

func (c *NatClient) ListPrivateNatTagsInvoker(request *model.ListPrivateNatTagsRequest) *ListPrivateNatTagsInvoker {
	requestDef := GenReqDefForListPrivateNatTags()
	return &ListPrivateNatTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListPrivateNats(request *model.ListPrivateNatsRequest) (*model.ListPrivateNatsResponse, error) {
	requestDef := GenReqDefForListPrivateNats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNatsResponse), nil
	}
}

func (c *NatClient) ListPrivateNatsInvoker(request *model.ListPrivateNatsRequest) *ListPrivateNatsInvoker {
	requestDef := GenReqDefForListPrivateNats()
	return &ListPrivateNatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListPrivateNatsByTags(request *model.ListPrivateNatsByTagsRequest) (*model.ListPrivateNatsByTagsResponse, error) {
	requestDef := GenReqDefForListPrivateNatsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNatsByTagsResponse), nil
	}
}

func (c *NatClient) ListPrivateNatsByTagsInvoker(request *model.ListPrivateNatsByTagsRequest) *ListPrivateNatsByTagsInvoker {
	requestDef := GenReqDefForListPrivateNatsByTags()
	return &ListPrivateNatsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowNatGateway(request *model.ShowNatGatewayRequest) (*model.ShowNatGatewayResponse, error) {
	requestDef := GenReqDefForShowNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayResponse), nil
	}
}

func (c *NatClient) ShowNatGatewayInvoker(request *model.ShowNatGatewayRequest) *ShowNatGatewayInvoker {
	requestDef := GenReqDefForShowNatGateway()
	return &ShowNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowPrivateNat(request *model.ShowPrivateNatRequest) (*model.ShowPrivateNatResponse, error) {
	requestDef := GenReqDefForShowPrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateNatResponse), nil
	}
}

func (c *NatClient) ShowPrivateNatInvoker(request *model.ShowPrivateNatRequest) *ShowPrivateNatInvoker {
	requestDef := GenReqDefForShowPrivateNat()
	return &ShowPrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowPrivateNatTags(request *model.ShowPrivateNatTagsRequest) (*model.ShowPrivateNatTagsResponse, error) {
	requestDef := GenReqDefForShowPrivateNatTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateNatTagsResponse), nil
	}
}

func (c *NatClient) ShowPrivateNatTagsInvoker(request *model.ShowPrivateNatTagsRequest) *ShowPrivateNatTagsInvoker {
	requestDef := GenReqDefForShowPrivateNatTags()
	return &ShowPrivateNatTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdateNatGateway(request *model.UpdateNatGatewayRequest) (*model.UpdateNatGatewayResponse, error) {
	requestDef := GenReqDefForUpdateNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayResponse), nil
	}
}

func (c *NatClient) UpdateNatGatewayInvoker(request *model.UpdateNatGatewayRequest) *UpdateNatGatewayInvoker {
	requestDef := GenReqDefForUpdateNatGateway()
	return &UpdateNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdatePrivateNat(request *model.UpdatePrivateNatRequest) (*model.UpdatePrivateNatResponse, error) {
	requestDef := GenReqDefForUpdatePrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateNatResponse), nil
	}
}

func (c *NatClient) UpdatePrivateNatInvoker(request *model.UpdatePrivateNatRequest) *UpdatePrivateNatInvoker {
	requestDef := GenReqDefForUpdatePrivateNat()
	return &UpdatePrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateTransitIp(request *model.CreateTransitIpRequest) (*model.CreateTransitIpResponse, error) {
	requestDef := GenReqDefForCreateTransitIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransitIpResponse), nil
	}
}

func (c *NatClient) CreateTransitIpInvoker(request *model.CreateTransitIpRequest) *CreateTransitIpInvoker {
	requestDef := GenReqDefForCreateTransitIp()
	return &CreateTransitIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteTransitIp(request *model.DeleteTransitIpRequest) (*model.DeleteTransitIpResponse, error) {
	requestDef := GenReqDefForDeleteTransitIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransitIpResponse), nil
	}
}

func (c *NatClient) DeleteTransitIpInvoker(request *model.DeleteTransitIpRequest) *DeleteTransitIpInvoker {
	requestDef := GenReqDefForDeleteTransitIp()
	return &DeleteTransitIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListTransitIps(request *model.ListTransitIpsRequest) (*model.ListTransitIpsResponse, error) {
	requestDef := GenReqDefForListTransitIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitIpsResponse), nil
	}
}

func (c *NatClient) ListTransitIpsInvoker(request *model.ListTransitIpsRequest) *ListTransitIpsInvoker {
	requestDef := GenReqDefForListTransitIps()
	return &ListTransitIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowTransitIp(request *model.ShowTransitIpRequest) (*model.ShowTransitIpResponse, error) {
	requestDef := GenReqDefForShowTransitIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransitIpResponse), nil
	}
}

func (c *NatClient) ShowTransitIpInvoker(request *model.ShowTransitIpRequest) *ShowTransitIpInvoker {
	requestDef := GenReqDefForShowTransitIp()
	return &ShowTransitIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreateNatGatewaySnatRule(request *model.CreateNatGatewaySnatRuleRequest) (*model.CreateNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForCreateNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) CreateNatGatewaySnatRuleInvoker(request *model.CreateNatGatewaySnatRuleRequest) *CreateNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForCreateNatGatewaySnatRule()
	return &CreateNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) CreatePrivateSnat(request *model.CreatePrivateSnatRequest) (*model.CreatePrivateSnatResponse, error) {
	requestDef := GenReqDefForCreatePrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateSnatResponse), nil
	}
}

func (c *NatClient) CreatePrivateSnatInvoker(request *model.CreatePrivateSnatRequest) *CreatePrivateSnatInvoker {
	requestDef := GenReqDefForCreatePrivateSnat()
	return &CreatePrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeleteNatGatewaySnatRule(request *model.DeleteNatGatewaySnatRuleRequest) (*model.DeleteNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) DeleteNatGatewaySnatRuleInvoker(request *model.DeleteNatGatewaySnatRuleRequest) *DeleteNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForDeleteNatGatewaySnatRule()
	return &DeleteNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) DeletePrivateSnat(request *model.DeletePrivateSnatRequest) (*model.DeletePrivateSnatResponse, error) {
	requestDef := GenReqDefForDeletePrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateSnatResponse), nil
	}
}

func (c *NatClient) DeletePrivateSnatInvoker(request *model.DeletePrivateSnatRequest) *DeletePrivateSnatInvoker {
	requestDef := GenReqDefForDeletePrivateSnat()
	return &DeletePrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListNatGatewaySnatRules(request *model.ListNatGatewaySnatRulesRequest) (*model.ListNatGatewaySnatRulesResponse, error) {
	requestDef := GenReqDefForListNatGatewaySnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaySnatRulesResponse), nil
	}
}

func (c *NatClient) ListNatGatewaySnatRulesInvoker(request *model.ListNatGatewaySnatRulesRequest) *ListNatGatewaySnatRulesInvoker {
	requestDef := GenReqDefForListNatGatewaySnatRules()
	return &ListNatGatewaySnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ListPrivateSnats(request *model.ListPrivateSnatsRequest) (*model.ListPrivateSnatsResponse, error) {
	requestDef := GenReqDefForListPrivateSnats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateSnatsResponse), nil
	}
}

func (c *NatClient) ListPrivateSnatsInvoker(request *model.ListPrivateSnatsRequest) *ListPrivateSnatsInvoker {
	requestDef := GenReqDefForListPrivateSnats()
	return &ListPrivateSnatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowNatGatewaySnatRule(request *model.ShowNatGatewaySnatRuleRequest) (*model.ShowNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForShowNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) ShowNatGatewaySnatRuleInvoker(request *model.ShowNatGatewaySnatRuleRequest) *ShowNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForShowNatGatewaySnatRule()
	return &ShowNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) ShowPrivateSnat(request *model.ShowPrivateSnatRequest) (*model.ShowPrivateSnatResponse, error) {
	requestDef := GenReqDefForShowPrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateSnatResponse), nil
	}
}

func (c *NatClient) ShowPrivateSnatInvoker(request *model.ShowPrivateSnatRequest) *ShowPrivateSnatInvoker {
	requestDef := GenReqDefForShowPrivateSnat()
	return &ShowPrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdateNatGatewaySnatRule(request *model.UpdateNatGatewaySnatRuleRequest) (*model.UpdateNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewaySnatRuleResponse), nil
	}
}

func (c *NatClient) UpdateNatGatewaySnatRuleInvoker(request *model.UpdateNatGatewaySnatRuleRequest) *UpdateNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForUpdateNatGatewaySnatRule()
	return &UpdateNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *NatClient) UpdatePrivateSnat(request *model.UpdatePrivateSnatRequest) (*model.UpdatePrivateSnatResponse, error) {
	requestDef := GenReqDefForUpdatePrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateSnatResponse), nil
	}
}

func (c *NatClient) UpdatePrivateSnatInvoker(request *model.UpdatePrivateSnatRequest) *UpdatePrivateSnatInvoker {
	requestDef := GenReqDefForUpdatePrivateSnat()
	return &UpdatePrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
