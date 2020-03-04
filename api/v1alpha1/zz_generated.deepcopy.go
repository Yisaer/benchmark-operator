// +build !ignore_autogenerated

/*

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataBaseBenchmarkPrepare) DeepCopyInto(out *DataBaseBenchmarkPrepare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataBaseBenchmarkPrepare.
func (in *DataBaseBenchmarkPrepare) DeepCopy() *DataBaseBenchmarkPrepare {
	if in == nil {
		return nil
	}
	out := new(DataBaseBenchmarkPrepare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataBaseBenchmarkPrepare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataBaseBenchmarkPrepareList) DeepCopyInto(out *DataBaseBenchmarkPrepareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataBaseBenchmarkPrepare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataBaseBenchmarkPrepareList.
func (in *DataBaseBenchmarkPrepareList) DeepCopy() *DataBaseBenchmarkPrepareList {
	if in == nil {
		return nil
	}
	out := new(DataBaseBenchmarkPrepareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataBaseBenchmarkPrepareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataBaseBenchmarkPrepareSpec) DeepCopyInto(out *DataBaseBenchmarkPrepareSpec) {
	*out = *in
	if in.Prepares != nil {
		in, out := &in.Prepares, &out.Prepares
		*out = make([]Prepare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataBaseBenchmarkPrepareSpec.
func (in *DataBaseBenchmarkPrepareSpec) DeepCopy() *DataBaseBenchmarkPrepareSpec {
	if in == nil {
		return nil
	}
	out := new(DataBaseBenchmarkPrepareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataBaseBenchmarkPrepareStatus) DeepCopyInto(out *DataBaseBenchmarkPrepareStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataBaseBenchmarkPrepareStatus.
func (in *DataBaseBenchmarkPrepareStatus) DeepCopy() *DataBaseBenchmarkPrepareStatus {
	if in == nil {
		return nil
	}
	out := new(DataBaseBenchmarkPrepareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prepare) DeepCopyInto(out *Prepare) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prepare.
func (in *Prepare) DeepCopy() *Prepare {
	if in == nil {
		return nil
	}
	out := new(Prepare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TidbClusterRef) DeepCopyInto(out *TidbClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TidbClusterRef.
func (in *TidbClusterRef) DeepCopy() *TidbClusterRef {
	if in == nil {
		return nil
	}
	out := new(TidbClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TpccBenchmark) DeepCopyInto(out *TpccBenchmark) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TpccBenchmark.
func (in *TpccBenchmark) DeepCopy() *TpccBenchmark {
	if in == nil {
		return nil
	}
	out := new(TpccBenchmark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TpccBenchmark) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TpccBenchmarkList) DeepCopyInto(out *TpccBenchmarkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TpccBenchmark, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TpccBenchmarkList.
func (in *TpccBenchmarkList) DeepCopy() *TpccBenchmarkList {
	if in == nil {
		return nil
	}
	out := new(TpccBenchmarkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TpccBenchmarkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TpccBenchmarkSpec) DeepCopyInto(out *TpccBenchmarkSpec) {
	*out = *in
	if in.Conn != nil {
		in, out := &in.Conn, &out.Conn
		*out = new(string)
		**out = **in
	}
	out.Cluster = in.Cluster
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TpccBenchmarkSpec.
func (in *TpccBenchmarkSpec) DeepCopy() *TpccBenchmarkSpec {
	if in == nil {
		return nil
	}
	out := new(TpccBenchmarkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TpccBenchmarkStatus) DeepCopyInto(out *TpccBenchmarkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TpccBenchmarkStatus.
func (in *TpccBenchmarkStatus) DeepCopy() *TpccBenchmarkStatus {
	if in == nil {
		return nil
	}
	out := new(TpccBenchmarkStatus)
	in.DeepCopyInto(out)
	return out
}
