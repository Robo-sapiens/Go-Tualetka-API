// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson9e1087fdDecodeMainInternalModels(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.ID = int(in.Int())
		case "nickname":
			out.Nickname = string(in.String())
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "phone":
			out.Phone = string(in.String())
		case "status":
			out.Status = string(in.String())
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
func easyjson9e1087fdEncodeMainInternalModels(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	if in.Name != nil {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(*in.Name))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeMainInternalModels(l, v)
}