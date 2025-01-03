package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupUnbindInstancesApi 云主机组批量移除云主机
// https://www.ctyun.cn/document/10026730/10463455
type EcsAffinityGroupUnbindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupUnbindInstancesApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupUnbindInstancesApi {
	return &EcsAffinityGroupUnbindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/unbind-instances",
		},
	}
}

func (this *EcsAffinityGroupUnbindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupUnbindInstancesRequest) (*EcsAffinityGroupUnbindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupUnbindInstancesRealRequest{
		RegionID:        req.RegionID,
		InstanceIDs:     req.InstanceIDs,
		AffinityGroupID: req.AffinityGroupID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupUnbindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupUnbindInstancesResponse{}, nil
}

type EcsAffinityGroupUnbindInstancesRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	InstanceIDs     *string `json:"instanceIDs,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupUnbindInstancesRequest struct {
	RegionID        *string
	InstanceIDs     *string
	AffinityGroupID *string
}

type EcsAffinityGroupUnbindInstancesRealResponse struct {
}

type EcsAffinityGroupUnbindInstancesResponse struct {
}
