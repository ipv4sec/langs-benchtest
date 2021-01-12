package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

func main() {
	startAt := time.Now().UnixNano()

	for i:=0; i < 2; i++ {
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
		f23()
		f24()
		f25()
		f26()
		f27()
		f28()
		f29()


		f30()
		f31()
		f32()
		f33()
		f34()
		f35()
		f36()
		f37()
		f38()
		f39()

		f40()
		f400()
		f41()
		f42()
		f43()
		f44()
		f45()
		f46()
		f47()
		f48()
		f49()

		f50()
		f51()
		f52()
		f53()
		f54()
		f55()
		f56()
		f57()
	}

	fmt.Print(time.Now().UnixNano()- startAt)
}

func f1() {
	fmt.Println("Hello, World!")
	fmt.Println("Google" + "Runoob")
	var age int
	age = 10
	fmt.Println(age)
	apples := "a"
	oranges := "b"
	fruit := apples+oranges
	fmt.Println(fruit)

	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明一个变量并初始化
	var aa = "RUNOOB"
	fmt.Println(aa)

	// 没有初始化就为零值
	var bb int
	fmt.Println(bb)

	// bool 零值为 false
	var cc bool
	fmt.Println(cc)

	var i int
	var f float64
	var bbb bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bbb, s)

	var d = true
	fmt.Println(d)

	ff := "Runoob" // var f string = "Runoob"

	fmt.Println(ff)

}

func f2() {
	var x, y int
	var (  // 这种因式分解关键字的写法一般用于声明全局变量
		a int
		b bool
	)

	var c, d int = 1, 2
	var e, f = 123, "hello"

	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}


func f3() {
	var a string = "abc"
	fmt.Println("hello, world", a)
}


func f4() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}


func f5(){
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a, b, c)
}

func f6() {
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}


func f7() {
	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}

func f8() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c )
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c )
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c )
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c )
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a )
	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )
}

func f9() {
	var a int = 21
	var b int = 10

	if( a == b ) {
		fmt.Printf("第一行 - a 等于 b\n" )
	} else {
		fmt.Printf("第一行 - a 不等于 b\n" )
	}
	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n" )
	} else {
		fmt.Printf("第二行 - a 不小于 b\n" )
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n" )
	} else {
		fmt.Printf("第三行 - a 不大于 b\n" )
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n" )
	}
	if ( b >= a ) {
		fmt.Printf("第五行 - b 大于等于 a\n" )
	}
}

func f10() {
	var a bool = true
	var b bool = false
	if ( a && b ) {
		fmt.Printf("第一行 - 条件为 true\n" )
	}
	if ( a || b ) {
		fmt.Printf("第二行 - 条件为 true\n" )
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if ( a && b ) {
		fmt.Printf("第三行 - 条件为 true\n" )
	} else {
		fmt.Printf("第三行 - 条件为 false\n" )
	}
	if ( !(a && b) ) {
		fmt.Printf("第四行 - 条件为 true\n" )
	}
}

func f11() {

	var a uint = 60      /* 60 = 0011 1100 */
	var b uint = 13      /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b       /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c )

	c = a | b       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c )

	c = a ^ b       /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c )

	c = a << 2     /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c )

	c = a >> 2     /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c )
}

func f12() {
	var a int = 21
	var c int

	c =  a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

	c +=  a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

	c -=  a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

	c *=  a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

	c /=  a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

	c  = 200;

	c <<=  2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

	c >>=  2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

	c &=  2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

	c ^=  2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

	c |=  2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )

}

func f13() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

	/*  & 和 * 运算符实例 */
	ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a);
	fmt.Printf("*ptr 为 %d\n", *ptr);
}

func f14() {
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int;

	e = (a + b) * c / d;      // ( 30 * 15 ) / 5
	fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

	e = ((a + b) * c) / d;    // (30 * 15 ) / 5
	fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

	e = (a + b) * (c / d);   // (30) * (15/5)
	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

	e = a + (b * c) / d;     //  20 + (150/5)
	fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );
}

func f15() {
	/* 定义局部变量 */
	var a int = 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n" )
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func f16() {
	/* 局部变量定义 */
	var a int = 100;

	/* 判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n" );
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n" );
	}
	fmt.Printf("a 的值为 : %d\n", a);

}

func f17() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
		}
	}
	fmt.Printf("a 值为 : %d\n", a );
	fmt.Printf("b 值为 : %d\n", b );
}

func f18() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade );
}

func f19() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 ", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
	}
}

func f20() {
	var a = 1

	switch {
	case a==2:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case a==1:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case a==3:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case a==1:
		fmt.Println("4、case 条件语句为 true")
	case a==4:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

func f21() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func f22() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a);
		a++;
		if a > 15 {
			/* 使用 break 语句跳出循环 */
			break;
		}
	}
}
func f23() {

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}
func f24() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1;
			continue;
		}
		fmt.Printf("a 的值为 : %d\n", a);
		a++;
	}
}

func f25() {

	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue re
		}
	}
}

func f26() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP: for a < 20 {
	if a == 15 {
		/* 跳过迭代 */
		a = a + 1
		goto LOOP
	}
	fmt.Printf("a的值为 : %d\n", a)
	a++
}
}


func f27() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf( "最大值是 : %d\n", ret )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func f28() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}

func f29(){
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

}

func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

func f30(){
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

type Circle struct {
	radius float64
}

func f31() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
func f32() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
}

var g int

func f33() {

	/* 声明局部变量 */
	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}

var a int = 20;

func f34() {
	/* main 函数中声明局部变量 */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中 a = %d\n",  a);
	c = sum( a, b);
	fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n",  a);
	fmt.Printf("sum() 函数中 b = %d\n",  b);

	return a + b;
}

func f35() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}
}

func f36() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a  )
}

func f37() {
	var a int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}
func f38() {
	var  ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr  )
}



func f39() {
	type Books struct {
		title string
		author string
		subject string
		book_id int
	}
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func f40() {

	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}



func f400() {

	printBook := func( book Books ) {
		fmt.Printf( "Book title : %s\n", book.title)
		fmt.Printf( "Book author : %s\n", book.author)
		fmt.Printf( "Book subject : %s\n", book.subject)
		fmt.Printf( "Book book_id : %d\n", book.book_id)
	}

	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)
}





func f41() {
	type Books struct {
		title string
		author string
		subject string
		book_id int
	}
	printBook := func ( book *Books ) {
		fmt.Printf( "Book title : %s\n", book.title)
		fmt.Printf( "Book author : %s\n", book.author)
		fmt.Printf( "Book subject : %s\n", book.subject)
		fmt.Printf( "Book book_id : %d\n", book.book_id)
	}
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
}

func f42() {
	var numbers = make([]int,3,5)

	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func f43() {

	printSlice := func (x []int){
		fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
	}
	var numbers []int

	printSlice(numbers)

	if(numbers == nil){
		fmt.Printf("切片是空的")
	}
}



func f44() {
	printSlice := func (x []int){
		fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
	}
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func f45() {
	printSlice := func (x []int){
		fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
	}
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)
}



func f46() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func f47() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func f48() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

	/*删除元素*/ delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
}
func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func f49() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func f50() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func f51() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func f52() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}


// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func f53() {

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func f54() {
	go say("world")
	say("hello")
}



func f55() {
	sum := func (s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum // 把 sum 发送到通道 c
	}
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

func f56() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}



func f57() {
	fibonacci := func (n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}


}