// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package route

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.Route{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetRouteWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.ApiKeyRequired != nil {
		ko.Spec.APIKeyRequired = resp.ApiKeyRequired
	}
	if resp.AuthorizationScopes != nil {
		f2 := []*string{}
		for _, f2iter := range resp.AuthorizationScopes {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		ko.Spec.AuthorizationScopes = f2
	}
	if resp.AuthorizationType != nil {
		ko.Spec.AuthorizationType = resp.AuthorizationType
	}
	if resp.AuthorizerId != nil {
		ko.Spec.AuthorizerID = resp.AuthorizerId
	}
	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	}
	if resp.OperationName != nil {
		ko.Spec.OperationName = resp.OperationName
	}
	if resp.RequestModels != nil {
		f7 := map[string]*string{}
		for f7key, f7valiter := range resp.RequestModels {
			var f7val string
			f7val = *f7valiter
			f7[f7key] = &f7val
		}
		ko.Spec.RequestModels = f7
	}
	if resp.RequestParameters != nil {
		f8 := map[string]*svcapitypes.ParameterConstraints{}
		for f8key, f8valiter := range resp.RequestParameters {
			f8val := &svcapitypes.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.Required = f8valiter.Required
			}
			f8[f8key] = f8val
		}
		ko.Spec.RequestParameters = f8
	}
	if resp.RouteId != nil {
		ko.Status.RouteID = resp.RouteId
	}
	if resp.RouteKey != nil {
		ko.Spec.RouteKey = resp.RouteKey
	}
	if resp.RouteResponseSelectionExpression != nil {
		ko.Spec.RouteResponseSelectionExpression = resp.RouteResponseSelectionExpression
	}
	if resp.Target != nil {
		ko.Spec.Target = resp.Target
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.APIID == nil || r.ko.Status.RouteID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetRouteInput, error) {
	res := &svcsdk.GetRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.RouteID != nil {
		res.SetRouteId(*r.ko.Status.RouteID)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.GetRoutesInput, error) {
	res := &svcsdk.GetRoutesInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateRouteWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.RouteId != nil {
		ko.Status.RouteID = resp.RouteId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateRouteInput, error) {
	res := &svcsdk.CreateRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.APIKeyRequired != nil {
		res.SetApiKeyRequired(*r.ko.Spec.APIKeyRequired)
	}
	if r.ko.Spec.AuthorizationScopes != nil {
		f2 := []*string{}
		for _, f2iter := range r.ko.Spec.AuthorizationScopes {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		res.SetAuthorizationScopes(f2)
	}
	if r.ko.Spec.AuthorizationType != nil {
		res.SetAuthorizationType(*r.ko.Spec.AuthorizationType)
	}
	if r.ko.Spec.AuthorizerID != nil {
		res.SetAuthorizerId(*r.ko.Spec.AuthorizerID)
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.SetModelSelectionExpression(*r.ko.Spec.ModelSelectionExpression)
	}
	if r.ko.Spec.OperationName != nil {
		res.SetOperationName(*r.ko.Spec.OperationName)
	}
	if r.ko.Spec.RequestModels != nil {
		f7 := map[string]*string{}
		for f7key, f7valiter := range r.ko.Spec.RequestModels {
			var f7val string
			f7val = *f7valiter
			f7[f7key] = &f7val
		}
		res.SetRequestModels(f7)
	}
	if r.ko.Spec.RequestParameters != nil {
		f8 := map[string]*svcsdk.ParameterConstraints{}
		for f8key, f8valiter := range r.ko.Spec.RequestParameters {
			f8val := &svcsdk.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.SetRequired(*f8valiter.Required)
			}
			f8[f8key] = f8val
		}
		res.SetRequestParameters(f8)
	}
	if r.ko.Spec.RouteKey != nil {
		res.SetRouteKey(*r.ko.Spec.RouteKey)
	}
	if r.ko.Spec.RouteResponseSelectionExpression != nil {
		res.SetRouteResponseSelectionExpression(*r.ko.Spec.RouteResponseSelectionExpression)
	}
	if r.ko.Spec.Target != nil {
		res.SetTarget(*r.ko.Spec.Target)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateRouteWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.RouteId != nil {
		ko.Status.RouteID = resp.RouteId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateRouteInput, error) {
	res := &svcsdk.UpdateRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.APIKeyRequired != nil {
		res.SetApiKeyRequired(*r.ko.Spec.APIKeyRequired)
	}
	if r.ko.Spec.AuthorizationScopes != nil {
		f2 := []*string{}
		for _, f2iter := range r.ko.Spec.AuthorizationScopes {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		res.SetAuthorizationScopes(f2)
	}
	if r.ko.Spec.AuthorizationType != nil {
		res.SetAuthorizationType(*r.ko.Spec.AuthorizationType)
	}
	if r.ko.Spec.AuthorizerID != nil {
		res.SetAuthorizerId(*r.ko.Spec.AuthorizerID)
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.SetModelSelectionExpression(*r.ko.Spec.ModelSelectionExpression)
	}
	if r.ko.Spec.OperationName != nil {
		res.SetOperationName(*r.ko.Spec.OperationName)
	}
	if r.ko.Spec.RequestModels != nil {
		f7 := map[string]*string{}
		for f7key, f7valiter := range r.ko.Spec.RequestModels {
			var f7val string
			f7val = *f7valiter
			f7[f7key] = &f7val
		}
		res.SetRequestModels(f7)
	}
	if r.ko.Spec.RequestParameters != nil {
		f8 := map[string]*svcsdk.ParameterConstraints{}
		for f8key, f8valiter := range r.ko.Spec.RequestParameters {
			f8val := &svcsdk.ParameterConstraints{}
			if f8valiter.Required != nil {
				f8val.SetRequired(*f8valiter.Required)
			}
			f8[f8key] = f8val
		}
		res.SetRequestParameters(f8)
	}
	if r.ko.Status.RouteID != nil {
		res.SetRouteId(*r.ko.Status.RouteID)
	}
	if r.ko.Spec.RouteKey != nil {
		res.SetRouteKey(*r.ko.Spec.RouteKey)
	}
	if r.ko.Spec.RouteResponseSelectionExpression != nil {
		res.SetRouteResponseSelectionExpression(*r.ko.Spec.RouteResponseSelectionExpression)
	}
	if r.ko.Spec.Target != nil {
		res.SetTarget(*r.ko.Spec.Target)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteRouteWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRouteInput, error) {
	res := &svcsdk.DeleteRouteInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.RouteID != nil {
		res.SetRouteId(*r.ko.Status.RouteID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Route,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}
