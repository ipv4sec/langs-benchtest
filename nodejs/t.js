
function f1() {
    console.log("Hello World");
}

function f2() {
    console.log("程序执行结束!");
}

function f3() {
    console.log("程序执行结束!");
}

function f4() {
    console.log("程序执行完毕。");
}

function f5() {
    console.log("程序执行完毕");
}

function f6() {
}

function f7() {
}

function f8() {
    console.log("程序执行完毕。");
}

function f9() {
    const buf = Buffer.from('runoob', 'ascii');

// 输出 72756e6f6f62
    console.log(buf.toString('hex'));

// 输出 cnVub29i
    console.log(buf.toString('base64'));
}

function f10() {
    buf = Buffer.alloc(256);
    len = buf.write("www.runoob.com");

    console.log("写入字节数 : "+  len);
}

function f11() {
    buf = Buffer.alloc(26);
    for (var i = 0 ; i < 26 ; i++) {
        buf[i] = i + 97;
    }

    console.log( buf.toString('ascii'));       // 输出: abcdefghijklmnopqrstuvwxyz
    console.log( buf.toString('ascii',0,5));   //使用 'ascii' 编码, 并输出: abcde
    console.log( buf.toString('utf8',0,5));    // 使用 'utf8' 编码, 并输出: abcde
    console.log( buf.toString(undefined,0,5)); // 使用默认的 'utf8' 编码, 并输出: abcde
}

function f12() {
    const buf = Buffer.from([0x1, 0x2, 0x3, 0x4, 0x5]);
    const json = JSON.stringify(buf);

// 输出: {"type":"Buffer","data":[1,2,3,4,5]}
    console.log(json);

    const copy = JSON.parse(json, (key, value) => {
        return value && value.type === 'Buffer' ?
            Buffer.from(value.data) :
            value;
    });

// 输出: <Buffer 01 02 03 04 05>
    console.log(copy);
}

function f13() {
    var buffer1 = Buffer.from(('菜鸟教程'));
    var buffer2 = Buffer.from(('www.runoob.com'));
    var buffer3 = Buffer.concat([buffer1,buffer2]);
    console.log("buffer3 内容: " + buffer3.toString());
}

function f14() {
    var buffer1 = Buffer.from('ABC');
    var buffer2 = Buffer.from('ABCD');
    var result = buffer1.compare(buffer2);

    if(result < 0) {
        console.log(buffer1 + " 在 " + buffer2 + "之前");
    }else if(result == 0){
        console.log(buffer1 + " 与 " + buffer2 + "相同");
    }else {
        console.log(buffer1 + " 在 " + buffer2 + "之后");
    }
}

function f15() {
    var buf1 = Buffer.from('abcdefghijkl');
    var buf2 = Buffer.from('RUNOOB');

//将 buf2 插入到 buf1 指定位置上
    buf2.copy(buf1, 2);

    console.log(buf1.toString());
}

function f16() {
    var buffer1 = Buffer.from('runoob');
// 剪切缓冲区
    var buffer2 = buffer1.slice(0,2);
    console.log("buffer2 content: " + buffer2.toString());
}

function f17() {
    var buffer = Buffer.from('www.runoob.com');
//  缓冲区长度
    console.log("buffer length: " + buffer.length);
}

function f18() {
    console.log("程序执行完毕");
}

function f19() {
    console.log("程序执行完毕");
}

function f20() {
    console.log("程序执行完毕");
}

function say(word) {
    console.log(word);
}

function execute(someFunction, value) {
    someFunction(value);
}
function printHello(){
    console.log( "Hello, World!");
}
function f21() {
    console.log('Hello world');
    console.log('byvoid%diovyb');
    console.log('byvoid%diovyb', 1991);
}

function f22() {
    execute(say, "Hello");
    execute(function(word){ console.log(word) }, "Hello");
    console.log( __filename );
    console.log( __dirname );
}

function main() {
    for (var i= 0; i < 10000; i++) {
        f1()
        f2()
        f3()
        f4()
        f5()
        f6()
        f7()
        f8()
        f9()
        f10()

        f11()
        f12()
        f13()
        f14()
        f15()
        f16()
        f17()
        f18()
        f19()
        f20()

        f21()
        f22()
    }
}

function getNanoSecTime() {
    var hrTime = process.hrtime();
    return hrTime[0] * 1000000000 + hrTime[1];
}

var startAt = getNanoSecTime()
main()
console.log(getNanoSecTime()-startAt)