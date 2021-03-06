$(function() {
    var id = -1;
    var msg = $('#msg-list');
    var rpc = jsonrpc.NewClient("ws://"+window.location.host+"/jsonrpc");
    var push;
    var memberlist=$('#user-list');
    var membernum=0;
    var username="";

    rpc.ws.onopen = function(){
        // var largs = {};
        // largs.Username = "";
        // largs.Psd = 0;
        // rpc.Call({
        //     method:"UserInfo.GetMember",
        //     params: new Array(largs),
        //     success: function(result){
        //         //memberlist.append(result);
        //         generateMemberList(result);
        //     },
        //     error: function(error){
        //         msg.prepend("<li>Get member error: " + error + "</li>");
        //     }
        // });

    };
    function generateMemberList(list){


        for (i=0;i<list.length;i++){
            if(list[i]!="me"){
                var m='<div class="member"> <strong>'+list[i]+'</strong>'+
                '<button class="button follow" id="'+list[i]+'" value="follow">follow</button></div>';
                memberlist.append(m);
            }
        }
    }
    $('#register').click(function(){
        var largs = {};
        largs.Username = $('#username').val();
        largs.Psd = parseInt($('#psd').val());
        largs.ServiceName = "UserInfo.Register";
        largs.Method="Register"

        rpc.Call({
            method: "Command.RequestHandler",
            params: new Array(largs),
            success: function(result){
                id = result.Val;
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
        largs.ServiceName = "UserInfo.Signin";
        largs.Method="Signin"

        rpc.Call({
            method: "Command.RequestHandler",
            params: new Array(largs),
            success: function(result){
                id = result.Val;
                console.log(id)
                if(id!=-1){
                    //msg.prepend("<li>id: " + result + "</li>");
                    username=$('#username').val();
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

                    var args = {};
                    args.Username = username;
                    args.Psd = id;
                    args.ServiceName = "UserInfo.GetMember";
                    args.Method="GetMember";

                    rpc.Call({
                        method:"Command.RequestHandler",
                        params: new Array(args),
                        success: function(result){
                            //memberlist.append(result);
                            memberlist.empty();
                            generateMemberList(result.List);
                        },
                        error: function(error){
                            msg.prepend("<li>Get member error: " + error + "</li>");
                        }
                    });
                }
            },
            error: function(error){
                msg.prepend("<li>user singn up error: " + error + "</li>");
            }
        });
    });
    $('#signout').click(function(){
        var largs = {};
        largs.Username = username;
        largs.Psd = parseInt($('#psd').val());
        largs.ServiceName = "UserInfo.Signout";
        largs.Method="Signout"

        rpc.Call({
            method: "Command.RequestHandler",
            params: new Array(largs),
            success: function(result){
                push.Close();
                $('#userform').show();
                $('#send').hide();
                $('#signout').hide();
                $("#welcome").text("Welcome, guest!");
                //msg.prepend("<li>rpc signout successfully</li>");
                push.Close();
            },
            error: function(error){

            }
        });
        jsonrpc.Close();
    });
    $('#send').click(function(){
        var args = {};
        args.Author = $('#username').val();
        args.Content = $('#blog').val();
        args.ServiceName = "BlogInfo.AddBlog";
        args.Method="AddBlog"
        //console.log(args.Content)
        //args.Heat=0
        rpc.Call({
            method: "Command.RequestHandler",
            params: new Array(args),
            success: function(result){
                blogid=result.Val;
            },
            error: function(error){
                msg.prepend("<li>rpc error: " + error + "</li>");
                $('#divide-result').val("0 ......0");
            }
        });
    });
    $('#Test').click(function(){
        var args = {};
        args.ServiceName = "Arith.Multiply";
        //args.Content = $('#blog').val();
        //console.log(args.Content)
        //args.Heat=0
        rpc.Call({
            method: "Command.RequestHandler",
            params: new Array(args),
            success: function(result){
                msg.prepend("<li>relication:kkk</li>");
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
              }else {
                  args.Num=-1;
              }
              args.ServiceName = "LikeInfo.LikeHandler";
              args.Method="LikeHandler"
              rpc.Call({
                  method: "Command.RequestHandler",
                  params: new Array(args),
                  success: function(result){
                      blogid=result;
                      if(like=="like"){
                         e.target.value="dislike";
                         e.target.style.backgroundColor='Tomato';

                      }else {
                          e.target.style.backgroundColor='MediumSeaGreen';
                          e.target.value="like";
                      }
                  },
                  error: function(error){
                      msg.prepend("<li>like error: " + error + "</li>");
                      $('#divide-result').val("0 ......0");
                  }
              });

          }
    });

    document.getElementById('user-list').addEventListener('click', function (e) {
        if (e.target.nodeName == "BUTTON") {
            var args = {};
            args.UserId=id;
            args.InterestId=e.target.id;
            console.log(args.UserId);
            console.log(args.InterestId);
            var action=e.target.value;
            console.log(action)
            if(action=="follow"){
                args.Action=1;
            }else {
                args.Action=0;
            }
            args.ServiceName = "FollowInfo.FollowHandler";
            args.Method="FollowHandler"
            rpc.Call({
                method: "Command.RequestHandler",
                params: new Array(args),
                success: function(result){
                    blogid=result.Val;
                    if(action=="follow"){
                        e.target.value="unfollow";
                        $(e.target).text("unfollow");
                    }else {
                        e.target.value="follow";
                        $(e.target).text("follow");
                    }
                },
                error: function(error){
                    msg.prepend("<li>follow error: " + error + "</li>");
                    $('#divide-result').val("0 ......0");
                }
            });
        }
    });


});

