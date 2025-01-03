package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyUnbindInstancesApi
type EcsBackupPolicyUnbindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupPolicyUnbindInstancesApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyUnbindInstancesApi {
	return &EcsBackupPolicyUnbindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/unbind-instances",
		},
	}
}

func (this *EcsBackupPolicyUnbindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyUnbindInstancesRequest) (*EcsBackupPolicyUnbindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupPolicyUnbindInstancesRealRequest{
		RegionID:       req.RegionID,
		PolicyID:       req.PolicyID,
		InstanceIDList: req.InstanceIDList,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupPolicyUnbindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupPolicyUnbindInstancesResponse{
		PolicyID:       realResponse.PolicyID,
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type EcsBackupPolicyUnbindInstancesRealRequest struct {
	RegionID       *string `json:"regionID,omitempty"`
	PolicyID       *string `json:"policyID,omitempty"`
	InstanceIDList *string `json:"instanceIDList,omitempty"`
}

type EcsBackupPolicyUnbindInstancesRequest struct {
	RegionID       *string
	PolicyID       *string
	InstanceIDList *string
}

type EcsBackupPolicyUnbindInstancesRealResponse struct {
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type EcsBackupPolicyUnbindInstancesResponse struct {
	PolicyID       string
	InstanceIDList string
}
