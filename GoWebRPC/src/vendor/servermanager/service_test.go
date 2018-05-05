package servermanager
import(
    "testing"
    "net/http"
    "net/rpc"
    "golang.org/x/net/websocket"

)

type Users struct {
    Username    string
    Psd         int
}
type Blogs struct {
    Author  string
    Content string
}
type Likes struct {
    Num int
    Id  int
}


func TestServices(t *testing.T){
    client, err := Dial("tcp", "localhost:8010")
	if err != nil {
		t.Fatal("dialing", err)
	}

    userargs := &Users{"aaa", 111}
    blogargs := &Blogs{"aaa", "123123"}
    likeargs := &Likes{1, 0}
	var reply
	err = client.Call("net.rpc.UserInfo.Signup", userargs, reply)
	if err != nil {
		t.Errorf("Signup: expected no error but got %q", err.Error())
	}
	if reply != 0 {
		t.Errorf("Signup: expected %d got %d", 0, reply)
	}

    err = client.Call("net.rpc.BlogInfo.AddBlog", blogargs, reply)
	if err != nil {
		t.Errorf("AddBlog: expected no error but got %q", err.Error())
	}
	if reply != 0 {
		t.Errorf("AddBlog: expected %d got %d", 0, reply)
	}

    err = client.Call("net.rpc.LikeInfo.LikeHandler", likeargs, reply)
	if err != nil {
		t.Errorf("LikeHandler: expected no error but got %q", err.Error())
	}
	if reply.C != 1 {
		t.Errorf("LikeHandler: expected %d got %d", 1, reply)
	}

}
