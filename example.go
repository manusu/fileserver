package main  
  
import (  
    "encoding/base64"  
    "fmt"  
)  
  
const (  
    //base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"  
    base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)  
  
var coder = base64.NewEncoding(base64Table)  
  
func base64Encode(src []byte) []byte {  
    return []byte(coder.EncodeToString(src))  
}  
  
func base64Decode(src []byte) ([]byte, error) {  
    return coder.DecodeString(string(src))  
}  
  
func test() {  
    // encode     
    hello := "hello world"  
    debyte := base64Encode([]byte(hello))  
  
    // decode     
    enbyte, err := base64Decode(debyte)  
    if err != nil {  
        fmt.Println(err.Error())  
    }  
  
    if hello != string(enbyte) {  
        fmt.Println("hello is not equal to enbyte")  
    }  
  
    fmt.Println(string(enbyte))  
}  
