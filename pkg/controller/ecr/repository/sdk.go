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

package repository

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/ecr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/ecr/apis/v1alpha1"
	"github.com/muvaf/test-generated-aws/apis/ecr/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ECR{}
	_ = &svcapitypes.Repository{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)
func GenerateDescribeRepositoriesInput(cr *v1alpha1.Repository) *svcsdk.DescribeRepositoriesInput {
	res := &svcsdk.DescribeRepositoriesInput{}

	if cr.Status.AtProvider.RegistryID != nil {
		res.SetRegistryId(*cr.Status.AtProvider.RegistryID)
	}

	return res
}


// GenerateCreateRepositoryInput returns a CreateRepositoryInput object.
func GenerateCreateRepositoryInput(cr *v1alpha1.Repository) *svcsdk.CreateRepositoryInput {
	res := &svcsdk.CreateRepositoryInput{}

	if cr.Spec.ForProvider.EncryptionConfiguration != nil {
		f0 := &svcsdk.EncryptionConfiguration{}
		if cr.Spec.ForProvider.EncryptionConfiguration.EncryptionType != nil {
			f0.SetEncryptionType(*cr.Spec.ForProvider.EncryptionConfiguration.EncryptionType)
		}
		if cr.Spec.ForProvider.EncryptionConfiguration.KMSKey != nil {
			f0.SetKmsKey(*cr.Spec.ForProvider.EncryptionConfiguration.KMSKey)
		}
		res.SetEncryptionConfiguration(f0)
	}
	if cr.Spec.ForProvider.ImageScanningConfiguration != nil {
		f1 := &svcsdk.ImageScanningConfiguration{}
		if cr.Spec.ForProvider.ImageScanningConfiguration.ScanOnPush != nil {
			f1.SetScanOnPush(*cr.Spec.ForProvider.ImageScanningConfiguration.ScanOnPush)
		}
		res.SetImageScanningConfiguration(f1)
	}
	if cr.Spec.ForProvider.ImageTagMutability != nil {
		res.SetImageTagMutability(*cr.Spec.ForProvider.ImageTagMutability)
	}
	if cr.Spec.ForProvider.RepositoryName != nil {
		res.SetRepositoryName(*cr.Spec.ForProvider.RepositoryName)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range cr.Spec.ForProvider.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
	}

	return res
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func GenerateDeleteRepositoryInput(cr *v1alpha1.Repository) *svcsdk.DeleteRepositoryInput {
	res := &svcsdk.DeleteRepositoryInput{}

	if cr.Status.AtProvider.RegistryID != nil {
		res.SetRegistryId(*cr.Status.AtProvider.RegistryID)
	}
	if cr.Spec.ForProvider.RepositoryName != nil {
		res.SetRepositoryName(*cr.Spec.ForProvider.RepositoryName)
	}

	return res
}// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeRepositoriesWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if len(resp.Repositories) == 0 {
		return nil, ackerr.NotFound
	}
	found := false
	for _, elem := range resp.Repositories {
		if elem.CreatedAt != nil {
			ko.Status.AtProvider.CreatedAt = &metav1.Time{*elem.CreatedAt}
		}
		if elem.EncryptionConfiguration != nil {
			f1 := &svcapitypes.EncryptionConfiguration{}
			if elem.EncryptionConfiguration.EncryptionType != nil {
				f1.EncryptionType = elem.EncryptionConfiguration.EncryptionType
			}
			if elem.EncryptionConfiguration.KmsKey != nil {
				f1.KMSKey = elem.EncryptionConfiguration.KmsKey
			}
			ko.Spec.ForProvider.EncryptionConfiguration = f1
		}
		if elem.ImageScanningConfiguration != nil {
			f2 := &svcapitypes.ImageScanningConfiguration{}
			if elem.ImageScanningConfiguration.ScanOnPush != nil {
				f2.ScanOnPush = elem.ImageScanningConfiguration.ScanOnPush
			}
			ko.Spec.ForProvider.ImageScanningConfiguration = f2
		}
		if elem.ImageTagMutability != nil {
			ko.Spec.ForProvider.ImageTagMutability = elem.ImageTagMutability
		}
		if elem.RegistryId != nil {
			ko.Status.AtProvider.RegistryID = elem.RegistryId
		}
		if elem.RepositoryArn != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.RepositoryArn)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.RepositoryName != nil {
			ko.Spec.ForProvider.RepositoryName = elem.RepositoryName
		}
		if elem.RepositoryUri != nil {
			ko.Status.AtProvider.RepositoryURI = elem.RepositoryUri
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	return &resource{ko}, nil
}
// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeRepositoriesInput, error) {
	res := &svcsdk.DescribeRepositoriesInput{}

	if r.ko.Status.AtProvider.RegistryID != nil {
		res.SetRegistryId(*r.ko.Status.AtProvider.RegistryID)
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

	resp, respErr := rm.sdkapi.CreateRepositoryWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Repository.CreatedAt != nil {
		ko.Status.CreatedAt = &metav1.Time{*resp.Repository.CreatedAt}
	}
	if resp.Repository.RegistryId != nil {
		ko.Status.RegistryID = resp.Repository.RegistryId
	}
	if resp.Repository.RepositoryUri != nil {
		ko.Status.RepositoryURI = resp.Repository.RepositoryUri
	}

	ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{OwnerAccountID: &rm.awsAccountID}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateRepositoryInput, error) {
	res := &svcsdk.CreateRepositoryInput{}

	if r.ko.Spec.ForProvider.EncryptionConfiguration != nil {
		f0 := &svcsdk.EncryptionConfiguration{}
		if r.ko.Spec.ForProvider.EncryptionConfiguration.EncryptionType != nil {
			f0.SetEncryptionType(*r.ko.Spec.ForProvider.EncryptionConfiguration.EncryptionType)
		}
		if r.ko.Spec.ForProvider.EncryptionConfiguration.KMSKey != nil {
			f0.SetKmsKey(*r.ko.Spec.ForProvider.EncryptionConfiguration.KMSKey)
		}
		res.SetEncryptionConfiguration(f0)
	}
	if r.ko.Spec.ForProvider.ImageScanningConfiguration != nil {
		f1 := &svcsdk.ImageScanningConfiguration{}
		if r.ko.Spec.ForProvider.ImageScanningConfiguration.ScanOnPush != nil {
			f1.SetScanOnPush(*r.ko.Spec.ForProvider.ImageScanningConfiguration.ScanOnPush)
		}
		res.SetImageScanningConfiguration(f1)
	}
	if r.ko.Spec.ForProvider.ImageTagMutability != nil {
		res.SetImageTagMutability(*r.ko.Spec.ForProvider.ImageTagMutability)
	}
	if r.ko.Spec.ForProvider.RepositoryName != nil {
		res.SetRepositoryName(*r.ko.Spec.ForProvider.RepositoryName)
	}
	if r.ko.Spec.ForProvider.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range r.ko.Spec.ForProvider.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, nil
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
	_, respErr := rm.sdkapi.DeleteRepositoryWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRepositoryInput, error) {
	res := &svcsdk.DeleteRepositoryInput{}

	if r.ko.Status.AtProvider.RegistryID != nil {
		res.SetRegistryId(*r.ko.Status.AtProvider.RegistryID)
	}
	if r.ko.Spec.ForProvider.RepositoryName != nil {
		res.SetRepositoryName(*r.ko.Spec.ForProvider.RepositoryName)
	}

	return res, nil
}