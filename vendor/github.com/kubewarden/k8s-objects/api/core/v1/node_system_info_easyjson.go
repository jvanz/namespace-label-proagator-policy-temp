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

func easyjson56ebc1a3DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *NodeSystemInfo) {
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
		case "architecture":
			if in.IsNull() {
				in.Skip()
				out.Architecture = nil
			} else {
				if out.Architecture == nil {
					out.Architecture = new(string)
				}
				*out.Architecture = string(in.String())
			}
		case "bootID":
			if in.IsNull() {
				in.Skip()
				out.BootID = nil
			} else {
				if out.BootID == nil {
					out.BootID = new(string)
				}
				*out.BootID = string(in.String())
			}
		case "containerRuntimeVersion":
			if in.IsNull() {
				in.Skip()
				out.ContainerRuntimeVersion = nil
			} else {
				if out.ContainerRuntimeVersion == nil {
					out.ContainerRuntimeVersion = new(string)
				}
				*out.ContainerRuntimeVersion = string(in.String())
			}
		case "kernelVersion":
			if in.IsNull() {
				in.Skip()
				out.KernelVersion = nil
			} else {
				if out.KernelVersion == nil {
					out.KernelVersion = new(string)
				}
				*out.KernelVersion = string(in.String())
			}
		case "kubeProxyVersion":
			if in.IsNull() {
				in.Skip()
				out.KubeProxyVersion = nil
			} else {
				if out.KubeProxyVersion == nil {
					out.KubeProxyVersion = new(string)
				}
				*out.KubeProxyVersion = string(in.String())
			}
		case "kubeletVersion":
			if in.IsNull() {
				in.Skip()
				out.KubeletVersion = nil
			} else {
				if out.KubeletVersion == nil {
					out.KubeletVersion = new(string)
				}
				*out.KubeletVersion = string(in.String())
			}
		case "machineID":
			if in.IsNull() {
				in.Skip()
				out.MachineID = nil
			} else {
				if out.MachineID == nil {
					out.MachineID = new(string)
				}
				*out.MachineID = string(in.String())
			}
		case "operatingSystem":
			if in.IsNull() {
				in.Skip()
				out.OperatingSystem = nil
			} else {
				if out.OperatingSystem == nil {
					out.OperatingSystem = new(string)
				}
				*out.OperatingSystem = string(in.String())
			}
		case "osImage":
			if in.IsNull() {
				in.Skip()
				out.OSImage = nil
			} else {
				if out.OSImage == nil {
					out.OSImage = new(string)
				}
				*out.OSImage = string(in.String())
			}
		case "systemUUID":
			if in.IsNull() {
				in.Skip()
				out.SystemUUID = nil
			} else {
				if out.SystemUUID == nil {
					out.SystemUUID = new(string)
				}
				*out.SystemUUID = string(in.String())
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
func easyjson56ebc1a3EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in NodeSystemInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"architecture\":"
		out.RawString(prefix[1:])
		if in.Architecture == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Architecture))
		}
	}
	{
		const prefix string = ",\"bootID\":"
		out.RawString(prefix)
		if in.BootID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.BootID))
		}
	}
	{
		const prefix string = ",\"containerRuntimeVersion\":"
		out.RawString(prefix)
		if in.ContainerRuntimeVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ContainerRuntimeVersion))
		}
	}
	{
		const prefix string = ",\"kernelVersion\":"
		out.RawString(prefix)
		if in.KernelVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.KernelVersion))
		}
	}
	{
		const prefix string = ",\"kubeProxyVersion\":"
		out.RawString(prefix)
		if in.KubeProxyVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.KubeProxyVersion))
		}
	}
	{
		const prefix string = ",\"kubeletVersion\":"
		out.RawString(prefix)
		if in.KubeletVersion == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.KubeletVersion))
		}
	}
	{
		const prefix string = ",\"machineID\":"
		out.RawString(prefix)
		if in.MachineID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.MachineID))
		}
	}
	{
		const prefix string = ",\"operatingSystem\":"
		out.RawString(prefix)
		if in.OperatingSystem == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.OperatingSystem))
		}
	}
	{
		const prefix string = ",\"osImage\":"
		out.RawString(prefix)
		if in.OSImage == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.OSImage))
		}
	}
	{
		const prefix string = ",\"systemUUID\":"
		out.RawString(prefix)
		if in.SystemUUID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SystemUUID))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NodeSystemInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56ebc1a3EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeSystemInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56ebc1a3EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeSystemInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56ebc1a3DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeSystemInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56ebc1a3DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}