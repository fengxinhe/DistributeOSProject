//document.write("<script type='text/javascript' src='client.js'></script>");

function handleMessage(data){
    // console.log("data:")
    // console.log(data)
    method=getMethod(data);
    msg=getMsg(data);
    $('#blog-list').prepend('<button type="button" id="'+msg[1]+'" value="like">like<div>0</div></button><br>')
    $('#blog-list').prepend('<div>"'+msg[2]+msg[3]+msg[4]+'"</div><br>')

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
