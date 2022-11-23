package proto

import (
	"testing"

	"github.com/nrc-no/notcore/internal/utils/pointers"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestMarshalForm(t *testing.T) {

	test := []struct {
		name  string
		input *Form
	}{
		{
			// Outputs
			//           "id": "test",
			//          "title": "Test Form",
			//          "description": "This is a test form",
			//          "fields": [
			//            {
			//              "id": "text-field-1",
			//              "typedConfig": {
			//                "@type": "type.googleapis.com/no.nrc.core.form.FormTextField",
			//                "label": "Text Field 1"
			//              }
			//            },
			//            {
			//              "id": "textarea-field-1",
			//              "typedConfig": {
			//                "@type": "type.googleapis.com/no.nrc.core.form.FormTextareaField",
			//                "label": "Textarea Field 1"
			//              }
			//            },
			//            {
			//              "id": "date-field-1",
			//              "typedConfig": {
			//                "@type": "type.googleapis.com/no.nrc.core.form.FormDateField",
			//                "label": "Date Field 1"
			//              }
			//            },
			//            {
			//              "id": "select-field-1",
			//              "typedConfig": {
			//                "@type": "type.googleapis.com/no.nrc.core.form.FormSelectField",
			//                "label": "Select Field 1",
			//                "options": [
			//                  {
			//                    "label": "Option 1",
			//                    "value": "option-1"
			//                  },
			//                  {
			//                    "label": "Option 2",
			//                    "value": "option-2"
			//                  }
			//                ]
			//              }
			//            }
			//          ]
			//        }
			name: "test",
			input: &Form{
				Id:          "test",
				Title:       "Test Form",
				Description: "This is a test form",
				Fields: []*FormField{
					{
						Id: "text-field-1",
						TypedConfig: mustAny(&FormTextField{
							Label: "Text Field 1",
						}),
					}, {
						Id: "textarea-field-1",
						TypedConfig: mustAny(&FormTextareaField{
							Label: "Textarea Field 1",
						}),
					}, {
						Id: "date-field-1",
						TypedConfig: mustAny(&FormDateField{
							Label: "Date Field 1",
						}),
					}, {
						Id: "select-field-1",
						TypedConfig: mustAny(&FormSelectField{
							Label: "Select Field 1",
							Options: []*FormSelectOption{
								{
									Label: "Option 1",
									Value: pointers.String("option-1"),
								}, {
									Label: "Option 2",
									Value: pointers.String("option-2"),
								},
							},
						}),
					},
				},
			},
		},
	}
	marshalOpts := protojson.MarshalOptions{
		Indent: "  ",
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			formJson, _ := marshalOpts.Marshal(tt.input)
			t.Log(string(formJson))
		})
	}
}

func TestMarshalPredicate(t *testing.T) {
	tests := []struct {
		name  string
		input *RecordPredicate
	}{
		{

			// Outputs
			//           "predicate": {
			//            "@type": "type.googleapis.com/no.nrc.core.form.AnyOfPredicate",
			//            "predicates": [
			//              {
			//                "predicate": {
			//                  "@type": "type.googleapis.com/no.nrc.core.form.EqualityPredicate",
			//                  "fieldId": "field-1",
			//                  "value": "value-1"
			//                }
			//              },
			//              {
			//                "predicate": {
			//                  "@type": "type.googleapis.com/no.nrc.core.form.EqualityPredicate",
			//                  "fieldId": "field-2",
			//                  "value": "value-1"
			//                }
			//              },
			//              {
			//                "predicate": {
			//                  "@type": "type.googleapis.com/no.nrc.core.form.AllOfPredicate",
			//                  "predicates": [
			//                    {
			//                      "predicate": {
			//                        "@type": "type.googleapis.com/no.nrc.core.form.EqualityPredicate",
			//                        "fieldId": "field-3",
			//                        "value": "value-1"
			//                      }
			//                    },
			//                    {
			//                      "predicate": {
			//                        "@type": "type.googleapis.com/no.nrc.core.form.EqualityPredicate",
			//                        "fieldId": "field-4",
			//                        "value": "value-1"
			//                      }
			//                    }
			//                  ]
			//                }
			//              }
			//            ]
			//          }
			//        }
			name: "simple",
			input: &RecordPredicate{
				Predicate: mustAny(&AnyOfPredicate{
					Predicates: []*RecordPredicate{
						{
							Predicate: mustAny(&EqualityPredicate{
								FieldId: "field-1",
								Value:   structpb.NewStringValue("value-1"),
							}),
						}, {
							Predicate: mustAny(&EqualityPredicate{
								FieldId: "field-2",
								Value:   structpb.NewStringValue("value-1"),
							}),
						}, {
							Predicate: mustAny(&AllOfPredicate{
								Predicates: []*RecordPredicate{
									{
										Predicate: mustAny(&EqualityPredicate{
											FieldId: "field-3",
											Value:   structpb.NewStringValue("value-1"),
										}),
									}, {
										Predicate: mustAny(&EqualityPredicate{
											FieldId: "field-4",
											Value:   structpb.NewStringValue("value-1"),
										}),
									},
								},
							}),
						},
					},
				}),
			},
		},
	}
	marshalOpts := protojson.MarshalOptions{
		Indent: "  ",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formJson, _ := marshalOpts.Marshal(tt.input)
			t.Log(string(formJson))
		})
	}
}

func mustAny(pb proto.Message) *anypb.Any {
	anyVal, err := anypb.New(pb)
	if err != nil {
		panic(err)
	}
	return anyVal
}
