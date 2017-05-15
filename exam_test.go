package main
import (
    "testing"
)

func Test_base64(t *testing.T){
    str := "test"
    res,_=base64Decode(base64Encode([]byte(str)))
    if string(res)!=str{
        t.Error("test failed")
    }
}


func Test_test(t *testing.T){
    test()
    t.Log("test successed")
}
