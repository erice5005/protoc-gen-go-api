package main

import "testing"

func makeTestMessage() GeneratableMessage {
	return GeneratableMessage{
		StructName: "TestObj",
		Fields: []MessageField{
			{
				Name:     "First",
				JSONName: "first",
				Kind:     "string",
			},
			{
				Name:     "Second",
				JSONName: "second",
				Kind:     "int32",
			},
		},
	}
}
func Test_GeneratableMessage(t *testing.T) {
	testMg := makeTestMessage()
	t.Logf("Struct: %v\n", generateStruct(testMg))
	t.Logf("Getters: %v\n", makeGetters(testMg))
	t.Logf("Setters: %v\n", makeSetters(testMg))
}
