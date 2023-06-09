// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package kubernetes

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes(in *jlexer.Lexer, out *ListResourcesByNamespaceRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "api_version":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "label_selector":
			out.LabelSelector = string(in.String())
		case "field_selector":
			out.FieldSelector = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes(out *jwriter.Writer, in ListResourcesByNamespaceRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"api_version\":"
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	{
		const prefix string = ",\"label_selector\":"
		out.RawString(prefix)
		out.String(string(in.LabelSelector))
	}
	{
		const prefix string = ",\"field_selector\":"
		out.RawString(prefix)
		out.String(string(in.FieldSelector))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListResourcesByNamespaceRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListResourcesByNamespaceRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListResourcesByNamespaceRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListResourcesByNamespaceRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes1(in *jlexer.Lexer, out *ListAllResourcesRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "api_version":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "label_selector":
			out.LabelSelector = string(in.String())
		case "field_selector":
			out.FieldSelector = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes1(out *jwriter.Writer, in ListAllResourcesRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"api_version\":"
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"label_selector\":"
		out.RawString(prefix)
		out.String(string(in.LabelSelector))
	}
	{
		const prefix string = ",\"field_selector\":"
		out.RawString(prefix)
		out.String(string(in.FieldSelector))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ListAllResourcesRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ListAllResourcesRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ListAllResourcesRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ListAllResourcesRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes1(l, v)
}
func easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes2(in *jlexer.Lexer, out *GetResourceRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "api_version":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "disable_cache":
			out.DisableCache = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes2(out *jwriter.Writer, in GetResourceRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"api_version\":"
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		out.String(string(in.Namespace))
	}
	{
		const prefix string = ",\"disable_cache\":"
		out.RawString(prefix)
		out.Bool(bool(in.DisableCache))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetResourceRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetResourceRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetResourceRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetResourceRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComKubewardenPolicySdkGoPkgCapabilitiesKubernetes2(l, v)
}
