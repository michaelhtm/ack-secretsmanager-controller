//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIErrorType) DeepCopyInto(out *APIErrorType) {
	*out = *in
	if in.SecretID != nil {
		in, out := &in.SecretID, &out.SecretID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIErrorType.
func (in *APIErrorType) DeepCopy() *APIErrorType {
	if in == nil {
		return nil
	}
	out := new(APIErrorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaRegionType) DeepCopyInto(out *ReplicaRegionType) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaRegionType.
func (in *ReplicaRegionType) DeepCopy() *ReplicaRegionType {
	if in == nil {
		return nil
	}
	out := new(ReplicaRegionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationStatusType) DeepCopyInto(out *ReplicationStatusType) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = (*in).DeepCopy()
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationStatusType.
func (in *ReplicationStatusType) DeepCopy() *ReplicationStatusType {
	if in == nil {
		return nil
	}
	out := new(ReplicationStatusType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RotationRulesType) DeepCopyInto(out *RotationRulesType) {
	*out = *in
	if in.AutomaticallyAfterDays != nil {
		in, out := &in.AutomaticallyAfterDays, &out.AutomaticallyAfterDays
		*out = new(int64)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.ScheduleExpression != nil {
		in, out := &in.ScheduleExpression, &out.ScheduleExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RotationRulesType.
func (in *RotationRulesType) DeepCopy() *RotationRulesType {
	if in == nil {
		return nil
	}
	out := new(RotationRulesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Secret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretList) DeepCopyInto(out *SecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Secret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretList.
func (in *SecretList) DeepCopy() *SecretList {
	if in == nil {
		return nil
	}
	out := new(SecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretListEntry) DeepCopyInto(out *SecretListEntry) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreatedDate != nil {
		in, out := &in.CreatedDate, &out.CreatedDate
		*out = (*in).DeepCopy()
	}
	if in.DeletedDate != nil {
		in, out := &in.DeletedDate, &out.DeletedDate
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = (*in).DeepCopy()
	}
	if in.LastChangedDate != nil {
		in, out := &in.LastChangedDate, &out.LastChangedDate
		*out = (*in).DeepCopy()
	}
	if in.LastRotatedDate != nil {
		in, out := &in.LastRotatedDate, &out.LastRotatedDate
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NextRotationDate != nil {
		in, out := &in.NextRotationDate, &out.NextRotationDate
		*out = (*in).DeepCopy()
	}
	if in.OwningService != nil {
		in, out := &in.OwningService, &out.OwningService
		*out = new(string)
		**out = **in
	}
	if in.PrimaryRegion != nil {
		in, out := &in.PrimaryRegion, &out.PrimaryRegion
		*out = new(string)
		**out = **in
	}
	if in.RotationEnabled != nil {
		in, out := &in.RotationEnabled, &out.RotationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RotationLambdaARN != nil {
		in, out := &in.RotationLambdaARN, &out.RotationLambdaARN
		*out = new(string)
		**out = **in
	}
	if in.RotationRules != nil {
		in, out := &in.RotationRules, &out.RotationRules
		*out = new(RotationRulesType)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretVersionsToStages != nil {
		in, out := &in.SecretVersionsToStages, &out.SecretVersionsToStages
		*out = make(map[string][]*string, len(*in))
		for key, val := range *in {
			var outVal []*string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]*string, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(string)
						**out = **in
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretListEntry.
func (in *SecretListEntry) DeepCopy() *SecretListEntry {
	if in == nil {
		return nil
	}
	out := new(SecretListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceOverwriteReplicaSecret != nil {
		in, out := &in.ForceOverwriteReplicaSecret, &out.ForceOverwriteReplicaSecret
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ReplicaRegions != nil {
		in, out := &in.ReplicaRegions, &out.ReplicaRegions
		*out = make([]*ReplicaRegionType, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicaRegionType)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.SecretString != nil {
		in, out := &in.SecretString, &out.SecretString
		*out = new(corev1alpha1.SecretKeyReference)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatus) DeepCopyInto(out *SecretStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ReplicationStatus != nil {
		in, out := &in.ReplicationStatus, &out.ReplicationStatus
		*out = make([]*ReplicationStatusType, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicationStatusType)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatus.
func (in *SecretStatus) DeepCopy() *SecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueEntry) DeepCopyInto(out *SecretValueEntry) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SecretBinary != nil {
		in, out := &in.SecretBinary, &out.SecretBinary
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SecretString != nil {
		in, out := &in.SecretString, &out.SecretString
		*out = new(string)
		**out = **in
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
	if in.VersionStages != nil {
		in, out := &in.VersionStages, &out.VersionStages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueEntry.
func (in *SecretValueEntry) DeepCopy() *SecretValueEntry {
	if in == nil {
		return nil
	}
	out := new(SecretValueEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVersionsListEntry) DeepCopyInto(out *SecretVersionsListEntry) {
	*out = *in
	if in.LastAccessedDate != nil {
		in, out := &in.LastAccessedDate, &out.LastAccessedDate
		*out = (*in).DeepCopy()
	}
	if in.VersionID != nil {
		in, out := &in.VersionID, &out.VersionID
		*out = new(string)
		**out = **in
	}
	if in.VersionStages != nil {
		in, out := &in.VersionStages, &out.VersionStages
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVersionsListEntry.
func (in *SecretVersionsListEntry) DeepCopy() *SecretVersionsListEntry {
	if in == nil {
		return nil
	}
	out := new(SecretVersionsListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationErrorsEntry) DeepCopyInto(out *ValidationErrorsEntry) {
	*out = *in
	if in.CheckName != nil {
		in, out := &in.CheckName, &out.CheckName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationErrorsEntry.
func (in *ValidationErrorsEntry) DeepCopy() *ValidationErrorsEntry {
	if in == nil {
		return nil
	}
	out := new(ValidationErrorsEntry)
	in.DeepCopyInto(out)
	return out
}
