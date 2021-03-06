// Code generated by zanzibar
// @generated
// Checksum : pmTCG1Dq7lD4AXp+UflKTQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package multi

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

func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello(in *jlexer.Lexer, out *ServiceCFront_Hello_Result) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
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
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello(out *jwriter.Writer, in ServiceCFront_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceCFront_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceCFront_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceCFront_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceCFront_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello1(in *jlexer.Lexer, out *ServiceCFront_Hello_Args) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello1(out *jwriter.Writer, in ServiceCFront_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceCFront_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceCFront_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceCFront_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceCFront_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceCFrontHello1(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello(in *jlexer.Lexer, out *ServiceBFront_Hello_Result) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
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
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello(out *jwriter.Writer, in ServiceBFront_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBFront_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBFront_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBFront_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBFront_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello1(in *jlexer.Lexer, out *ServiceBFront_Hello_Args) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello1(out *jwriter.Writer, in ServiceBFront_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBFront_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBFront_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBFront_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBFront_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceBFrontHello1(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello(in *jlexer.Lexer, out *ServiceAFront_Hello_Result) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
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
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello(out *jwriter.Writer, in ServiceAFront_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceAFront_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceAFront_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceAFront_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceAFront_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello1(in *jlexer.Lexer, out *ServiceAFront_Hello_Args) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello1(out *jwriter.Writer, in ServiceAFront_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceAFront_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceAFront_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceAFront_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceAFront_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsMultiMultiServiceAFrontHello1(l, v)
}
