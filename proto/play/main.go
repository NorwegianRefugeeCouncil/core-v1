package main

import (
	"os"
	"path"

	nrcproto "github.com/nrc-no/notcore/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {

	form := &nrcproto.Form{
		Id:          "test",
		Title:       "Test Form",
		Description: "This is a test form",
		Fields: []*nrcproto.FormField{
			{
				Id: "text-field-1",
				TypedConfig: mustAny(&nrcproto.FormTextField{
					Label: "Text Field 1",
				}),
			}, {
				Id: "textarea-field-1",
				TypedConfig: mustAny(&nrcproto.FormTextareaField{
					Label: "Textarea Field 1",
				}),
			}, {
				Id: "date-field-1",
				TypedConfig: mustAny(&nrcproto.FormDateField{
					Label: "Date Field 1",
				}),
			},
		},
	}

	marshalOpts := protojson.MarshalOptions{
		Indent: "  ",
	}

	formJson, _ := marshalOpts.Marshal(form)
	formJsonFile, _ := os.Create(path.Join("proto/play", "form.json"))
	_, _ = formJsonFile.Write(formJson)

}

func mustAny(pb proto.Message) *anypb.Any {
	any, err := anypb.New(pb)
	if err != nil {
		panic(err)
	}
	return any
}
