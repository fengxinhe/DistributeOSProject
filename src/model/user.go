package model

import(
    //"fmt"
    "encoding/json"
    "fmt"
    "log"
    //"os"
    "io/ioutil"

)

type User struct{
    UserName    string      `json:"user_name"`
    UserPsd     string      `json:"user_psd"`
    UserLike    int      `json:"user_like"`
}


func UserRead(fn string) (string, error) {
    path:="/home/firebug/go/gobird/src/db/account/"+fn+".json"
    data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("no such user! error username!")
        return "", err
    }
    var user User
    err = json.Unmarshal(data, &user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("ppp---"+user.UserPsd)
    return user.UserPsd,nil

    //return data[:],nil
}

func UserWrite(fn string, user User) {
    path := "/home/firebug/go/gobird/src/db/account/"+fn+".json"
    data, err := json.Marshal(user)
    if err != nil {
        log.Fatal(err)
    }
    ioutil.WriteFile(path, data,0600)
}

func UserCreate(userName string, psd string){
    var user User
    user.UserName=userName
    user.UserPsd=psd
    user.UserLike=0
    UserWrite(userName, user)
}

func UserByName(name string) (string, error){
    psd, err := UserRead(name)
    if err!=nil{
        return "",err
    }
    //var user User
    //UserUnmarshal(data,user)
    //fmt.Println(user.UserPsd)
    return psd,nil
}
