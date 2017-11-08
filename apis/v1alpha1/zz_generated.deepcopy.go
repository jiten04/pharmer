// +build !ignore_autogenerated

/*
Copyright 2017 The Pharmer Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	reflect "reflect"

	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*API).DeepCopyInto(out.(*API))
			return nil
		}, InType: reflect.TypeOf(&API{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AWSSpec).DeepCopyInto(out.(*AWSSpec))
			return nil
		}, InType: reflect.TypeOf(&AWSSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AWSStatus).DeepCopyInto(out.(*AWSStatus))
			return nil
		}, InType: reflect.TypeOf(&AWSStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Action).DeepCopyInto(out.(*Action))
			return nil
		}, InType: reflect.TypeOf(&Action{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureCloudConfig).DeepCopyInto(out.(*AzureCloudConfig))
			return nil
		}, InType: reflect.TypeOf(&AzureCloudConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureSpec).DeepCopyInto(out.(*AzureSpec))
			return nil
		}, InType: reflect.TypeOf(&AzureSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureStorageSpec).DeepCopyInto(out.(*AzureStorageSpec))
			return nil
		}, InType: reflect.TypeOf(&AzureStorageSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CloudSpec).DeepCopyInto(out.(*CloudSpec))
			return nil
		}, InType: reflect.TypeOf(&CloudSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CloudStatus).DeepCopyInto(out.(*CloudStatus))
			return nil
		}, InType: reflect.TypeOf(&CloudStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Cluster).DeepCopyInto(out.(*Cluster))
			return nil
		}, InType: reflect.TypeOf(&Cluster{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterDeleteRequest).DeepCopyInto(out.(*ClusterDeleteRequest))
			return nil
		}, InType: reflect.TypeOf(&ClusterDeleteRequest{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterSpec).DeepCopyInto(out.(*ClusterSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterState).DeepCopyInto(out.(*ClusterState))
			return nil
		}, InType: reflect.TypeOf(&ClusterState{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterStatus).DeepCopyInto(out.(*ClusterStatus))
			return nil
		}, InType: reflect.TypeOf(&ClusterStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Credential).DeepCopyInto(out.(*Credential))
			return nil
		}, InType: reflect.TypeOf(&Credential{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CredentialSpec).DeepCopyInto(out.(*CredentialSpec))
			return nil
		}, InType: reflect.TypeOf(&CredentialSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GCECloudConfig).DeepCopyInto(out.(*GCECloudConfig))
			return nil
		}, InType: reflect.TypeOf(&GCECloudConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GCSSpec).DeepCopyInto(out.(*GCSSpec))
			return nil
		}, InType: reflect.TypeOf(&GCSSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GoogleSpec).DeepCopyInto(out.(*GoogleSpec))
			return nil
		}, InType: reflect.TypeOf(&GoogleSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LinodeCloudConfig).DeepCopyInto(out.(*LinodeCloudConfig))
			return nil
		}, InType: reflect.TypeOf(&LinodeCloudConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LinodeSpec).DeepCopyInto(out.(*LinodeSpec))
			return nil
		}, InType: reflect.TypeOf(&LinodeSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LocalSpec).DeepCopyInto(out.(*LocalSpec))
			return nil
		}, InType: reflect.TypeOf(&LocalSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Networking).DeepCopyInto(out.(*Networking))
			return nil
		}, InType: reflect.TypeOf(&Networking{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeGroup).DeepCopyInto(out.(*NodeGroup))
			return nil
		}, InType: reflect.TypeOf(&NodeGroup{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeGroupSpec).DeepCopyInto(out.(*NodeGroupSpec))
			return nil
		}, InType: reflect.TypeOf(&NodeGroupSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeGroupStatus).DeepCopyInto(out.(*NodeGroupStatus))
			return nil
		}, InType: reflect.TypeOf(&NodeGroupStatus{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeInfo).DeepCopyInto(out.(*NodeInfo))
			return nil
		}, InType: reflect.TypeOf(&NodeInfo{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeSpec).DeepCopyInto(out.(*NodeSpec))
			return nil
		}, InType: reflect.TypeOf(&NodeSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NodeTemplateSpec).DeepCopyInto(out.(*NodeTemplateSpec))
			return nil
		}, InType: reflect.TypeOf(&NodeTemplateSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*NullNameGenerator).DeepCopyInto(out.(*NullNameGenerator))
			return nil
		}, InType: reflect.TypeOf(&NullNameGenerator{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PacketCloudConfig).DeepCopyInto(out.(*PacketCloudConfig))
			return nil
		}, InType: reflect.TypeOf(&PacketCloudConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PharmerConfig).DeepCopyInto(out.(*PharmerConfig))
			return nil
		}, InType: reflect.TypeOf(&PharmerConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ReservedIP).DeepCopyInto(out.(*ReservedIP))
			return nil
		}, InType: reflect.TypeOf(&ReservedIP{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*S3Spec).DeepCopyInto(out.(*S3Spec))
			return nil
		}, InType: reflect.TypeOf(&S3Spec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SSHConfig).DeepCopyInto(out.(*SSHConfig))
			return nil
		}, InType: reflect.TypeOf(&SSHConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ScalewayCloudConfig).DeepCopyInto(out.(*ScalewayCloudConfig))
			return nil
		}, InType: reflect.TypeOf(&ScalewayCloudConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SoftlayerCloudConfig).DeepCopyInto(out.(*SoftlayerCloudConfig))
			return nil
		}, InType: reflect.TypeOf(&SoftlayerCloudConfig{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StorageBackend).DeepCopyInto(out.(*StorageBackend))
			return nil
		}, InType: reflect.TypeOf(&StorageBackend{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SwiftSpec).DeepCopyInto(out.(*SwiftSpec))
			return nil
		}, InType: reflect.TypeOf(&SwiftSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Upgrade).DeepCopyInto(out.(*Upgrade))
			return nil
		}, InType: reflect.TypeOf(&Upgrade{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VultrCloudConfig).DeepCopyInto(out.(*VultrCloudConfig))
			return nil
		}, InType: reflect.TypeOf(&VultrCloudConfig{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSpec) DeepCopyInto(out *AWSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSpec.
func (in *AWSSpec) DeepCopy() *AWSSpec {
	if in == nil {
		return nil
	}
	out := new(AWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSStatus) DeepCopyInto(out *AWSStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSStatus.
func (in *AWSStatus) DeepCopy() *AWSStatus {
	if in == nil {
		return nil
	}
	out := new(AWSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Action) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCloudConfig) DeepCopyInto(out *AzureCloudConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCloudConfig.
func (in *AzureCloudConfig) DeepCopy() *AzureCloudConfig {
	if in == nil {
		return nil
	}
	out := new(AzureCloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSpec) DeepCopyInto(out *AzureSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSpec.
func (in *AzureSpec) DeepCopy() *AzureSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureStorageSpec) DeepCopyInto(out *AzureStorageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureStorageSpec.
func (in *AzureStorageSpec) DeepCopy() *AzureStorageSpec {
	if in == nil {
		return nil
	}
	out := new(AzureStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSpec) DeepCopyInto(out *CloudSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSSpec)
			**out = **in
		}
	}
	if in.GCE != nil {
		in, out := &in.GCE, &out.GCE
		if *in == nil {
			*out = nil
		} else {
			*out = new(GoogleSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		if *in == nil {
			*out = nil
		} else {
			*out = new(AzureSpec)
			**out = **in
		}
	}
	if in.Linode != nil {
		in, out := &in.Linode, &out.Linode
		if *in == nil {
			*out = nil
		} else {
			*out = new(LinodeSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSpec.
func (in *CloudSpec) DeepCopy() *CloudSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStatus) DeepCopyInto(out *CloudStatus) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStatus.
func (in *CloudStatus) DeepCopy() *CloudStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeleteRequest) DeepCopyInto(out *ClusterDeleteRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeleteRequest.
func (in *ClusterDeleteRequest) DeepCopy() *ClusterDeleteRequest {
	if in == nil {
		return nil
	}
	out := new(ClusterDeleteRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.Cloud.DeepCopyInto(&out.Cloud)
	out.API = in.API
	out.Networking = in.Networking
	if in.KubeletExtraArgs != nil {
		in, out := &in.KubeletExtraArgs, &out.KubeletExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIServerExtraArgs != nil {
		in, out := &in.APIServerExtraArgs, &out.APIServerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ControllerManagerExtraArgs != nil {
		in, out := &in.ControllerManagerExtraArgs, &out.ControllerManagerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SchedulerExtraArgs != nil {
		in, out := &in.SchedulerExtraArgs, &out.SchedulerExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AuthorizationModes != nil {
		in, out := &in.AuthorizationModes, &out.AuthorizationModes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIServerCertSANs != nil {
		in, out := &in.APIServerCertSANs, &out.APIServerCertSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterState) DeepCopyInto(out *ClusterState) {
	*out = *in
	if in.KubeletVersions != nil {
		in, out := &in.KubeletVersions, &out.KubeletVersions
		*out = make(map[string]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterState.
func (in *ClusterState) DeepCopy() *ClusterState {
	if in == nil {
		return nil
	}
	out := new(ClusterState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.Cloud.DeepCopyInto(&out.Cloud)
	if in.APIAddresses != nil {
		in, out := &in.APIAddresses, &out.APIAddresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.ReservedIPs != nil {
		in, out := &in.ReservedIPs, &out.ReservedIPs
		*out = make([]ReservedIP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Credential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialSpec) DeepCopyInto(out *CredentialSpec) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialSpec.
func (in *CredentialSpec) DeepCopy() *CredentialSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCECloudConfig) DeepCopyInto(out *GCECloudConfig) {
	*out = *in
	if in.NodeTags != nil {
		in, out := &in.NodeTags, &out.NodeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCECloudConfig.
func (in *GCECloudConfig) DeepCopy() *GCECloudConfig {
	if in == nil {
		return nil
	}
	out := new(GCECloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSSpec) DeepCopyInto(out *GCSSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSSpec.
func (in *GCSSpec) DeepCopy() *GCSSpec {
	if in == nil {
		return nil
	}
	out := new(GCSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleSpec) DeepCopyInto(out *GoogleSpec) {
	*out = *in
	if in.NodeTags != nil {
		in, out := &in.NodeTags, &out.NodeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeScopes != nil {
		in, out := &in.NodeScopes, &out.NodeScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleSpec.
func (in *GoogleSpec) DeepCopy() *GoogleSpec {
	if in == nil {
		return nil
	}
	out := new(GoogleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinodeCloudConfig) DeepCopyInto(out *LinodeCloudConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinodeCloudConfig.
func (in *LinodeCloudConfig) DeepCopy() *LinodeCloudConfig {
	if in == nil {
		return nil
	}
	out := new(LinodeCloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinodeSpec) DeepCopyInto(out *LinodeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinodeSpec.
func (in *LinodeSpec) DeepCopy() *LinodeSpec {
	if in == nil {
		return nil
	}
	out := new(LinodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpec) DeepCopyInto(out *LocalSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpec.
func (in *LocalSpec) DeepCopy() *LocalSpec {
	if in == nil {
		return nil
	}
	out := new(LocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroup) DeepCopyInto(out *NodeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroup.
func (in *NodeGroup) DeepCopy() *NodeGroup {
	if in == nil {
		return nil
	}
	out := new(NodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupSpec) DeepCopyInto(out *NodeGroupSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupSpec.
func (in *NodeGroupSpec) DeepCopy() *NodeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(NodeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeGroupStatus) DeepCopyInto(out *NodeGroupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeGroupStatus.
func (in *NodeGroupStatus) DeepCopy() *NodeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(NodeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInfo) DeepCopyInto(out *NodeInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInfo.
func (in *NodeInfo) DeepCopy() *NodeInfo {
	if in == nil {
		return nil
	}
	out := new(NodeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	if in.KubeletExtraArgs != nil {
		in, out := &in.KubeletExtraArgs, &out.KubeletExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTemplateSpec) DeepCopyInto(out *NodeTemplateSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTemplateSpec.
func (in *NodeTemplateSpec) DeepCopy() *NodeTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(NodeTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NullNameGenerator) DeepCopyInto(out *NullNameGenerator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NullNameGenerator.
func (in *NullNameGenerator) DeepCopy() *NullNameGenerator {
	if in == nil {
		return nil
	}
	out := new(NullNameGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PacketCloudConfig) DeepCopyInto(out *PacketCloudConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PacketCloudConfig.
func (in *PacketCloudConfig) DeepCopy() *PacketCloudConfig {
	if in == nil {
		return nil
	}
	out := new(PacketCloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PharmerConfig) DeepCopyInto(out *PharmerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]Credential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Store.DeepCopyInto(&out.Store)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PharmerConfig.
func (in *PharmerConfig) DeepCopy() *PharmerConfig {
	if in == nil {
		return nil
	}
	out := new(PharmerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PharmerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedIP) DeepCopyInto(out *ReservedIP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedIP.
func (in *ReservedIP) DeepCopy() *ReservedIP {
	if in == nil {
		return nil
	}
	out := new(ReservedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Spec) DeepCopyInto(out *S3Spec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Spec.
func (in *S3Spec) DeepCopy() *S3Spec {
	if in == nil {
		return nil
	}
	out := new(S3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHConfig) DeepCopyInto(out *SSHConfig) {
	*out = *in
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHConfig.
func (in *SSHConfig) DeepCopy() *SSHConfig {
	if in == nil {
		return nil
	}
	out := new(SSHConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalewayCloudConfig) DeepCopyInto(out *ScalewayCloudConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalewayCloudConfig.
func (in *ScalewayCloudConfig) DeepCopy() *ScalewayCloudConfig {
	if in == nil {
		return nil
	}
	out := new(ScalewayCloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftlayerCloudConfig) DeepCopyInto(out *SoftlayerCloudConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftlayerCloudConfig.
func (in *SoftlayerCloudConfig) DeepCopy() *SoftlayerCloudConfig {
	if in == nil {
		return nil
	}
	out := new(SoftlayerCloudConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBackend) DeepCopyInto(out *StorageBackend) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		if *in == nil {
			*out = nil
		} else {
			*out = new(LocalSpec)
			**out = **in
		}
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		if *in == nil {
			*out = nil
		} else {
			*out = new(S3Spec)
			**out = **in
		}
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		if *in == nil {
			*out = nil
		} else {
			*out = new(GCSSpec)
			**out = **in
		}
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		if *in == nil {
			*out = nil
		} else {
			*out = new(AzureStorageSpec)
			**out = **in
		}
	}
	if in.Swift != nil {
		in, out := &in.Swift, &out.Swift
		if *in == nil {
			*out = nil
		} else {
			*out = new(SwiftSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBackend.
func (in *StorageBackend) DeepCopy() *StorageBackend {
	if in == nil {
		return nil
	}
	out := new(StorageBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftSpec) DeepCopyInto(out *SwiftSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftSpec.
func (in *SwiftSpec) DeepCopy() *SwiftSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Upgrade) DeepCopyInto(out *Upgrade) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Before.DeepCopyInto(&out.Before)
	in.After.DeepCopyInto(&out.After)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Upgrade.
func (in *Upgrade) DeepCopy() *Upgrade {
	if in == nil {
		return nil
	}
	out := new(Upgrade)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Upgrade) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VultrCloudConfig) DeepCopyInto(out *VultrCloudConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VultrCloudConfig.
func (in *VultrCloudConfig) DeepCopy() *VultrCloudConfig {
	if in == nil {
		return nil
	}
	out := new(VultrCloudConfig)
	in.DeepCopyInto(out)
	return out
}
