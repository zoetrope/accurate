//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v2

import (
	unsafe "unsafe"

	v2alpha1 "github.com/cybozu-go/accurate/api/accurate/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*SubNamespace)(nil), (*v2alpha1.SubNamespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_SubNamespace_To_v2alpha1_SubNamespace(a.(*SubNamespace), b.(*v2alpha1.SubNamespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.SubNamespace)(nil), (*SubNamespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_SubNamespace_To_v2_SubNamespace(a.(*v2alpha1.SubNamespace), b.(*SubNamespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SubNamespaceList)(nil), (*v2alpha1.SubNamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_SubNamespaceList_To_v2alpha1_SubNamespaceList(a.(*SubNamespaceList), b.(*v2alpha1.SubNamespaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.SubNamespaceList)(nil), (*SubNamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_SubNamespaceList_To_v2_SubNamespaceList(a.(*v2alpha1.SubNamespaceList), b.(*SubNamespaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SubNamespaceSpec)(nil), (*v2alpha1.SubNamespaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_SubNamespaceSpec_To_v2alpha1_SubNamespaceSpec(a.(*SubNamespaceSpec), b.(*v2alpha1.SubNamespaceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.SubNamespaceSpec)(nil), (*SubNamespaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_SubNamespaceSpec_To_v2_SubNamespaceSpec(a.(*v2alpha1.SubNamespaceSpec), b.(*SubNamespaceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SubNamespaceStatus)(nil), (*v2alpha1.SubNamespaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2_SubNamespaceStatus_To_v2alpha1_SubNamespaceStatus(a.(*SubNamespaceStatus), b.(*v2alpha1.SubNamespaceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.SubNamespaceStatus)(nil), (*SubNamespaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_SubNamespaceStatus_To_v2_SubNamespaceStatus(a.(*v2alpha1.SubNamespaceStatus), b.(*SubNamespaceStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2_SubNamespace_To_v2alpha1_SubNamespace(in *SubNamespace, out *v2alpha1.SubNamespace, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2_SubNamespaceSpec_To_v2alpha1_SubNamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2_SubNamespaceStatus_To_v2alpha1_SubNamespaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2_SubNamespace_To_v2alpha1_SubNamespace is an autogenerated conversion function.
func Convert_v2_SubNamespace_To_v2alpha1_SubNamespace(in *SubNamespace, out *v2alpha1.SubNamespace, s conversion.Scope) error {
	return autoConvert_v2_SubNamespace_To_v2alpha1_SubNamespace(in, out, s)
}

func autoConvert_v2alpha1_SubNamespace_To_v2_SubNamespace(in *v2alpha1.SubNamespace, out *SubNamespace, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2alpha1_SubNamespaceSpec_To_v2_SubNamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_SubNamespaceStatus_To_v2_SubNamespaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2alpha1_SubNamespace_To_v2_SubNamespace is an autogenerated conversion function.
func Convert_v2alpha1_SubNamespace_To_v2_SubNamespace(in *v2alpha1.SubNamespace, out *SubNamespace, s conversion.Scope) error {
	return autoConvert_v2alpha1_SubNamespace_To_v2_SubNamespace(in, out, s)
}

func autoConvert_v2_SubNamespaceList_To_v2alpha1_SubNamespaceList(in *SubNamespaceList, out *v2alpha1.SubNamespaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v2alpha1.SubNamespace)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2_SubNamespaceList_To_v2alpha1_SubNamespaceList is an autogenerated conversion function.
func Convert_v2_SubNamespaceList_To_v2alpha1_SubNamespaceList(in *SubNamespaceList, out *v2alpha1.SubNamespaceList, s conversion.Scope) error {
	return autoConvert_v2_SubNamespaceList_To_v2alpha1_SubNamespaceList(in, out, s)
}

func autoConvert_v2alpha1_SubNamespaceList_To_v2_SubNamespaceList(in *v2alpha1.SubNamespaceList, out *SubNamespaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]SubNamespace)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2alpha1_SubNamespaceList_To_v2_SubNamespaceList is an autogenerated conversion function.
func Convert_v2alpha1_SubNamespaceList_To_v2_SubNamespaceList(in *v2alpha1.SubNamespaceList, out *SubNamespaceList, s conversion.Scope) error {
	return autoConvert_v2alpha1_SubNamespaceList_To_v2_SubNamespaceList(in, out, s)
}

func autoConvert_v2_SubNamespaceSpec_To_v2alpha1_SubNamespaceSpec(in *SubNamespaceSpec, out *v2alpha1.SubNamespaceSpec, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_v2_SubNamespaceSpec_To_v2alpha1_SubNamespaceSpec is an autogenerated conversion function.
func Convert_v2_SubNamespaceSpec_To_v2alpha1_SubNamespaceSpec(in *SubNamespaceSpec, out *v2alpha1.SubNamespaceSpec, s conversion.Scope) error {
	return autoConvert_v2_SubNamespaceSpec_To_v2alpha1_SubNamespaceSpec(in, out, s)
}

func autoConvert_v2alpha1_SubNamespaceSpec_To_v2_SubNamespaceSpec(in *v2alpha1.SubNamespaceSpec, out *SubNamespaceSpec, s conversion.Scope) error {
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_v2alpha1_SubNamespaceSpec_To_v2_SubNamespaceSpec is an autogenerated conversion function.
func Convert_v2alpha1_SubNamespaceSpec_To_v2_SubNamespaceSpec(in *v2alpha1.SubNamespaceSpec, out *SubNamespaceSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_SubNamespaceSpec_To_v2_SubNamespaceSpec(in, out, s)
}

func autoConvert_v2_SubNamespaceStatus_To_v2alpha1_SubNamespaceStatus(in *SubNamespaceStatus, out *v2alpha1.SubNamespaceStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v2_SubNamespaceStatus_To_v2alpha1_SubNamespaceStatus is an autogenerated conversion function.
func Convert_v2_SubNamespaceStatus_To_v2alpha1_SubNamespaceStatus(in *SubNamespaceStatus, out *v2alpha1.SubNamespaceStatus, s conversion.Scope) error {
	return autoConvert_v2_SubNamespaceStatus_To_v2alpha1_SubNamespaceStatus(in, out, s)
}

func autoConvert_v2alpha1_SubNamespaceStatus_To_v2_SubNamespaceStatus(in *v2alpha1.SubNamespaceStatus, out *SubNamespaceStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v2alpha1_SubNamespaceStatus_To_v2_SubNamespaceStatus is an autogenerated conversion function.
func Convert_v2alpha1_SubNamespaceStatus_To_v2_SubNamespaceStatus(in *v2alpha1.SubNamespaceStatus, out *SubNamespaceStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_SubNamespaceStatus_To_v2_SubNamespaceStatus(in, out, s)
}
