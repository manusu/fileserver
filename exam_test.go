package main
import (
    "testing"
)

func Test_base64(){
    str := "test"
    if string(base64Decode(base64Encode([]byte(str))))!=str{
        t.Error("test failed")
    }
}


func Test_test(){
    test()
    t.Log("test successed")
}
