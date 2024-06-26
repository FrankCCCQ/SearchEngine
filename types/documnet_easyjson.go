package types

import (
	"encoding/json"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson18605acbDecodeGithubComCocaineCongTangsengAppSearchEngineTypes(in *jlexer.Lexer, out *Document) {
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
		case "doc_id":
			out.DocId = int64(in.Int64())
		case "title":
			out.Title = string(in.String())
		case "body":
			out.Body = string(in.String())
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
func easyjson18605acbEncodeGithubComCocaineCongTangsengAppSearchEngineTypes(out *jwriter.Writer, in Document) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"doc_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.DocId))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Document) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18605acbEncodeGithubComCocaineCongTangsengAppSearchEngineTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Document) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18605acbEncodeGithubComCocaineCongTangsengAppSearchEngineTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Document) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18605acbDecodeGithubComCocaineCongTangsengAppSearchEngineTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Document) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18605acbDecodeGithubComCocaineCongTangsengAppSearchEngineTypes(l, v)
}
