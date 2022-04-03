package introspect

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

type otherStruct struct {
	OtherString string
	OtherInt    int
	Interface   map[interface{}]interface{}
}

type otherStructPtr struct {
	OtherStringPtr string
}

type testStruct struct {
	FieldString   string
	FieldInt      int
	privateString string
	Other         otherStruct
	OtherPtr      *otherStructPtr
	SliceString   []string
	SliceInt      []int
}

func TestStruct(t *testing.T) {
	test := &testStruct{
		FieldString:   "Hello world !",
		FieldInt:      1337,
		privateString: "private str",
		SliceString:   []string{"LE", "ET"},
		SliceInt:      []int{13, 37},
	}

	yamlFile, err := ioutil.ReadFile("yaml_file_test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatal(err)
	}

	testOtherPtr := &otherStructPtr{OtherStringPtr: "other struct ptr"}
	test.OtherPtr = testOtherPtr
	test.Other.Interface = data
	test.Other.OtherString = "test"

	is := NewStruct(test)
	k := is.Keys()
	n := len(k)

	log.Println("Keys :")
	for _, i := range k {
		log.Println(i, "=", is.Value(i))
	}

	if n != 19 {
		t.Errorf("Wrong number of keys got %d", n)
	} else {
		log.Println("keys", n)
	}

	if is.TypeOf("/NIL") != "nil" {
		t.Errorf("TypeOf should be nil for unknown path, but got %s", is.TypeOf("NIL"))
	}

	if is.TypeOf("/otherStruct/Interface/a") != "int" {
		t.Errorf("Value should be 'int', but got %s", is.TypeOf("/otherStruct/Interface/a"))
	}

	if is.Value("/otherStruct/Interface/a") != 1 {
		t.Errorf("Value should be '1', but got %s", is.Value("/otherStruct/Interface/a"))
	}

	if reflect.TypeOf(is.Get("/otherStruct/Interface/a")).Kind() != reflect.Ptr {
		t.Errorf("Value should be 'reflect.Ptr', but got %s", reflect.TypeOf(is.Value("/otherStruct/Interface/a")).Kind())
	}

	is.Set("/otherStruct/Interface/a", 2)

	if is.Value("/otherStruct/Interface/a") != 2 {
		t.Errorf("Value should be '2', but got %s", is.Value("/otherStruct/Interface/a"))
	}

	if is.TypeOf("/otherStruct/Interface/vars/hello") != "string" {
		t.Errorf("Value should be 'int', but got %s", is.TypeOf("/otherStruct/Interface/vars/hello"))
	}

	if is.Value("/otherStruct/Interface/vars/hello") != "world" {
		t.Errorf("Value should be '1', but got %s", is.Value("/otherStruct/Interface/vars/hello"))
	}
}
