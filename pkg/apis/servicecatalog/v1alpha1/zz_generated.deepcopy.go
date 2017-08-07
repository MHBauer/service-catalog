// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_AlphaPodPresetTemplate, InType: reflect.TypeOf(&AlphaPodPresetTemplate{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BasicAuthConfig, InType: reflect.TypeOf(&BasicAuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BearerTokenAuthConfig, InType: reflect.TypeOf(&BearerTokenAuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Binding, InType: reflect.TypeOf(&Binding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingCondition, InType: reflect.TypeOf(&BindingCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingList, InType: reflect.TypeOf(&BindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingSpec, InType: reflect.TypeOf(&BindingSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingStatus, InType: reflect.TypeOf(&BindingStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Broker, InType: reflect.TypeOf(&Broker{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerAuthInfo, InType: reflect.TypeOf(&BrokerAuthInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerCondition, InType: reflect.TypeOf(&BrokerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerList, InType: reflect.TypeOf(&BrokerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerSpec, InType: reflect.TypeOf(&BrokerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerStatus, InType: reflect.TypeOf(&BrokerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Instance, InType: reflect.TypeOf(&Instance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceCondition, InType: reflect.TypeOf(&InstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceList, InType: reflect.TypeOf(&InstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceSpec, InType: reflect.TypeOf(&InstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceStatus, InType: reflect.TypeOf(&InstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Plan, InType: reflect.TypeOf(&Plan{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PlanList, InType: reflect.TypeOf(&PlanList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceClass, InType: reflect.TypeOf(&ServiceClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceClassList, InType: reflect.TypeOf(&ServiceClassList{})},
	)
}

// DeepCopy_v1alpha1_AlphaPodPresetTemplate is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_AlphaPodPresetTemplate(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*AlphaPodPresetTemplate)
		out := out.(*AlphaPodPresetTemplate)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Selector); err != nil {
			return err
		} else {
			out.Selector = *newVal.(*v1.LabelSelector)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BasicAuthConfig is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BasicAuthConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BasicAuthConfig)
		out := out.(*BasicAuthConfig)
		*out = *in
		if in.SecretRef != nil {
			in, out := &in.SecretRef, &out.SecretRef
			*out = new(api_v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BearerTokenAuthConfig is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BearerTokenAuthConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BearerTokenAuthConfig)
		out := out.(*BearerTokenAuthConfig)
		*out = *in
		if in.SecretRef != nil {
			in, out := &in.SecretRef, &out.SecretRef
			*out = new(api_v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_Binding is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Binding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Binding)
		out := out.(*Binding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_BindingSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_BindingStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BindingCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BindingCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingCondition)
		out := out.(*BindingCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1alpha1_BindingList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingList)
		out := out.(*BindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Binding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Binding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BindingSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BindingSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingSpec)
		out := out.(*BindingSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.AlphaPodPresetTemplate != nil {
			in, out := &in.AlphaPodPresetTemplate, &out.AlphaPodPresetTemplate
			*out = new(AlphaPodPresetTemplate)
			if err := DeepCopy_v1alpha1_AlphaPodPresetTemplate(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BindingStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BindingStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingStatus)
		out := out.(*BindingStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BindingCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_BindingCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Checksum != nil {
			in, out := &in.Checksum, &out.Checksum
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_Broker is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Broker(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Broker)
		out := out.(*Broker)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_BrokerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_BrokerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BrokerAuthInfo is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BrokerAuthInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerAuthInfo)
		out := out.(*BrokerAuthInfo)
		*out = *in
		if in.Basic != nil {
			in, out := &in.Basic, &out.Basic
			*out = new(BasicAuthConfig)
			if err := DeepCopy_v1alpha1_BasicAuthConfig(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Bearer != nil {
			in, out := &in.Bearer, &out.Bearer
			*out = new(BearerTokenAuthConfig)
			if err := DeepCopy_v1alpha1_BearerTokenAuthConfig(*in, *out, c); err != nil {
				return err
			}
		}
		if in.BasicAuthSecret != nil {
			in, out := &in.BasicAuthSecret, &out.BasicAuthSecret
			*out = new(api_v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BrokerCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BrokerCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerCondition)
		out := out.(*BrokerCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1alpha1_BrokerList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BrokerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerList)
		out := out.(*BrokerList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Broker, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Broker(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BrokerSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BrokerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerSpec)
		out := out.(*BrokerSpec)
		*out = *in
		if in.AuthInfo != nil {
			in, out := &in.AuthInfo, &out.AuthInfo
			*out = new(BrokerAuthInfo)
			if err := DeepCopy_v1alpha1_BrokerAuthInfo(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_BrokerStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_BrokerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerStatus)
		out := out.(*BrokerStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BrokerCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_BrokerCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Checksum != nil {
			in, out := &in.Checksum, &out.Checksum
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_Instance is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Instance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Instance)
		out := out.(*Instance)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_InstanceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_InstanceStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_InstanceCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_InstanceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceCondition)
		out := out.(*InstanceCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1alpha1_InstanceList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_InstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceList)
		out := out.(*InstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Instance, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Instance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_InstanceSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_InstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceSpec)
		out := out.(*InstanceSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_InstanceStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_InstanceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceStatus)
		out := out.(*InstanceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]InstanceCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_InstanceCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.LastOperation != nil {
			in, out := &in.LastOperation, &out.LastOperation
			*out = new(string)
			**out = **in
		}
		if in.DashboardURL != nil {
			in, out := &in.DashboardURL, &out.DashboardURL
			*out = new(string)
			**out = **in
		}
		if in.Checksum != nil {
			in, out := &in.Checksum, &out.Checksum
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1alpha1_Plan is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_Plan(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Plan)
		out := out.(*Plan)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Bindable != nil {
			in, out := &in.Bindable, &out.Bindable
			*out = new(bool)
			**out = **in
		}
		if in.ExternalMetadata != nil {
			in, out := &in.ExternalMetadata, &out.ExternalMetadata
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.AlphaInstanceCreateParameterSchema != nil {
			in, out := &in.AlphaInstanceCreateParameterSchema, &out.AlphaInstanceCreateParameterSchema
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.AlphaInstanceUpdateParameterSchema != nil {
			in, out := &in.AlphaInstanceUpdateParameterSchema, &out.AlphaInstanceUpdateParameterSchema
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.AlphaBindingCreateParameterSchema != nil {
			in, out := &in.AlphaBindingCreateParameterSchema, &out.AlphaBindingCreateParameterSchema
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_PlanList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_PlanList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PlanList)
		out := out.(*PlanList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Plan, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Plan(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceClass is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceClass(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClass)
		out := out.(*ServiceClass)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.ExternalMetadata != nil {
			in, out := &in.ExternalMetadata, &out.ExternalMetadata
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.AlphaTags != nil {
			in, out := &in.AlphaTags, &out.AlphaTags
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.AlphaRequires != nil {
			in, out := &in.AlphaRequires, &out.AlphaRequires
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceClassList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceClassList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClassList)
		out := out.(*ServiceClassList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ServiceClass, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceClass(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
