//go:build !ignore_autogenerated

/*
Copyright 2025.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbSpecs) DeepCopyInto(out *DbSpecs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbSpecs.
func (in *DbSpecs) DeepCopy() *DbSpecs {
	if in == nil {
		return nil
	}
	out := new(DbSpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbWatcher) DeepCopyInto(out *DbWatcher) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbWatcher.
func (in *DbWatcher) DeepCopy() *DbWatcher {
	if in == nil {
		return nil
	}
	out := new(DbWatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbWatcher) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbWatcherList) DeepCopyInto(out *DbWatcherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DbWatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbWatcherList.
func (in *DbWatcherList) DeepCopy() *DbWatcherList {
	if in == nil {
		return nil
	}
	out := new(DbWatcherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DbWatcherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbWatcherSpec) DeepCopyInto(out *DbWatcherSpec) {
	*out = *in
	out.DbSpec = in.DbSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbWatcherSpec.
func (in *DbWatcherSpec) DeepCopy() *DbWatcherSpec {
	if in == nil {
		return nil
	}
	out := new(DbWatcherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbWatcherStatus) DeepCopyInto(out *DbWatcherStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbWatcherStatus.
func (in *DbWatcherStatus) DeepCopy() *DbWatcherStatus {
	if in == nil {
		return nil
	}
	out := new(DbWatcherStatus)
	in.DeepCopyInto(out)
	return out
}
