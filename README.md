
# 某一种语言在某一种环境下的性能表现

## 某些语言选型

暂定 `Go` `C` `Rust` `Java` `Nodejs` `Python` `Php`

## 测试代码原则

把基础教程里面的教学代码循环执行一万遍输出执行时间, 测试10次输出平均时间

[go-tutorial](https://www.runoob.com/go/go-tutorial.html)
[c-tutorial](https://www.runoob.com/cprogramming/c-tutorial.html)
[rust-tutorial](https://www.runoob.com/rust/rust-tutorial.html)
[java-tutorial](https://www.runoob.com/java/java-tutorial.html)
[nodejs-tutorial](https://www.runoob.com/nodejs/nodejs-tutorial.html)
[python-tutorial](https://www.runoob.com/python3/python3-tutorial.html)
[php-tutorial](https://www.runoob.com/php/php-tutorial.html)


## 环境

基础环境安装
```shell script
[root@i-okjxke92 ~]# cat /etc/redhat-release 
CentOS Linux release 7.9.2009 (Core)

[root@i-okjxke92 ~]# gcc -v
使用内建 specs。
COLLECT_GCC=gcc
COLLECT_LTO_WRAPPER=/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/lto-wrapper
目标：x86_64-redhat-linux
配置为：../configure --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --with-bugurl=http://bugzilla.redhat.com/bugzilla --enable-bootstrap --enable-shared --enable-threads=posix --enable-checking=release --with-system-zlib --enable-__cxa_atexit --disable-libunwind-exceptions --enable-gnu-unique-object --enable-linker-build-id --with-linker-hash-style=gnu --enable-languages=c,c++,objc,obj-c++,java,fortran,ada,go,lto --enable-plugin --enable-initfini-array --disable-libgcj --with-isl=/builddir/build/BUILD/gcc-4.8.5-20150702/obj-x86_64-redhat-linux/isl-install --with-cloog=/builddir/build/BUILD/gcc-4.8.5-20150702/obj-x86_64-redhat-linux/cloog-install --enable-gnu-indirect-function --with-tune=generic --with-arch_32=x86-64 --build=x86_64-redhat-linux
线程模型：posix
gcc 版本 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC)
[root@i-okjxke92 ~]# yum install python3 -y
[root@i-okjxke92 ~]# python3 --version
Python 3.6.8

yum install epel-release -y
```

### Wasmtime

官网 https://wasmtime.dev/
安装 `curl https://wasmtime.dev/install.sh -sSf | bash`
配置 `source ~/.bashrc`


### Wasmer

官网 https://wasmer.io/
安装 `curl https://get.wasmer.io -sSfL | sh`

安装错误, 错误信息如下:
```shell script
[root@i-okjxke92 ~]# curl https://get.wasmer.io -sSfL | sh
which: no wasmer in (/root/.wasmtime/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin)
Welcome to the Wasmer bash installer!

               ww
               wwwww
        ww     wwwwww  w
        wwwww      wwwwwwwww
ww      wwwwww  w     wwwwwww
wwwww      wwwwwwwwww   wwwww
wwwwww  w      wwwwwww  wwwww
wwwwwwwwwwwwww   wwwww  wwwww
wwwwwwwwwwwwwww  wwwww  wwwww
wwwwwwwwwwwwwww  wwwww  wwwww
wwwwwwwwwwwwwww  wwwww  wwwww
wwwwwwwwwwwwwww  wwwww   wwww
wwwwwwwwwwwwwww  wwwww
   wwwwwwwwwwww   wwww
       wwwwwwww
           wwww

downloading: wasmer-linux-amd64
Latest release: 1.0.0
which: no wasmer in (/root/.wasmer/bin)
Downloading archive from https://github.com/wasmerio/wasmer/releases/download/1.0.0/wasmer-linux-amd64.tar.gz
######################################################################## 100.0%
installing: /root/.wasmer
Updating bash profile /root/.bashrc
we've added the following to your /root/.bashrc
If you have a different profile please add the following:

# Wasmer
export WASMER_DIR="/root/.wasmer"
[ -s "$WASMER_DIR/wasmer.sh" ] && source "$WASMER_DIR/wasmer.sh"
/root/.wasmer/bin/wasmer: /lib64/libm.so.6: version `GLIBC_2.27' not found (required by /root/.wasmer/bin/wasmer)
/root/.wasmer/bin/wasmer: /lib64/libstdc++.so.6: version `GLIBCXX_3.4.20' not found (required by /root/.wasmer/bin/wasmer)
/root/.wasmer/bin/wasmer: /lib64/libstdc++.so.6: version `CXXABI_1.3.9' not found (required by /root/.wasmer/bin/wasmer)
/root/.wasmer/bin/wasmer: /lib64/libstdc++.so.6: version `GLIBCXX_3.4.21' not found (required by /root/.wasmer/bin/wasmer)
/root/.wasmer/bin/wasmer: /lib64/libc.so.6: version `GLIBC_2.18' not found (required by /root/.wasmer/bin/wasmer)
error: wasmer was installed, but doesn't seem to be working :(
```

安装失败 https://github.com/wasmerio/wasmer/pull/1755
需要升级 GLIBC 到 2.17, GLIBCXX 升级到 3.4.19



###  Graalvm Native
TODO

### Graalvm
TODO

### Native
TODO

## 具体的语言

### Go

#### 转换成 `WebAssembly`

依据 https://github.com/appcypher/awesome-wasm-langs#go, 有两种方式转换成WebAssembly
1. Golang https://github.com/golang/go
2. tinygo https://github.com/tinygo-org/tinygo

##### Golang 官方

测试代码:
```
package main

import (
	"fmt"
)
func main() {
	fmt.Println("Hello, Go")
}

```

编译:
```shell script
GOOS=js GOARCH=wasm go build -o t.wasm t.go
```

在node中运行
```
node $(go env GOROOT)/misc/wasm/wasm_exec.js t.wasm
```

查询Go能编译到哪些平台和哪些分支
```shell script
➜  ~ go tool dist list |grep wasm                       
js/wasm
```
仅支持js平台

用graalvm运行编译的wasm文件, 错误
```
[root@acg go]# wasm main.wasm 
ERROR: No entry-point function found, cannot start program.
```

在mac下用wasmer运行编译的wasm文件, 错误
```
➜  go git:(master) ✗ wasmer main.wasm
error: failed to run `main.wasm`
╰─> 1: Error while importing "go"."debug": unknown import. Expected Function(FunctionType { params: [I32], results: [] })
```

在mac下用wasmtime运行编译的wasm文件, 错误
```
➜  go git:(master) ✗ wasmtime main.wasm
Error: failed to run main module `main.wasm`

Caused by:
    0: failed to instantiate "main.wasm"
    1: unknown import: `go::debug` has not been defined
```

结论: 可以转换成 wasm 格式, 但仅可以在 js 中调用

##### tinygo

官方安装命令错误
```shell script
[root@i-okjxke92 ~]# sudo dnf install tinygo
上次元数据过期检查：0:01:09 前，执行于 2021年01月12日 星期二 15时30分26秒。
未找到匹配的参数： tinygo
错误：没有任何匹配: tinygo
```

在mac下用wasmtime运行用mac的tinygo编译的wasm文件, 错误
```
➜  go git:(master) ✗ wasmtime tiny.wasm
Error: failed to run main module `tiny.wasm`

Caused by:
    0: failed to instantiate "tiny.wasm"
    1: unknown import: `env::syscall/js.valueGet` has not been defined
```

用graalvm运行用mac的tinygo编译的wasm文件, 错误
```
[root@acg go]# wasm tiny.wasm 
ERROR: The module 'env', referenced by the import 'syscall/js.valueNew' in the module 'main', does not exist.
```

在mac下用wasmer运行用mac的tinygo编译的wasm文件, 错误
```
➜  go git:(master) ✗ wasmer tiny.wasm
error: failed to run `tiny.wasm`
╰─> 1: Error while importing "wasi_unstable"."fd_write": unknown import. Expected Function(FunctionType { params: [I32, I32, I32, I32], results: [I32] })
➜  go git:(master) ✗ 
```

#### 运行结果

1. yum install golang1.15.5 运行时间(纳秒): 1002707570
2. 官网下载golang1.15.5 运行时间(纳秒): 1002898515
3. 编译安装golang1.15.5 运行时间(纳秒): 1002841179

(一开始测试的是官网下载golang1.15.6 运行时间(纳秒): 1002665315, 为了和yum版本一致就没用此版本)

![](./go/ECharts.png)

### Java
yum install java-11-openjdk -y


```
[root@i-okjxke92 java]# native-image Main
[main:23021]    classlist:   3,765.57 ms,  0.96 GB
[main:23021]        (cap):     939.82 ms,  0.96 GB
[main:23021]        setup:   3,241.91 ms,  0.96 GB
[main:23021]     (clinit):     449.99 ms,  1.06 GB
[main:23021]   (typeflow):  13,203.28 ms,  1.06 GB
[main:23021]    (objects):   8,329.93 ms,  1.06 GB
[main:23021]   (features):     356.95 ms,  1.06 GB
[main:23021]     analysis:  22,653.08 ms,  1.06 GB
[main:23021]     universe:     722.38 ms,  1.06 GB
[main:23021]      (parse):   5,198.06 ms,  1.05 GB
[main:23021]     (inline):   3,783.86 ms,  1.01 GB
[main:23021]    (compile):  23,500.11 ms,  1.05 GB
[main:23021]      compile:  33,155.32 ms,  1.05 GB
[main:23021]        image:   1,857.48 ms,  1.05 GB
[main:23021]        write:     171.04 ms,  1.05 GB
Fatal error:java.lang.RuntimeException: java.lang.RuntimeException: There was an error linking the native image: Linker command exited with 1

Linker command executed:
/usr/bin/gcc -z noexecstack -Wl,--gc-sections -Wl,--dynamic-list -Wl,/tmp/SVM-10949438881074898817/exported_symbols.list -Wl,--exclude-libs,ALL -Wl,-x -o /root/langs-benchtest/java/main main.o /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/liblibchelper.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnet.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libffi.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnio.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libjava.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libfdlibm.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libzip.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libjvm.a -v -L/tmp/SVM-10949438881074898817 -L/root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc -L/root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64 -lpthread -ldl -lz -lrt 

Linker command output:
使用内建 specs。
COLLECT_GCC=/usr/bin/gcc
COLLECT_LTO_WRAPPER=/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/lto-wrapper
目标：x86_64-redhat-linux
配置为：../configure --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --with-bugurl=http://bugzilla.redhat.com/bugzilla --enable-bootstrap --enable-shared --enable-threads=posix --enable-checking=release --with-system-zlib --enable-__cxa_atexit --disable-libunwind-exceptions --enable-gnu-unique-object --enable-linker-build-id --with-linker-hash-style=gnu --enable-languages=c,c++,objc,obj-c++,java,fortran,ada,go,lto --enable-plugin --enable-initfini-array --disable-libgcj --with-isl=/builddir/build/BUILD/gcc-4.8.5-20150702/obj-x86_64-redhat-linux/isl-install --with-cloog=/builddir/build/BUILD/gcc-4.8.5-20150702/obj-x86_64-redhat-linux/cloog-install --enable-gnu-indirect-function --with-tune=generic --with-arch_32=x86-64 --build=x86_64-redhat-linux
线程模型：posix
gcc 版本 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC) 
COMPILER_PATH=/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/:/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/:/usr/libexec/gcc/x86_64-redhat-linux/:/usr/lib/gcc/x86_64-redhat-linux/4.8.5/:/usr/lib/gcc/x86_64-redhat-linux/
LIBRARY_PATH=/usr/lib/gcc/x86_64-redhat-linux/4.8.5/:/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/:/lib/../lib64/:/usr/lib/../lib64/:/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../:/lib/:/usr/lib/
COLLECT_GCC_OPTIONS='-z' 'noexecstack' '-o' '/root/langs-benchtest/java/main' '-v' '-L/tmp/SVM-10949438881074898817' '-L/root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc' '-L/root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64' '-mtune=generic' '-march=x86-64'
 /usr/libexec/gcc/x86_64-redhat-linux/4.8.5/collect2 --build-id --no-add-needed --eh-frame-hdr --hash-style=gnu -m elf_x86_64 -dynamic-linker /lib64/ld-linux-x86-64.so.2 -o /root/langs-benchtest/java/main -z noexecstack /usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/crt1.o /usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/4.8.5/crtbegin.o -L/tmp/SVM-10949438881074898817 -L/root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc -L/root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64 -L/usr/lib/gcc/x86_64-redhat-linux/4.8.5 -L/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../.. --gc-sections --dynamic-list /tmp/SVM-10949438881074898817/exported_symbols.list --exclude-libs ALL -x main.o /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/liblibchelper.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnet.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libffi.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnio.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libjava.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libfdlibm.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libzip.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libjvm.a -lpthread -ldl -lz -lrt -lgcc --as-needed -lgcc_s --no-as-needed -lc -lgcc --as-needed -lgcc_s --no-as-needed /usr/lib/gcc/x86_64-redhat-linux/4.8.5/crtend.o /usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/crtn.o
/usr/bin/ld: 找不到 -lz
collect2: 错误：ld 返回 1
	at java.base/jdk.internal.reflect.NativeConstructorAccessorImpl.newInstance0(Native Method)
	at java.base/jdk.internal.reflect.NativeConstructorAccessorImpl.newInstance(NativeConstructorAccessorImpl.java:62)
	at java.base/jdk.internal.reflect.DelegatingConstructorAccessorImpl.newInstance(DelegatingConstructorAccessorImpl.java:45)
	at java.base/java.lang.reflect.Constructor.newInstance(Constructor.java:490)
	at java.base/java.util.concurrent.ForkJoinTask.getThrowableException(ForkJoinTask.java:600)
	at java.base/java.util.concurrent.ForkJoinTask.get(ForkJoinTask.java:1006)
	at com.oracle.svm.hosted.NativeImageGenerator.run(NativeImageGenerator.java:483)
	at com.oracle.svm.hosted.NativeImageGeneratorRunner.buildImage(NativeImageGeneratorRunner.java:350)
	at com.oracle.svm.hosted.NativeImageGeneratorRunner.build(NativeImageGeneratorRunner.java:509)
	at com.oracle.svm.hosted.NativeImageGeneratorRunner.main(NativeImageGeneratorRunner.java:115)
	at com.oracle.svm.hosted.NativeImageGeneratorRunner$JDK9Plus.main(NativeImageGeneratorRunner.java:541)
Caused by: java.lang.RuntimeException: There was an error linking the native image: Linker command exited with 1

Linker command executed:
/usr/bin/gcc -z noexecstack -Wl,--gc-sections -Wl,--dynamic-list -Wl,/tmp/SVM-10949438881074898817/exported_symbols.list -Wl,--exclude-libs,ALL -Wl,-x -o /root/langs-benchtest/java/main main.o /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/liblibchelper.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnet.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libffi.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnio.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libjava.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libfdlibm.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libzip.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libjvm.a -v -L/tmp/SVM-10949438881074898817 -L/root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc -L/root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64 -lpthread -ldl -lz -lrt 

Linker command output:
使用内建 specs。
COLLECT_GCC=/usr/bin/gcc
COLLECT_LTO_WRAPPER=/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/lto-wrapper
目标：x86_64-redhat-linux
配置为：../configure --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --with-bugurl=http://bugzilla.redhat.com/bugzilla --enable-bootstrap --enable-shared --enable-threads=posix --enable-checking=release --with-system-zlib --enable-__cxa_atexit --disable-libunwind-exceptions --enable-gnu-unique-object --enable-linker-build-id --with-linker-hash-style=gnu --enable-languages=c,c++,objc,obj-c++,java,fortran,ada,go,lto --enable-plugin --enable-initfini-array --disable-libgcj --with-isl=/builddir/build/BUILD/gcc-4.8.5-20150702/obj-x86_64-redhat-linux/isl-install --with-cloog=/builddir/build/BUILD/gcc-4.8.5-20150702/obj-x86_64-redhat-linux/cloog-install --enable-gnu-indirect-function --with-tune=generic --with-arch_32=x86-64 --build=x86_64-redhat-linux
线程模型：posix
gcc 版本 4.8.5 20150623 (Red Hat 4.8.5-44) (GCC) 
COMPILER_PATH=/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/:/usr/libexec/gcc/x86_64-redhat-linux/4.8.5/:/usr/libexec/gcc/x86_64-redhat-linux/:/usr/lib/gcc/x86_64-redhat-linux/4.8.5/:/usr/lib/gcc/x86_64-redhat-linux/
LIBRARY_PATH=/usr/lib/gcc/x86_64-redhat-linux/4.8.5/:/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/:/lib/../lib64/:/usr/lib/../lib64/:/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../:/lib/:/usr/lib/
COLLECT_GCC_OPTIONS='-z' 'noexecstack' '-o' '/root/langs-benchtest/java/main' '-v' '-L/tmp/SVM-10949438881074898817' '-L/root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc' '-L/root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64' '-mtune=generic' '-march=x86-64'
 /usr/libexec/gcc/x86_64-redhat-linux/4.8.5/collect2 --build-id --no-add-needed --eh-frame-hdr --hash-style=gnu -m elf_x86_64 -dynamic-linker /lib64/ld-linux-x86-64.so.2 -o /root/langs-benchtest/java/main -z noexecstack /usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/crt1.o /usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/4.8.5/crtbegin.o -L/tmp/SVM-10949438881074898817 -L/root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc -L/root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64 -L/usr/lib/gcc/x86_64-redhat-linux/4.8.5 -L/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../.. --gc-sections --dynamic-list /tmp/SVM-10949438881074898817/exported_symbols.list --exclude-libs ALL -x main.o /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/liblibchelper.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnet.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libffi.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libnio.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libjava.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libfdlibm.a /root/downloads/graalvm-ce-java11-20.3.0/lib/static/linux-amd64/glibc/libzip.a /root/downloads/graalvm-ce-java11-20.3.0/lib/svm/clibraries/linux-amd64/libjvm.a -lpthread -ldl -lz -lrt -lgcc --as-needed -lgcc_s --no-as-needed -lc -lgcc --as-needed -lgcc_s --no-as-needed /usr/lib/gcc/x86_64-redhat-linux/4.8.5/crtend.o /usr/lib/gcc/x86_64-redhat-linux/4.8.5/../../../../lib64/crtn.o
/usr/bin/ld: 找不到 -lz
collect2: 错误：ld 返回 1
	at com.oracle.svm.hosted.image.NativeBootImageViaCC.handleLinkerFailure(NativeBootImageViaCC.java:463)
	at com.oracle.svm.hosted.image.NativeBootImageViaCC.write(NativeBootImageViaCC.java:430)
	at com.oracle.svm.hosted.NativeImageGenerator.doRun(NativeImageGenerator.java:680)
	at com.oracle.svm.hosted.NativeImageGenerator.lambda$run$0(NativeImageGenerator.java:471)
	at java.base/java.util.concurrent.ForkJoinTask$AdaptedRunnableAction.exec(ForkJoinTask.java:1407)
	at java.base/java.util.concurrent.ForkJoinTask.doExec(ForkJoinTask.java:290)
	at java.base/java.util.concurrent.ForkJoinPool$WorkQueue.topLevelExec(ForkJoinPool.java:1020)
	at java.base/java.util.concurrent.ForkJoinPool.scan(ForkJoinPool.java:1656)
	at java.base/java.util.concurrent.ForkJoinPool.runWorker(ForkJoinPool.java:1594)
	at java.base/java.util.concurrent.ForkJoinWorkerThread.run(ForkJoinWorkerThread.java:183)
Error: Image build request failed with exit status 1
```
yum install zlib-devel -y

#### 运行结果
1. yum install jdk11 运行时间(纳秒):  12881111702
2. graalvm 运行时间(纳秒): 12841016645
3. graalvm native 运行时间(纳秒): 