# Go Introspection

    package main

    import (
        "fmt"

        "github.com/cgeninet/introspect"
    )

    var test struct {
        Field1  string
        Field2  int
        private string
        Field3  bool
    }

    type test2 struct {
        Field1  string
        Field2  int
        private string
        Field3  bool
    }

    func main() {
        // Test 1
        test.Field1 = "hello"
        test.Field2 = 1337
    
        i := introspect.NewStruct(test)
        fmt.Println(i.Keys())

        // Test 2
        s := new(test2)
        s.Field1 = "hello"
        s.Field2 = 1337

        i2 := introspect.NewStruct(s)
        fmt.Println(i2.Keys())
    }
