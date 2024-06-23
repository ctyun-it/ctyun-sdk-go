package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAgentBatchActionApi
type EcsAgentBatchActionApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAgentBatchActionApi(client *ctyunsdk.CtyunClient) *EcsAgentBatchActionApi {
	return &EcsAgentBatchActionApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/agent/batch-action",
		},
	}
}

func (this *EcsAgentBatchActionApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAgentBatchActionRequest) (*EcsAgentBatchActionResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var actionInfo []EcsAgentBatchActionActionInfoRealRequest
	for _, request := range req.ActionInfo {
		actionInfo = append(actionInfo, EcsAgentBatchActionActionInfoRealRequest{
			InstanceID:    request.InstanceID,
			SystemType:    request.SystemType,
			SystemArch:    request.SystemArch,
			SystemVersion: request.SystemVersion,
		})
	}

	_, err := builder.WriteJson(&EcsAgentBatchActionRealRequest{
		RegionID:   req.RegionID,
		Action:     req.Action,
		ActionInfo: actionInfo,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAgentBatchActionRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAgentBatchActionResponse{}, nil
}

type EcsAgentBatchActionActionInfoRealRequest struct {
	InstanceID    string `json:"instanceID,omitempty"`
	SystemType    string `json:"systemType,omitempty"`
	SystemArch    string `json:"systemArch,omitempty"`
	SystemVersion string `json:"systemVersion,omitempty"`
}

type EcsAgentBatchActionRealRequest struct {
	RegionID   string                                     `json:"regionID,omitempty"`
	Action     string                                     `json:"action,omitempty"`
	ActionInfo []EcsAgentBatchActionActionInfoRealRequest `json:"actionInfo,omitempty"`
}

type EcsAgentBatchActionActionInfoRequest struct {
	InstanceID    string
	SystemType    string
	SystemArch    string
	SystemVersion string
}

type EcsAgentBatchActionRequest struct {
	RegionID   string
	Action     string
	ActionInfo []EcsAgentBatchActionActionInfoRequest
}

type EcsAgentBatchActionRealResponse struct {
}

type EcsAgentBatchActionResponse struct {
}
