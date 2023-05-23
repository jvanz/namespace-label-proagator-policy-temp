// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
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

func easyjsonDe232084DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *Probe) {
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
		case "exec":
			if in.IsNull() {
				in.Skip()
				out.Exec = nil
			} else {
				if out.Exec == nil {
					out.Exec = new(ExecAction)
				}
				(*out.Exec).UnmarshalEasyJSON(in)
			}
		case "failureThreshold":
			out.FailureThreshold = int32(in.Int32())
		case "grpc":
			if in.IsNull() {
				in.Skip()
				out.GRPC = nil
			} else {
				if out.GRPC == nil {
					out.GRPC = new(GRPCAction)
				}
				(*out.GRPC).UnmarshalEasyJSON(in)
			}
		case "httpGet":
			if in.IsNull() {
				in.Skip()
				out.HTTPGet = nil
			} else {
				if out.HTTPGet == nil {
					out.HTTPGet = new(HTTPGetAction)
				}
				(*out.HTTPGet).UnmarshalEasyJSON(in)
			}
		case "initialDelaySeconds":
			out.InitialDelaySeconds = int32(in.Int32())
		case "periodSeconds":
			out.PeriodSeconds = int32(in.Int32())
		case "successThreshold":
			out.SuccessThreshold = int32(in.Int32())
		case "tcpSocket":
			if in.IsNull() {
				in.Skip()
				out.TCPSocket = nil
			} else {
				if out.TCPSocket == nil {
					out.TCPSocket = new(TCPSocketAction)
				}
				easyjsonDe232084DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.TCPSocket)
			}
		case "terminationGracePeriodSeconds":
			out.TerminationGracePeriodSeconds = int64(in.Int64())
		case "timeoutSeconds":
			out.TimeoutSeconds = int32(in.Int32())
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
func easyjsonDe232084EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in Probe) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Exec != nil {
		const prefix string = ",\"exec\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Exec).MarshalEasyJSON(out)
	}
	if in.FailureThreshold != 0 {
		const prefix string = ",\"failureThreshold\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FailureThreshold))
	}
	if in.GRPC != nil {
		const prefix string = ",\"grpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.GRPC).MarshalEasyJSON(out)
	}
	if in.HTTPGet != nil {
		const prefix string = ",\"httpGet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.HTTPGet).MarshalEasyJSON(out)
	}
	if in.InitialDelaySeconds != 0 {
		const prefix string = ",\"initialDelaySeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.InitialDelaySeconds))
	}
	if in.PeriodSeconds != 0 {
		const prefix string = ",\"periodSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PeriodSeconds))
	}
	if in.SuccessThreshold != 0 {
		const prefix string = ",\"successThreshold\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SuccessThreshold))
	}
	if in.TCPSocket != nil {
		const prefix string = ",\"tcpSocket\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonDe232084EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.TCPSocket)
	}
	if in.TerminationGracePeriodSeconds != 0 {
		const prefix string = ",\"terminationGracePeriodSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TerminationGracePeriodSeconds))
	}
	if in.TimeoutSeconds != 0 {
		const prefix string = ",\"timeoutSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TimeoutSeconds))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Probe) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe232084EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Probe) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe232084EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Probe) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe232084DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Probe) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe232084DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonDe232084DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *TCPSocketAction) {
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
		case "host":
			out.Host = string(in.String())
		case "port":
			if in.IsNull() {
				in.Skip()
				out.Port = nil
			} else {
				if out.Port == nil {
					out.Port = new(intstr.IntOrString)
				}
				*out.Port = intstr.IntOrString(in.String())
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
func easyjsonDe232084EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in TCPSocketAction) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Host != "" {
		const prefix string = ",\"host\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Port == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Port))
		}
	}
	out.RawByte('}')
}
