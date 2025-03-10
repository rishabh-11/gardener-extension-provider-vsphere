//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/gardener-extension-provider-vsphere/pkg/apis/config"
	apisconfig "github.com/gardener/gardener/extensions/pkg/apis/config"
	apisconfigv1alpha1 "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ControllerConfiguration)(nil), (*config.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(a.(*ControllerConfiguration), b.(*config.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfiguration)(nil), (*ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*config.ControllerConfiguration), b.(*ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCD)(nil), (*config.ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCD_To_config_ETCD(a.(*ETCD), b.(*config.ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCD)(nil), (*ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCD_To_v1alpha1_ETCD(a.(*config.ETCD), b.(*ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDStorage)(nil), (*config.ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(a.(*ETCDStorage), b.(*config.ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDStorage)(nil), (*ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(a.(*config.ETCDStorage), b.(*ETCDStorage), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	out.GardenId = in.GardenId
	out.ClientConnection = (*componentbaseconfig.ClientConnectionConfiguration)(unsafe.Pointer(in.ClientConnection))
	if err := Convert_v1alpha1_ETCD_To_config_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*apisconfig.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	out.GardenId = in.GardenId
	out.ClientConnection = (*configv1alpha1.ClientConnectionConfiguration)(unsafe.Pointer(in.ClientConnection))
	if err := Convert_config_ETCD_To_v1alpha1_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*apisconfigv1alpha1.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	return nil
}

// Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	if err := Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ETCD_To_config_ETCD is an autogenerated conversion function.
func Convert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCD_To_config_ETCD(in, out, s)
}

func autoConvert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	if err := Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ETCD_To_v1alpha1_ETCD is an autogenerated conversion function.
func Convert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	return autoConvert_config_ETCD_To_v1alpha1_ETCD(in, out, s)
}

func autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage is an autogenerated conversion function.
func Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in, out, s)
}

func autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage is an autogenerated conversion function.
func Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	return autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in, out, s)
}
