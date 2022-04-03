package introspect

import (
	"fmt"
	"log"
	"reflect"
)

// Struct : object struct to do introspection
type Struct struct {
	Separator string
	obj       interface{}
	data      map[string]*interface{}
	keys      []string
}

// NewStruct : create object Struct to do introspection
func NewStruct(obj interface{}) *Struct {
	me := new(Struct)
	me.Separator = "/"

	me.obj = obj
	me.data = make(map[string]*interface{})
	me.walk(me.obj, "")

	keys := reflect.ValueOf(me.data).MapKeys()
	me.keys = make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		me.keys[i] = keys[i].String()
	}

	return me
}

// Get an array of paths
func (me *Struct) Keys() []string {
	return me.keys
}

// Get type of value for path
func (me *Struct) TypeOf(path string) string {
	if me.data[path] == nil {
		return "nil"
	}
	return reflect.TypeOf(*me.data[path]).String()
}

// Get value from path
func (me *Struct) Value(path string) interface{} {
	return *me.data[path]
}

// Get ptr for path
func (me *Struct) Get(path string) *interface{} {
	return me.data[path]
}

// Set value for path
func (me *Struct) Set(path string, value interface{}) {
	*me.data[path] = value
}

// From : https://stackoverflow.com/questions/24348184/get-pointer-to-value-using-reflection
func (me *Struct) walk(obj interface{}, path string) {
	var prefix string
	if t := reflect.TypeOf(obj); t.Kind() == reflect.Ptr {
		prefix = t.Elem().Name()
	} else {
		prefix = t.Name()
	}
	if prefix != "" {
		path = me.Separator + prefix + me.Separator
	} else {
		path += prefix + me.Separator
	}

	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			item := val.Field(i)
			itemType := val.Type().Field(i)
			if item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			fullpath := path + itemType.Name
			if item.Kind() == reflect.Struct && item.CanInterface() {
				me.walk(item.Interface(), fullpath)
			} else if (item.CanAddr() || item.Kind() != reflect.Invalid) && item.CanInterface() {
				itemInterface := item.Interface()
				switch reflect.TypeOf(itemInterface).Kind() {
				case reflect.Map:
					me.walk(itemInterface, fullpath)
				case reflect.Slice:
					me.walk(itemInterface, fullpath)
				default:
					log.Println("struct item>", fullpath, "=", itemInterface)
					me.data[fullpath] = &itemInterface
				}
			} else {
				log.Println("priv   item>", fullpath)
			}
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			item := val.MapIndex(key)
			itemInterface := item.Interface()
			if itemInterface == nil {
				continue
			}
			fullpath := path + fmt.Sprint(key.Interface())
			switch reflect.TypeOf(item.Interface()).Kind() {
			case reflect.Map:
				me.walk(itemInterface, fullpath)
			case reflect.Slice:
				me.walk(itemInterface.([]interface{}), fullpath)
			default:
				log.Println("map    item>", fullpath, "=", item)
				me.data[fullpath] = &itemInterface
			}
		}
	case reflect.Slice:
		items := reflect.ValueOf(obj)
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			itemInterface := item.Interface()
			itemType := reflect.ValueOf(itemInterface).Kind()
			fullpath := path + fmt.Sprint(i)
			switch itemType {
			case reflect.Map:
				me.walk(itemInterface, fullpath)
			case reflect.Slice:
				me.walk(itemInterface, fullpath)
			default:
				log.Println("slice  item>", fullpath, "=", itemInterface)
				me.data[fullpath] = &itemInterface
			}
		}
	}
}
