package com.example.zed;

import java.util.Calendar;
import java.util.Date;
import java.util.GregorianCalendar;
import java.util.Locale;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class FreshJuice {
    enum FreshJuiceSize{ SMALL, MEDIUM , LARGE }
    FreshJuiceSize size;
}

class Employee{
    String name;
    int age;
    String designation;
    double salary;
    // Employee 类的构造器
    public Employee(String name){
        this.name = name;
    }
    // 设置age的值
    public void empAge(int empAge){
        age =  empAge;
    }
    /* 设置designation的值*/
    public void empDesignation(String empDesig){
        designation = empDesig;
    }
    /* 设置salary的值*/
    public void empSalary(double empSalary){
        salary = empSalary;
    }
    /* 打印信息 */
    public void printEmployee(){
        System.out.println("名字:"+ name );
        System.out.println("年龄:" + age );
        System.out.println("职位:" + designation );
        System.out.println("薪水:" + salary);
    }
}

public class Zed {
    private static boolean myFlag = true;
    static final double weeks = 9.5;
    protected static final int BOXWIDTH = 42;
    public transient int limit = 55;

    public static void main(String[] args) {
        // TODO Auto-generated method stub

        long startTime = System.currentTimeMillis() * 1000000L + System.nanoTime() % 1000000L;    //获取开始时间
        for(int i = 0; i < 10000; i++) {
            print_HelloWorld();
            System.out.println(myFlag);
            System.out.println(weeks);
            FreshJuice juice = new FreshJuice();
            juice.size = FreshJuice.FreshJuiceSize.SMALL;
            System.out.println("juice.size = "+juice.size);
            /* 使用构造器创建两个对象 */
            Employee empOne = new Employee("RUNOOB1");
            Employee empTwo = new Employee("RUNOOB2");

            // 调用这两个对象的成员方法
            empOne.empAge(26);
            empOne.empDesignation("高级程序员");
            empOne.empSalary(1000);
            empOne.printEmployee();

            empTwo.empAge(21);
            empTwo.empDesignation("菜鸟程序员");
            empTwo.empSalary(500);
            empTwo.printEmployee();
            PrimitiveTypeTest();
            ZiDongLeiZhuan();
            QiangZhiZhuanHuan();
            operator();
            LoopTest();
            ConditionStatementTest();
            Test();
            StringTest();
            ArrayTest();
            printArray(new int[]{3, 1, 2, 6, 4, 2});
            DateTest();
            RegexExample();
            System.out.println(max(5, 6));
        }
        long endTime = System.currentTimeMillis() * 1000000L + System.nanoTime() % 1000000L;    //获取开始时间

        System.out.println(endTime - startTime);    //输出程序运行时间
    }

    public static void print_HelloWorld() {
        System.out.println("Hello,World");
    }

    public static void PrimitiveTypeTest() {
        // byte
        System.out.println("基本类型：byte 二进制位数：" + Byte.SIZE);
        System.out.println("包装类：java.lang.Byte");
        System.out.println("最小值：Byte.MIN_VALUE=" + Byte.MIN_VALUE);
        System.out.println("最大值：Byte.MAX_VALUE=" + Byte.MAX_VALUE);
        System.out.println();

        // short
        System.out.println("基本类型：short 二进制位数：" + Short.SIZE);
        System.out.println("包装类：java.lang.Short");
        System.out.println("最小值：Short.MIN_VALUE=" + Short.MIN_VALUE);
        System.out.println("最大值：Short.MAX_VALUE=" + Short.MAX_VALUE);
        System.out.println();

        // int
        System.out.println("基本类型：int 二进制位数：" + Integer.SIZE);
        System.out.println("包装类：java.lang.Integer");
        System.out.println("最小值：Integer.MIN_VALUE=" + Integer.MIN_VALUE);
        System.out.println("最大值：Integer.MAX_VALUE=" + Integer.MAX_VALUE);
        System.out.println();

        // long
        System.out.println("基本类型：long 二进制位数：" + Long.SIZE);
        System.out.println("包装类：java.lang.Long");
        System.out.println("最小值：Long.MIN_VALUE=" + Long.MIN_VALUE);
        System.out.println("最大值：Long.MAX_VALUE=" + Long.MAX_VALUE);
        System.out.println();

        // float
        System.out.println("基本类型：float 二进制位数：" + Float.SIZE);
        System.out.println("包装类：java.lang.Float");
        System.out.println("最小值：Float.MIN_VALUE=" + Float.MIN_VALUE);
        System.out.println("最大值：Float.MAX_VALUE=" + Float.MAX_VALUE);
        System.out.println();

        // double
        System.out.println("基本类型：double 二进制位数：" + Double.SIZE);
        System.out.println("包装类：java.lang.Double");
        System.out.println("最小值：Double.MIN_VALUE=" + Double.MIN_VALUE);
        System.out.println("最大值：Double.MAX_VALUE=" + Double.MAX_VALUE);
        System.out.println();

        // char
        System.out.println("基本类型：char 二进制位数：" + Character.SIZE);
        System.out.println("包装类：java.lang.Character");
        // 以数值形式而不是字符形式将Character.MIN_VALUE输出到控制台
        System.out.println("最小值：Character.MIN_VALUE="
                + (int) Character.MIN_VALUE);
        // 以数值形式而不是字符形式将Character.MAX_VALUE输出到控制台
        System.out.println("最大值：Character.MAX_VALUE="
                + (int) Character.MAX_VALUE);

        final double PI = 3.1415927;
        byte a = 68;
        char b = 'A';
    }

    public static void ZiDongLeiZhuan(){
        char c1='a';//定义一个char类型
        int i1 = c1;//char自动类型转换为int
        System.out.println("char自动类型转换为int后的值等于"+i1);
        char c2 = 'A';//定义一个char类型
        int i2 = c2+1;//char 类型和 int 类型计算
        System.out.println("char类型和int计算后的值等于"+i2);
    }

    public static void QiangZhiZhuanHuan(){
        int i1 = 123;
        byte b = (byte)i1;//强制类型转换为byte
        System.out.println("int强制类型转换为byte后的值等于"+b);
    }

    public static void operator() {
        int a = 10;
        int b = 20;
        int c = 25;
        int d = 25;
        System.out.println("算术运算符---------:");
        System.out.println("a + b = " + (a + b) );
        System.out.println("a - b = " + (a - b) );
        System.out.println("a * b = " + (a * b) );
        System.out.println("b / a = " + (b / a) );
        System.out.println("b % a = " + (b % a) );
        System.out.println("c % a = " + (c % a) );
        System.out.println("a++   = " +  (a++) );
        System.out.println("a--   = " +  (a--) );
        // 查看  d++ 与 ++d 的不同
        System.out.println("d++   = " +  (d++) );
        System.out.println("++d   = " +  (++d) );

        System.out.println("自增自减运算符----------:");
        int a1 = 5;//定义一个变量；
        int b1 = 5;
        int x1 = 2*++a1;
        int y1 = 2*b1++;
        System.out.println("自增运算符前缀运算后a1="+a1+",x1="+x1);
        System.out.println("自增运算符后缀运算后b1="+b1+",y1="+y1);

        System.out.println("关系运算符-------------------:");
        int a2 = 10;
        int b2 = 20;
        System.out.println("a2 == b2 = " + (a2 == b2) );
        System.out.println("a2 != b2 = " + (a2 != b2) );
        System.out.println("a2 > b2 = " + (a2 > b2) );
        System.out.println("a2 < b2 = " + (a2 < b2) );
        System.out.println("b2 >= a2 = " + (b2 >= a2) );
        System.out.println("b2 <= a2 = " + (b2 <= a2) );

        System.out.println("位运算符");
        int a3 = 60; /* 60 = 0011 1100 */
        int b3 = 13; /* 13 = 0000 1101 */
        int c3 = 0;
        c3 = a3 & b3;       /* 12 = 0000 1100 */
        System.out.println("a3 & b3 = " + c3 );

        c3 = a3 | b3;       /* 61 = 0011 1101 */
        System.out.println("a3 | b3 = " + c3 );

        c3 = a3 ^ b3;       /* 49 = 0011 0001 */
        System.out.println("a3 ^ b3 = " + c3 );

        c3 = ~a3;          /*-61 = 1100 0011 */
        System.out.println("~a3 = " + c3 );

        c3 = a3 << 2;     /* 240 = 1111 0000 */
        System.out.println("a3 << 2 = " + c3 );

        c3 = a3 >> 2;     /* 15 = 1111 */
        System.out.println("a3 >> 2  = " + c3 );

        c3 = a3 >>> 2;     /* 15 = 0000 1111 */
        System.out.println("a3 >>> 2 = " + c3 );

        System.out.println("逻辑运算符---------:");
        boolean a4 = true;
        boolean b4 = false;
        System.out.println("a4 && b4 = " + (a4&&b4));
        System.out.println("a4 || b4 = " + (a4||b4) );
        System.out.println("!(a4 && b4) = " + !(a4 && b4));

        System.out.println("短路逻辑运算符---------:");
        int a5 = 5;//定义一个变量；
        boolean b5 = (a5<4)&&(a5++<10);
        System.out.println("使用短路逻辑运算符的结果为"+b5);
        System.out.println("a的结果为"+a5);

        System.out.println("赋值运算符---------:");
        int a6 = 10;
        int b6 = 20;
        int c6 = 0;
        c6 = a6 + b6;
        System.out.println("c6 = a6 + b6 = " + c6 );
        c6 += a6 ;
        System.out.println("c6 += a6  = " + c6 );
        c6 -= a6 ;
        System.out.println("c6 -= a6 = " + c6 );
        c6 *= a6 ;
        System.out.println("c6 *= a6 = " + c6 );
        a6 = 10;
        c6 = 15;
        c6 /= a6 ;
        System.out.println("c6 /= a6 = " + c6 );
        a6 = 10;
        c6 = 15;
        c6 %= a6 ;
        System.out.println("c6 %= a6  = " + c6 );
        c6 <<= 2 ;
        System.out.println("c6 <<= 2 = " + c6 );
        c6 >>= 2 ;
        System.out.println("c6 >>= 2 = " + c6 );
        c6 >>= 2 ;
        System.out.println("c6 >>= 2 = " + c6 );
        c6 &= a6 ;
        System.out.println("c6 &= a6  = " + c6 );
        c6 ^= a6 ;
        System.out.println("c6 ^= a6   = " + c6 );
        c6 |= a6 ;
        System.out.println("c6 |= a6   = " + c6 );

        System.out.println("条件运算符---------:");
        int a7 , b7;
        a7 = 10;
        // 如果 a7 等于 1 成立，则设置 b7 为 20，否则为 30
        b7 = (a7 == 1) ? 20 : 30;
        System.out.println( "Value of b7 is : " +  b7 );

        // 如果 a 等于 10 成立，则设置 b 为 20，否则为 30
        b7 = (a7 == 10) ? 20 : 30;
        System.out.println( "Value of b7 is : " + b7 );
    }

    public static void LoopTest() {
        System.out.println("while循环-------:");
        int x = 10;
        while( x < 20 ) {
            System.out.print("value of x : " + x );
            x++;
            System.out.print("\n");
        }

        System.out.println("do⋯while循环-------:");
        int x1 = 10;

        do{
            System.out.print("value of x1 : " + x1 );
            x1++;
            System.out.print("\n");
        }while( x1 < 20 );

        System.out.println("for循环-------:");
        for(int x2 = 10; x2 < 20; x2 = x2+1) {
            System.out.print("value of x2 : " + x2 );
            System.out.print("\n");
        }

        System.out.println("for循环加强-------:");
        int [] numbers = {10, 20, 30, 40, 50};

        for(int x3 : numbers ){
            System.out.print( x3 );
            System.out.print(",");
        }
        System.out.print("\n");
        String [] names ={"James", "Larry", "Tom", "Lacy"};
        for( String name : names ) {
            System.out.print( name );
            System.out.print(",");
        }

        System.out.println("break关键字-------:");
        int [] numbers1 = {10, 20, 30, 40, 50};

        for(int x4 : numbers ) {
            // x 等于 30 时跳出循环
            if( x4 == 30 ) {
                break;
            }
            System.out.print( x4 );
            System.out.print("\n");
        }

        System.out.println("continue关键字-------:");
        int [] numbers2 = {10, 20, 30, 40, 50};

        for(int x5 : numbers ) {
            if( x5 == 30 ) {
                continue;
            }
            System.out.print( x5 );
            System.out.print("\n");
        }
    }

    public static void ConditionStatementTest() {
        System.out.println("if语句----------:");
        int x = 10;

        if( x < 20 ){
            System.out.println("这是 if 语句");
        }

        System.out.println("if-else语句----------:");
        int x1 = 30;

        if( x1 < 20 ){
            System.out.println("这是 if 语句");
        }else{
            System.out.println("这是 else 语句");
        }

        System.out.println("if...else if...else 语句----------:");
        int x2 = 30;

        if( x2 == 10 ){
            System.out.println("Value of X2 is 10");
        }else if( x2 == 20 ){
            System.out.println("Value of X2 is 20");
        }else if( x2 == 30 ){
            System.out.println("Value of X2 is 30");
        }else{
            System.out.println("这是 else 语句");
        }

        System.out.println("嵌套的 if⋯else 语句----------:");
        int x3 = 30;
        int y3 = 10;

        if( x3 == 30 ){
            if( y3 == 10 ){
                System.out.println("X3 = 30 and Y3 = 10");
            }
        }

        System.out.println("switch case 语句----------:");
        char grade = 'C';

        switch(grade)
        {
            case 'A' :
                System.out.println("优秀");
                break;
            case 'B' :
            case 'C' :
                System.out.println("良好");
                break;
            case 'D' :
                System.out.println("及格");
                break;
            case 'F' :
                System.out.println("你需要再努力努力");
                break;
            default :
                System.out.println("未知等级");
        }
        System.out.println("你的等级是 " + grade);
    }

    public static void Test() {
        System.out.println("Number类测试------:");
        Integer x = 5;
        x =  x + 10;
        System.out.println(x);

        System.out.println("Math类测试--------:");
        System.out.println("90 度的正弦值：" + Math.sin(Math.PI/2));
        System.out.println("0度的余弦值：" + Math.cos(0));
        System.out.println("60度的正切值：" + Math.tan(Math.PI/3));
        System.out.println("1的反正切值： " + Math.atan(1));
        System.out.println("π/2的角度值：" + Math.toDegrees(Math.PI/2));
        System.out.println(Math.PI);

        System.out.println("Math 的 floor,round 和 ceil 方法实例比较----------:");
        double[] nums = { 1.4, 1.5, 1.6, -1.4, -1.5, -1.6 };
        for (double num : nums) {
            System.out.println("Math.floor(" + num + ")=" + Math.floor(num));
            System.out.println("Math.round(" + num + ")=" + Math.round(num));
            System.out.println("Math.ceil(" + num + ")=" + Math.ceil(num));
        }

        Character ch = 'a';
        ch.toUpperCase(ch);
        System.out.println("ch = "+ch);
    }

    public static void StringTest() {
        char[] helloArray = { 'r', 'u', 'n', 'o', 'o', 'b'};
        String helloString = new String(helloArray);
        System.out.println( helloString );

        String site = "www.runoob.com";
        int len = site.length();
        System.out.println( "菜鸟教程网址长度 : " + len );

        String string1 = "菜鸟教程网址：";
        System.out.println("1、" + string1 + "www.runoob.com");

        StringBuilder sb = new StringBuilder(10);
        sb.append("Runoob..");
        System.out.println(sb);
        sb.append("!");
        System.out.println(sb);
        sb.insert(8, "Java");
        System.out.println(sb);
        sb.delete(5,8);
        System.out.println(sb);

        StringBuffer sBuffer = new StringBuffer("菜鸟教程官网：");
        sBuffer.append("www");
        sBuffer.append(".runoob");
        sBuffer.append(".com");
        System.out.println(sBuffer);
    }

    public static void ArrayTest() {
        // 数组大小
        int size = 10;
        // 定义数组
        double[] myList = new double[size];
        myList[0] = 5.6;
        myList[1] = 4.5;
        myList[2] = 3.3;
        myList[3] = 13.2;
        myList[4] = 4.0;
        myList[5] = 34.33;
        myList[6] = 34.0;
        myList[7] = 45.45;
        myList[8] = 99.993;
        myList[9] = 11123;
        // 计算所有元素的总和
        double total = 0;
        for (int i = 0; i < size; i++) {
            total += myList[i];
        }
        System.out.println("总和为： " + total);

        double[] myList1 = {1.9, 2.9, 3.4, 3.5};

        // 打印所有数组元素
        for (int i = 0; i < myList1.length; i++) {
            System.out.println(myList1[i] + " ");
        }
        // 计算所有元素的总和
        double total1 = 0;
        for (int i = 0; i < myList1.length; i++) {
            total1 += myList1[i];
        }
        System.out.println("Total1 is " + total1);
        // 查找最大元素
        double max = myList1[0];
        for (int i = 1; i < myList1.length; i++) {
            if (myList1[i] > max) max = myList1[i];
        }
        System.out.println("Max is " + max);

        double[] myList2 = {1.9, 2.9, 3.4, 3.5};

        // 打印所有数组元素
        for (double element: myList2) {
            System.out.println(element);
        }

        String str[][] = new String[3][4];
        String s[][] = new String[2][];
        s[0] = new String[2];
        s[1] = new String[3];
        s[0][0] = new String("Good");
        s[0][1] = new String("Luck");
        s[1][0] = new String("to");
        s[1][1] = new String("you");
        s[1][2] = new String("!");
    }

    public static void printArray(int[] array) {
        for (int i = 0; i < array.length; i++) {
            System.out.println(array[i] + " ");
        }
    }

    public static void DateTest() {
        Date date=new Date();
        //b的使用，月份简称
        String str=String.format(Locale.US,"英文月份简称：%tb",date);
        System.out.println(str);
        System.out.printf("本地月份简称：%tb%n",date);
        //B的使用，月份全称
        str=String.format(Locale.US,"英文月份全称：%tB",date);
        System.out.println(str);
        System.out.printf("本地月份全称：%tB%n",date);
        //a的使用，星期简称
        str=String.format(Locale.US,"英文星期的简称：%ta",date);
        System.out.println(str);
        //A的使用，星期全称
        System.out.printf("本地星期的简称：%tA%n",date);
        //C的使用，年前两位
        System.out.printf("年的前两位数字（不足两位前面补0）：%tC%n",date);
        //y的使用，年后两位
        System.out.printf("年的后两位数字（不足两位前面补0）：%ty%n",date);
        //j的使用，一年的天数
        System.out.printf("一年中的天数（即年的第几天）：%tj%n",date);
        //m的使用，月份
        System.out.printf("两位数字的月份（不足两位前面补0）：%tm%n",date);
        //d的使用，日（二位，不够补零）
        System.out.printf("两位数字的日（不足两位前面补0）：%td%n",date);
        //e的使用，日（一位不补零）
        System.out.printf("月份的日（前面不补0）：%te",date);

        try {
            System.out.println(new Date( ) + "\n");
            //Thread.sleep(1000*3);   // 休眠3秒
            System.out.println(new Date( ) + "\n");
        } catch (Exception e) {
            System.out.println("Got an exception!");
        }

        try {
            long start = System.currentTimeMillis( );
            System.out.println(new Date( ) + "\n");
            //Thread.sleep(5*60*10);
            System.out.println(new Date( ) + "\n");
            long end = System.currentTimeMillis( );
            long diff = end - start;
            System.out.println("Difference is : " + diff);
        } catch (Exception e) {
            System.out.println("Got an exception!");
        }

        Calendar c1 = Calendar.getInstance();
        // 获得年份
        int year = c1.get(Calendar.YEAR);
        // 获得月份
        int month = c1.get(Calendar.MONTH) + 1;
        // 获得日期
        int date1 = c1.get(Calendar.DATE);
        // 获得小时
        int hour = c1.get(Calendar.HOUR_OF_DAY);
        // 获得分钟
        int minute = c1.get(Calendar.MINUTE);
        // 获得秒
        int second = c1.get(Calendar.SECOND);
        // 获得星期几（注意（这个与Date类是不同的）：1代表星期日、2代表星期1、3代表星期二，以此类推）
        int day = c1.get(Calendar.DAY_OF_WEEK);

        String months[] = {
                "Jan", "Feb", "Mar", "Apr",
                "May", "Jun", "Jul", "Aug",
                "Sep", "Oct", "Nov", "Dec"};

        int year1;
        // 初始化 Gregorian 日历
        // 使用当前时间和日期
        // 默认为本地时间和时区
        GregorianCalendar gcalendar = new GregorianCalendar();
        // 显示当前时间和日期的信息
        System.out.print("Date: ");
        System.out.print(months[gcalendar.get(Calendar.MONTH)]);
        System.out.print(" " + gcalendar.get(Calendar.DATE) + " ");
        System.out.println(year1 = gcalendar.get(Calendar.YEAR));
        System.out.print("Time: ");
        System.out.print(gcalendar.get(Calendar.HOUR) + ":");
        System.out.print(gcalendar.get(Calendar.MINUTE) + ":");
        System.out.println(gcalendar.get(Calendar.SECOND));

        // 测试当前年份是否为闰年
        if(gcalendar.isLeapYear(year1)) {
            System.out.println("当前年份是闰年");
        }
        else {
            System.out.println("当前年份不是闰年");
        }
    }

    public static void RegexExample() {
        String content = "I am noob " +
                "from runoob.com.";

        String pattern = ".*runoob.*";

        boolean isMatch = Pattern.matches(pattern, content);
        System.out.println("字符串中是否包含了 'runoob' 子字符串? " + isMatch);

        // 按指定模式在字符串查找
        String line = "This order was placed for QT3000! OK?";
        String pattern1 = "(\\D*)(\\d+)(.*)";

        // 创建 Pattern 对象
        Pattern r = Pattern.compile(pattern1);

        // 现在创建 matcher 对象
        Matcher m = r.matcher(line);
        if (m.find( )) {
            System.out.println("Found value: " + m.group(0) );
            System.out.println("Found value: " + m.group(1) );
            System.out.println("Found value: " + m.group(2) );
            System.out.println("Found value: " + m.group(3) );
        } else {
            System.out.println("NO MATCH");
        }
    }

    public static int max(int num1, int num2) {
        int result;
        if (num1 > num2)
            result = num1;
        else
            result = num2;

        return result;
    }
}
