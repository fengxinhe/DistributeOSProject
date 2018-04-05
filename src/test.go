// package main
//
// import (
//     "encoding/json"
//     "fmt"
//     "log"
//     "os"
// )
//
// type User struct {
//     Name string
//     Age [2]int
// }
//
//
// func testMarshal() []byte {
//     user := User{
//         Name: "Tab",
//         Age: [2]int{11,22},
//     }
//     data, err := json.Marshal(user)
//     if err != nil {
//         log.Fatal(err)
//     }
//     return data
// }
//
// func testUnmarshal(data []byte) {
//     var user User
//     err := json.Unmarshal(data, &user)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(user)
// }
//
// func testRead() []byte {
//     fp, err := os.OpenFile("./data.json", os.O_RDONLY, 0755)
//     defer fp.Close()
//     if err != nil {
//         log.Fatal(err)
//     }
//     data := make([]byte, 100)
//     n, err := fp.Read(data)
//     if err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(string(data[:n]))
//     return data[:n]
// }
//
// func testWrite(data []byte) {
//     fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer fp.Close()
//     _, err = fp.Write(data)
//     if err != nil {
//         log.Fatal(err)
//     }
// }
//
// func main() {
//     var data []byte
//     var data2 []byte
//     data = testMarshal()
//     data2=testMarshal()
//     fmt.Println(string(data))
//     testWrite(data)
//     testWrite(data2)
//     data = testRead()
//     testUnmarshal(data)
// }
