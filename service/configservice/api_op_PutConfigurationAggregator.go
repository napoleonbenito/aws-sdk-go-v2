// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutConfigurationAggregatorRequest
type PutConfigurationAggregatorInput struct {
	_ struct{} `type:"structure"`

	// A list of AccountAggregationSource object.
	AccountAggregationSources []AccountAggregationSource `type:"list"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`

	// An OrganizationAggregationSource object.
	OrganizationAggregationSource *OrganizationAggregationSource `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationAggregatorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationAggregatorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationAggregatorInput"}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}
	if s.AccountAggregationSources != nil {
		for i, v := range s.AccountAggregationSources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AccountAggregationSources", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.OrganizationAggregationSource != nil {
		if err := s.OrganizationAggregationSource.Validate(); err != nil {
			invalidParams.AddNested("OrganizationAggregationSource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutConfigurationAggregatorResponse
type PutConfigurationAggregatorOutput struct {
	_ struct{} `type:"structure"`

	// Returns a ConfigurationAggregator object.
	ConfigurationAggregator *ConfigurationAggregator `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationAggregatorOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutConfigurationAggregator = "PutConfigurationAggregator"

// PutConfigurationAggregatorRequest returns a request value for making API operation for
// AWS Config.
//
// Creates and updates the configuration aggregator with the selected source
// accounts and regions. The source account can be individual account(s) or
// an organization.
//
// AWS Config should be enabled in source accounts and regions you want to aggregate.
//
// If your source type is an organization, you must be signed in to the master
// account and all features must be enabled in your organization. AWS Config
// calls EnableAwsServiceAccess API to enable integration between AWS Config
// and AWS Organizations.
//
//    // Example sending a request using PutConfigurationAggregatorRequest.
//    req := client.PutConfigurationAggregatorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutConfigurationAggregator
func (c *Client) PutConfigurationAggregatorRequest(input *PutConfigurationAggregatorInput) PutConfigurationAggregatorRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationAggregator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutConfigurationAggregatorInput{}
	}

	req := c.newRequest(op, input, &PutConfigurationAggregatorOutput{})
	return PutConfigurationAggregatorRequest{Request: req, Input: input, Copy: c.PutConfigurationAggregatorRequest}
}

// PutConfigurationAggregatorRequest is the request type for the
// PutConfigurationAggregator API operation.
type PutConfigurationAggregatorRequest struct {
	*aws.Request
	Input *PutConfigurationAggregatorInput
	Copy  func(*PutConfigurationAggregatorInput) PutConfigurationAggregatorRequest
}

// Send marshals and sends the PutConfigurationAggregator API request.
func (r PutConfigurationAggregatorRequest) Send(ctx context.Context) (*PutConfigurationAggregatorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationAggregatorResponse{
		PutConfigurationAggregatorOutput: r.Request.Data.(*PutConfigurationAggregatorOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationAggregatorResponse is the response type for the
// PutConfigurationAggregator API operation.
type PutConfigurationAggregatorResponse struct {
	*PutConfigurationAggregatorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationAggregator request.
func (r *PutConfigurationAggregatorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}