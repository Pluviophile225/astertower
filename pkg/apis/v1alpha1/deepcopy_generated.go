//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// astertower
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Astro) DeepCopyInto(out *Astro) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Astro.
func (in *Astro) DeepCopy() *Astro {
	if in == nil {
		return nil
	}
	out := new(Astro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Astro) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstroCondition) DeepCopyInto(out *AstroCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstroCondition.
func (in *AstroCondition) DeepCopy() *AstroCondition {
	if in == nil {
		return nil
	}
	out := new(AstroCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstroList) DeepCopyInto(out *AstroList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Astro, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstroList.
func (in *AstroList) DeepCopy() *AstroList {
	if in == nil {
		return nil
	}
	out := new(AstroList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AstroList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstroRef) DeepCopyInto(out *AstroRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstroRef.
func (in *AstroRef) DeepCopy() *AstroRef {
	if in == nil {
		return nil
	}
	out := new(AstroRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstroSpec) DeepCopyInto(out *AstroSpec) {
	*out = *in
	if in.Stars != nil {
		in, out := &in.Stars, &out.Stars
		*out = make([]AstroStar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstroSpec.
func (in *AstroSpec) DeepCopy() *AstroSpec {
	if in == nil {
		return nil
	}
	out := new(AstroSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstroStar) DeepCopyInto(out *AstroStar) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstroStar.
func (in *AstroStar) DeepCopy() *AstroStar {
	if in == nil {
		return nil
	}
	out := new(AstroStar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AstroStatus) DeepCopyInto(out *AstroStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AstroCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeploymentRef != nil {
		in, out := &in.DeploymentRef, &out.DeploymentRef
		*out = make([]AstroRef, len(*in))
		copy(*out, *in)
	}
	if in.ServiceRef != nil {
		in, out := &in.ServiceRef, &out.ServiceRef
		*out = make([]AstroRef, len(*in))
		copy(*out, *in)
	}
	out.AstermuleRef = in.AstermuleRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AstroStatus.
func (in *AstroStatus) DeepCopy() *AstroStatus {
	if in == nil {
		return nil
	}
	out := new(AstroStatus)
	in.DeepCopyInto(out)
	return out
}
