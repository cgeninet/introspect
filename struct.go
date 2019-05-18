package introspect

import (
	"log"
	"reflect"
)

// Struct : object struct to do introspection
type Struct struct {
	obj  interface{}
	data map[string]interface{}
	keys []string
}

// NewStruct : create object Struct to do introspection
func NewStruct(obj interface{}) *Struct {
	me := new(Struct)
	me.obj = obj
	me.data = make(map[string]interface{})
	me.walk(me.obj, "")

	keys := reflect.ValueOf(me.data).MapKeys()
	me.keys = make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		me.keys[i] = keys[i].String()
	}

	return me
}

// Get an array of fields
func (me *Struct) Keys() []string {
	return me.keys
}

// Get type of a field from his path
func (me *Struct) TypeOf(path string) string {
	if me.data[path] == nil {
		return "nil"
	}
	return reflect.TypeOf(me.data[path]).String()
}

// Get value of a field from his path
func (me *Struct) Value(path string) interface{} {
	return me.data[path]
}

// From : https://stackoverflow.com/questions/24348184/get-pointer-to-value-using-reflection
func (me *Struct) walk(obj interface{}, path string) {
	if path != "" {
		path += "."
	}

	if t := reflect.TypeOf(obj); t.Kind() == reflect.Ptr {
		path += t.Elem().Name() // *obj
	} else {
		path += t.Name()
	}

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}
		fullpath := path + "." + typeField.Name
		if valueField.Kind() == reflect.Struct && valueField.CanInterface() {
			me.walk(valueField.Interface(), path)
		} else if (valueField.CanAddr() || valueField.Kind() != reflect.Invalid) && valueField.CanInterface() {
			log.Print(fullpath, "= '", valueField.Interface(), "'")
			me.data[fullpath] = valueField.Interface()
		} else {
			log.Print("# ", fullpath)
		}
	}
}
