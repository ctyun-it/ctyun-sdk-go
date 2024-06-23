package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// InstanceBackupPolicyUnbindInstancesApi
type InstanceBackupPolicyUnbindInstancesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewInstanceBackupPolicyUnbindInstancesApi(client *ctyunsdk.CtyunClient) *InstanceBackupPolicyUnbindInstancesApi {
	return &InstanceBackupPolicyUnbindInstancesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/unbind-instances",
		},
	}
}

func (this *InstanceBackupPolicyUnbindInstancesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *InstanceBackupPolicyUnbindInstancesRequest) (*InstanceBackupPolicyUnbindInstancesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&InstanceBackupPolicyUnbindInstancesRealRequest{
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

	var realResponse InstanceBackupPolicyUnbindInstancesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &InstanceBackupPolicyUnbindInstancesResponse{
		PolicyID:       realResponse.PolicyID,
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type InstanceBackupPolicyUnbindInstancesRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type InstanceBackupPolicyUnbindInstancesRequest struct {
	RegionID       string
	PolicyID       string
	InstanceIDList string
}

type InstanceBackupPolicyUnbindInstancesRealResponse struct {
	PolicyID       string `json:"policyID,omitempty"`
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type InstanceBackupPolicyUnbindInstancesResponse struct {
	PolicyID       string
	InstanceIDList string
}
