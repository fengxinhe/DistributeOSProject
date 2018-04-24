$(function() {
    var id = -1;
    var msg = $('#msg-list');
    var rpc = jsonrpc.NewClient("ws://"+window.location.host+"/jsonrpc");
    var push;

    rpc.ws.onopen = function(){

    };
    // rpc.ws.onmessage=function({
    //
    //     var data = event.data;
    //     console.log(data)
    // });

    $('#register').click(function(){
        var largs = {};
        largs.Username = $('#username').val();
        largs.Psd = parseInt($('#psd').val());

        rpc.Call({
            method: "UserInfo.Register",
            params: new Array(largs),
            success: function(result){
                id = result;
                console.log(id)
                if(id!=-1){
                    $('#register').hide();

                }
            },
            error: function(error){
                msg.prepend("<li>user singn up error: " + error + "</li>");
            }
        });
    });
    $('#login').click(function(){
        var largs = {};
        largs.Username = $('#username').val();
        largs.Psd = parseInt($('#psd').val());

        rpc.Call({
            method: "UserInfo.Signin",
            params: new Array(largs),
            success: function(result){
                id = result;
                if(id!=-1){
                    //msg.prepend("<li>id: " + result + "</li>");

                    $('#userform').hide();
                    $("#welcome").text("Welcome, "+largs.Username+"!");
                    push = jsonrpc.NewServer("ws://"+window.location.host+"/push");
                    //msghandler = jsonrpc.NewServer("ws://"+window.location.host+"/notify");
                    console.log(id)

                    push.Register('User.Getpsd', function(parmars){
                                    return {
                                        result: id,
                                        error: null
                                    };
                                });
                    push.Connect();
                    $('#signout').show();
                    $('#send').show();
                }
            },
            error: function(error){
                msg.prepend("<li>user singn up error: " + error + "</li>");
            }
        });
    });
    $('#signout').click(function(){
        var largs = {};
        largs.Username = $('#username').val();
        largs.Psd = parseInt($('#psd').val());
        rpc.Call({
            method: "UserInfo.Signout",
            params: new Array(largs),
            success: function(result){
                push.Close();
                $('#userform').show();
                $('#send').hide();
                $('#signout').hide();
                $("#welcome").text("Welcome, guest!");
                //msg.prepend("<li>rpc signout successfully</li>");
            },
            error: function(error){

            }
        });
        //push.Close();
    });
    $('#send').click(function(){
        var args = {};
        args.Author = $('#username').val();
        args.Content = $('#blog').val();
        //console.log(args.Content)
        //args.Heat=0
        rpc.Call({
            method: "BlogInfo.AddBlog",
            params: new Array(args),
            success: function(result){
                blogid=result;
            },
            error: function(error){
                msg.prepend("<li>rpc error: " + error + "</li>");
                $('#divide-result').val("0 ......0");
            }
        });
    });

    document.getElementById('blog-list').addEventListener('click', function (e) {
            //console.log(e.target.nodeName);
          if (e.target.nodeName == "BUTTON") {
              var args = {};
              args.Id=parseInt(e.target.id);

              console.log(e.target.value);
              var like=e.target.value;
              if(like=="like"){
                  args.Num=1;
                 e.target.value="dislike";
                 e.target.style.backgroundColor='Tomato';
                 //$(e.target).text("dislike")
                // $(e.target).children().text("1")
              }else {
                  args.Num=-1;
                 // $(e.target).text("like")
                  //$(e.target).children().text("0")
                  e.target.style.backgroundColor='MediumSeaGreen';
                  e.target.value="like";
              }
              rpc.Call({
                  method: "LikeInfo.LikeHandler",
                  params: new Array(args),
                  success: function(result){
                      blogid=result;
                  },
                  error: function(error){
                      msg.prepend("<li>like error: " + error + "</li>");
                      $('#divide-result').val("0 ......0");
                  }
              });

          }
    });


});
