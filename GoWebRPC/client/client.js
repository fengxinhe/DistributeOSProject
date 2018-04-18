$(function() {
    var id = -1;
    var msg = $('#msg-list');
    var rpc = jsonrpc.NewClient("ws://"+window.location.host+"/jsonrpc");
    var push;

    rpc.ws.onopen = function(){

    };

    rpc.ws.addEventListener("message", function(event) {

        var data = event.data;
        //console.log(data)
    });

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
            method: "UserInfo.Signup",
            params: new Array(largs),
            success: function(result){
                id = result;
                console.log(id)
                if(id!=-1){
                    msg.prepend("<li>id: " + result + "</li>");
                    $('#userform').hide();
                    push = jsonrpc.NewServer("ws://"+window.location.host+"/push");
                    msghand=jsonrpc.NewServer("ws://"+window.location.host+"/notify");
                    push.Connect();
                }
            },
            error: function(error){
                msg.prepend("<li>user singn up error: " + error + "</li>");
            }
        });
    });

    $('#send').click(function(){
        var args = {};
        args.Content = $('#blog').val();
        //console.log(args.Content)
        args.Heat=0
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

    $('#multiply').click(function(){
        var args = {};
        args.A = parseInt($('#A').val());
        args.B = parseInt($('#B').val());
        $('#A').val(args.A);
        $('#B').val(args.B);
        args.id = id;
        rpc.Call({
            method: "Arith.Multiply",
            params: new Array(args),
            success: function(result){
                $('#multiply-result').val(result);
            },
            error: function(error){
                msg.prepend("<li>rpc error: " + error + "</li>");
                $('#multiply-result').val("0.0");
            }
        });
    });

});
