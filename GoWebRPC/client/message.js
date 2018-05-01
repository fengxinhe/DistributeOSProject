//document.write("<script type='text/javascript' src='client.js'></script>");

function handleMessage(data){
    // console.log("data:")
    // console.log(data)
    method=getMethod(data);
    msg=getMsg(data);
    if(method=="addblog"){
        addBlog(msg);
    }else if(method=="modifylike"){

        console.log(msg);
        var btnid=msg[1];
        var vote=msg[2];
        var btn=document.getElementById(btnid);
        $(btn).text(vote);
    }else if (method == "register") {
        addMember(msg[1])
    }

}
function addBlog(msg){
    blogcontent=''
    +'<div class="blog">Author: '//+ //'<img src="' + getGravatar(msg[2]) + '">'
    +msg[2]+''
    +'<br>Blog:   '+msg.slice(3).join(" ")+''
    +'<button style="background-color:MediumSeaGreen; color:white;" class="likebtn" id="'+msg[1]+'" value="like">0</button>'
    +'</div><br>';
    //$('#blog-list').prepend('Star <button style="background-color:Orange; color:MediumSeaGreen;" type="button" id="'+msg[1]+'" value="like">0</button><br>');
    $('#blog-list').prepend(blogcontent);
}
function addMember(member){
    var m='<div class="member"> <strong>'+member+'</strong>'+
    '<button class="button follow" id="'+member+'" value="follow">follow</button></div>';
    $('#user-list').append(m);
}
function getMethod(str){
    var method = str.split(" ",1);
    return method;
}
function getMsg(str){
    var msg = str.split(" ");
    console.log(msg[0])
    return msg;
}

function getGravatar(user) {
    return 'http://www.gravatar.com/avatar/' + user;
}
