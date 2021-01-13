
console.log("Hello World");
var fs = require("fs");

var data = fs.readFileSync('input.txt');

console.log(data.toString());
console.log("程序执行结束!");

var fs = require("fs");

fs.readFile('input.txt', function (err, data) {
    if (err) return console.error(err);
    console.log(data.toString());
});

console.log("程序执行结束!");

// 引入 events 模块
var events = require('events');
// 创建 eventEmitter 对象
var eventEmitter = new events.EventEmitter();

// 创建事件处理程序
var connectHandler = function connected() {
    console.log('连接成功。');

    // 触发 data_received 事件
    eventEmitter.emit('data_received');
}

// 绑定 connection 事件处理程序
eventEmitter.on('connection', connectHandler);

// 使用匿名函数绑定 data_received 事件
eventEmitter.on('data_received', function(){
    console.log('数据接收成功。');
});

// 触发 connection 事件
eventEmitter.emit('connection');

console.log("程序执行完毕。");


var EventEmitter = require('events').EventEmitter;
var event = new EventEmitter();
event.on('some_event', function() {
    console.log('some_event 事件触发');
});
setTimeout(function() {
    event.emit('some_event');
}, 1000);


//event.js 文件
var events = require('events');
var emitter = new events.EventEmitter();
emitter.on('someEvent', function(arg1, arg2) {
    console.log('listener1', arg1, arg2);
});
emitter.on('someEvent', function(arg1, arg2) {
    console.log('listener2', arg1, arg2);
});
emitter.emit('someEvent', 'arg1 参数', 'arg2 参数');



var events = require('events');
var eventEmitter = new events.EventEmitter();

// 监听器 #1
var listener1 = function listener1() {
    console.log('监听器 listener1 执行。');
}

// 监听器 #2
var listener2 = function listener2() {
    console.log('监听器 listener2 执行。');
}

// 绑定 connection 事件，处理函数为 listener1
eventEmitter.addListener('connection', listener1);

// 绑定 connection 事件，处理函数为 listener2
eventEmitter.on('connection', listener2);

var eventListeners = eventEmitter.listenerCount('connection');
console.log(eventListeners + " 个监听器监听连接事件。");

// 处理 connection 事件
eventEmitter.emit('connection');

// 移除监绑定的 listener1 函数
eventEmitter.removeListener('connection', listener1);
console.log("listener1 不再受监听。");

// 触发连接事件
eventEmitter.emit('connection');

eventListeners = eventEmitter.listenerCount('connection');
console.log(eventListeners + " 个监听器监听连接事件。");

console.log("程序执行完毕。");

function f2() {
    const buf = Buffer.from('runoob', 'ascii');

// 输出 72756e6f6f62
    console.log(buf.toString('hex'));

// 输出 cnVub29i
    console.log(buf.toString('base64'));
}

function f0() {
    buf = Buffer.alloc(256);
    len = buf.write("www.runoob.com");

    console.log("写入字节数 : "+  len);
}

function f1() {
    buf = Buffer.alloc(256);
    len = buf.write("www.runoob.com");

    console.log("写入字节数 : "+  len);
}

function f3() {
    const buf = Buffer.from('runoob', 'ascii');

// 输出 72756e6f6f62
    console.log(buf.toString('hex'));

// 输出 cnVub29i
    console.log(buf.toString('base64'));
}

function f4() {
    var buffer1 = Buffer.from(('菜鸟教程'));
    var buffer2 = Buffer.from(('www.runoob.com'));
    var buffer3 = Buffer.concat([buffer1,buffer2]);
    console.log("buffer3 内容: " + buffer3.toString());
}

function f5() {
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

function f6() {
    var buf1 = Buffer.from('abcdefghijkl');
    var buf2 = Buffer.from('RUNOOB');

//将 buf2 插入到 buf1 指定位置上
    buf2.copy(buf1, 2);

    console.log(buf1.toString());
}

function f7() {
    var buffer1 = Buffer.from('runoob');
// 剪切缓冲区
    var buffer2 = buffer1.slice(0,2);
    console.log("buffer2 content: " + buffer2.toString());
}

function f8() {
    var buffer = Buffer.from('www.runoob.com');
//  缓冲区长度
    console.log("buffer length: " + buffer.length);
}

var fs = require("fs");
var data = '';

// 创建可读流
var readerStream = fs.createReadStream('input.txt');

// 设置编码为 utf8。
readerStream.setEncoding('UTF8');

// 处理流事件 --> data, end, and error
readerStream.on('data', function(chunk) {
    data += chunk;
});

readerStream.on('end',function(){
    console.log(data);
});

readerStream.on('error', function(err){
    console.log(err.stack);
});

console.log("程序执行完毕");


var fs = require("fs");
var data = '菜鸟教程官网地址：www.runoob.com';

// 创建一个可以写入的流，写入到文件 output.txt 中
var writerStream = fs.createWriteStream('output.txt');

// 使用 utf8 编码写入数据
writerStream.write(data,'UTF8');

// 标记文件末尾
writerStream.end();

// 处理流事件 --> finish、error
writerStream.on('finish', function() {
    console.log("写入完成。");
});

writerStream.on('error', function(err){
    console.log(err.stack);
});

console.log("程序执行完毕");


var fs = require("fs");

// 创建一个可读流
var readerStream = fs.createReadStream('input.txt');

// 创建一个可写流
var writerStream = fs.createWriteStream('output.txt');

// 管道读写操作
// 读取 input.txt 文件内容，并将内容写入到 output.txt 文件中
readerStream.pipe(writerStream);

console.log("程序执行完毕");

var fs = require("fs");
var zlib = require('zlib');

// 压缩 input.txt 文件为 input.txt.gz
fs.createReadStream('input.txt')
    .pipe(zlib.createGzip())
    .pipe(fs.createWriteStream('input.txt.gz'));

console.log("文件压缩完成。");


var fs = require("fs");
var zlib = require('zlib');

// 解压 input.txt.gz 文件为 input.txt
fs.createReadStream('input.txt.gz')
    .pipe(zlib.createGunzip())
    .pipe(fs.createWriteStream('input.txt'));

console.log("文件解压完成。");

function say(word) {
    console.log(word);
}

function execute(someFunction, value) {
    someFunction(value);
}

execute(say, "Hello");


execute(function(word){ console.log(word) }, "Hello");

console.log( __filename );
console.log( __dirname );

function printHello(){
    console.log( "Hello, World!");
}
// 两秒后执行以上函数
setTimeout(printHello, 2000);

function printHello(){
    console.log( "Hello, World!");
}
// 两秒后执行以上函数
var t = setTimeout(printHello, 2000);

// 清除定时器
clearTimeout(t);


function printHello(){
    console.log( "Hello, World!");
}
// 两秒后执行以上函数
setInterval(printHello, 2000);

console.log('Hello world');
console.log('byvoid%diovyb');
console.log('byvoid%diovyb', 1991);


console.info("程序开始执行：");

var counter = 10;
console.log("计数: %d", counter);

console.time("获取数据");
//
// 执行一些代码
//
console.timeEnd('获取数据');

console.info("程序执行完毕。")

process.on('exit', function(code) {

    // 以下代码永远不会执行
    setTimeout(function() {
        console.log("该代码不会执行");
    }, 0);

    console.log('退出码为:', code);
});
console.log("程序执行结束");

process.stdout.write("Hello World!" + "\n");

// 通过参数读取
process.argv.forEach(function(val, index, array) {
    console.log(index + ': ' + val);
});

// 获取执行路径
console.log(process.execPath);


// 平台信息
console.log(process.platform);

// 输出当前目录
console.log('当前目录: ' + process.cwd());

// 输出当前版本
console.log('当前版本: ' + process.version);

// 输出内存使用情况
console.log(process.memoryUsage());