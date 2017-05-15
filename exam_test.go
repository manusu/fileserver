package main
import (
    "testing"
)

func Test_base64(t *testing.T){
    str := "test"
    if string(base64Decode(base64Encode([]byte(str))))!=str{
        t.Error("test failed")
    }
}


func Test_test(t *testing.T){
    test()
    t.Log("test successed")
}
