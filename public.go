package riffle

import (
    "fmt"
)

var sess *session

func someCall(a int, b int) int {
    ret := a + b

    fmt.Println("Call received with args and return:", a, b, ret)
    return ret
}

func somePub(a int, b int) {
    fmt.Println("Pub received with args:", a, b)
    sess.Leave()
}

func Tester(name string) string {
    fmt.Println("I can print!")
    
    //return fmt.Sprintf("Hello, %s!", name)

    s, err := Start("ws://ec2-52-26-83-61.us-west-2.compute.amazonaws.com:8000/ws", "xs.damouse.go")
    sess = s

    if err != nil {
        fmt.Println(err)
        return "NO"
    }

    s.Register("xs.damouse.go/hello", someCall, nil)
    fmt.Println("Registered")

    s.Subscribe("xs.damouse.go/sub", somePub)
    fmt.Println("Subscribed")

    // Block and recieve
    go s.Receive()

    fmt.Println("Done!")

    return "YES"
}
