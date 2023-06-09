// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson2ffb899eDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *FlexPersistentVolumeSource) {
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
		case "driver":
			if in.IsNull() {
				in.Skip()
				out.Driver = nil
			} else {
				if out.Driver == nil {
					out.Driver = new(string)
				}
				*out.Driver = string(in.String())
			}
		case "fsType":
			out.FSType = string(in.String())
		case "options":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Options = make(map[string]string)
				} else {
					out.Options = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Options)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "readOnly":
			out.ReadOnly = bool(in.Bool())
		case "secretRef":
			if in.IsNull() {
				in.Skip()
				out.SecretRef = nil
			} else {
				if out.SecretRef == nil {
					out.SecretRef = new(SecretReference)
				}
				easyjson2ffb899eDecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.SecretRef)
			}
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
func easyjson2ffb899eEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in FlexPersistentVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"driver\":"
		out.RawString(prefix[1:])
		if in.Driver == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Driver))
		}
	}
	if in.FSType != "" {
		const prefix string = ",\"fsType\":"
		out.RawString(prefix)
		out.String(string(in.FSType))
	}
	if len(in.Options) != 0 {
		const prefix string = ",\"options\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Options {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	if in.ReadOnly {
		const prefix string = ",\"readOnly\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnly))
	}
	if in.SecretRef != nil {
		const prefix string = ",\"secretRef\":"
		out.RawString(prefix)
		easyjson2ffb899eEncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.SecretRef)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FlexPersistentVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ffb899eEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FlexPersistentVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ffb899eEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FlexPersistentVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2ffb899eDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FlexPersistentVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ffb899eDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson2ffb899eDecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *SecretReference) {
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
		case "name":
			out.Name = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
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
func easyjson2ffb899eEncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in SecretReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	out.RawByte('}')
}
