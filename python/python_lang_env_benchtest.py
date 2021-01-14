import os,sys
sys.path.append(os.getcwd())
import time
from collections import deque

def print_helloWorld():
    print("Hello, World!")

def stringTest():
    word = '字符串'
    sentence = "这是一个句子。"
    paragraph = """这是一个段落，可以由多行组成"""

    str = 'Runoob'
    print("输出字符串",str)  # 输出字符串
    print("输出第一个到倒数第二个的所有字符",str[0:-1])  # 输出第一个到倒数第二个的所有字符
    print("输出字符串第一个字符",str[0])  # 输出字符串第一个字符
    print("输出从第三个开始到第五个的字符",str[2:5])  # 输出从第三个开始到第五个的字符
    print("输出从第三个开始后的所有字符",str[2:])  # 输出从第三个开始后的所有字符
    print("输出从第二个开始到第五个且每隔两个的字符",str[1:5:2])  # 输出从第二个开始到第五个且每隔两个的字符
    print("输出字符串两次",str * 2)  # 输出字符串两次
    print("连接字符串:",str + '你好')  # 连接字符串
    print('------------------------------')
    print('hello\nrunoob')  # 使用反斜杠(\)+n转义特殊字符
    print(r'hello\nrunoob')  # 在字符串前面添加一个 r，表示原始字符串，不会发生转义

def basic_dataType():
    counter = 100  # 整型变量
    miles = 1000.0  # 浮点型变量
    name = "runoob"  # 字符串
    print("counter = ",counter)
    print("miles = ",miles)
    print("name=",name)

    a, b, c, d = 20, 5.5, True, 4 + 3j  #数字类型，分别为int、float、bool、complex(复数)
    print(type(a), type(b), type(c), type(d))

    var = 1
    del var # del语句删除一些对象引用

    print("算术运算符--------:")
    print("加法：",5 + 4)  # 加法
    print("减法:",4.3 - 2)  # 减法
    print("乘法:",3 * 7)  # 乘法
    print("除法得到一个浮点数:",2 / 4)  # 除法，得到一个浮点数
    print("除法得到一个整数:",2 // 4)  # 除法，得到一个整数
    print("取余:", 17 % 3)  # 取余
    print("求乘方:",2 ** 5)  # 乘方

    print("比较运算符------:")
    a1 = 21
    b1 = 10
    c1 = 0
    if (a1 == b1):
        print("1 - a1 等于 b1")
    else:
        print("1 - a1 不等于 b1")

    if (a1 != b1):
        print("2 - a1 不等于 b1")
    else:
        print("2 - a1 等于 b1")

    if (a1 < b1):
        print("3 - a1 小于 b1")
    else:
        print("3 - a1 大于等于 b1")

    if (a1 > b1):
        print("4 - a1 大于 b1")
    else:
        print("4 - a1 小于等于 b1")
    # 修改变量 a1 和 b1 的值
    a1 = 5
    b1 = 20
    if (a1 <= b1):
        print("5 - a1 小于等于 b1")
    else:
        print("5 - a1 大于  b1")

    if (b1 >= a1):
        print("6 - b1 大于等于 a1")
    else:
        print("6 - b1 小于 a1")

    print("赋值运算符-----:")
    a2 = 21
    b2 = 10
    c2 = 0
    c2 = a2 + b2
    print("1 - c2 的值为：", c2)
    c2 += a2
    print("2 - c2 的值为：", c2)
    c2 *= a2
    print("3 - c2 的值为：", c2)
    c2 /= a2
    print("4 - c2 的值为：", c2)
    c2 = 2
    c2 %= a2
    print("5 - c2 的值为：", c2)
    c2 **= a2
    print("6 - c2 的值为：", c2)
    c2 //= a2
    print("7 - c2 的值为：", c2)

    print("位运算符------:")
    a3 = 60  # 60 = 0011 1100
    b3 = 13  # 13 = 0000 1101
    c3 = 0
    c3 = a3 & b3  # 12 = 0000 1100
    print("1 - c3 的值为：", c3)
    c3 = a3 | b3  # 61 = 0011 1101
    print("2 - c3 的值为：", c3)
    c3 = a3 ^ b3  # 49 = 0011 0001
    print("3 - c3 的值为：", c3)
    c3 = ~a3  # -61 = 1100 0011
    print("4 - c3 的值为：", c3)
    c3 = a3 << 2  # 240 = 1111 0000
    print("5 - c3 的值为：", c3)
    c3 = a3 >> 2  # 15 = 0000 1111
    print("6 - c3 的值为：", c3)

    print("逻辑运算符------:")
    a4 = 10
    b4 = 20
    if (a4 and b4):
        print("1 - 变量 a4 和 b4 都为 true")
    else:
        print("1 - 变量 a4 和 b4 有一个不为 true")

    if (a4 or b4):
        print("2 - 变量 a4 和 b4 都为 true，或其中一个变量为 true")
    else:
        print("2 - 变量 a4 和 b4 都不为 true")
    # 修改变量 a4 的值
    a4 = 0
    if (a4 and b4):
        print("3 - 变量 a4 和 b4 都为 true")
    else:
        print("3 - 变量 a4 和 b4 有一个不为 true")

    if (a4 or b4):
        print("4 - 变量 a4 和 b4 都为 true，或其中一个变量为 true")
    else:
        print("4 - 变量 a4 和 b4 都不为 true")

    if not (a4 and b4):
        print("5 - 变量 a4 和 b4 都为 false，或其中一个变量为 false")
    else:
        print("5 - 变量 a4 和 b4 都为 true")

    print("成员运算符------:")
    a5 = 10
    b5 = 20
    list = [1, 2, 3, 4, 5]

    if (a5 in list):
        print("1 - 变量 a5 在给定的列表中 list 中")
    else:
        print("1 - 变量 a5 不在给定的列表中 list 中")

    if (b5 not in list):
        print("2 - 变量 b5 不在给定的列表中 list 中")
    else:
        print("2 - 变量 b5 在给定的列表中 list 中")
    # 修改变量 a 的值
    a5 = 2
    if (a5 in list):
        print("3 - 变量 a5 在给定的列表中 list 中")
    else:
        print("3 - 变量 a5 不在给定的列表中 list 中")

    print("身份运算符-------:")
    a6 = 20
    b6 = 20

    if (a6 is b6):
        print("1 - a6 和 b6 有相同的标识")
    else:
        print("1 - a6 和 b6 没有相同的标识")

    if (id(a6) == id(b6)):
        print("2 - a6 和 b6 有相同的标识")
    else:
        print("2 - a6 和 b6 没有相同的标识")

    # 修改变量 b6 的值
    b6 = 30
    if (a6 is b6):
        print("3 - a6 和 b6 有相同的标识")
    else:
        print("3 - a6 和 b6 没有相同的标识")

    if (a6 is not b6):
        print("4 - a6 和 b6 没有相同的标识")
    else:
        print("4 - a6 和 b6 有相同的标识")

    print("运算符优先级-----:")
    a7 = 20
    b7 = 10
    c7 = 15
    d7 = 5
    e7 = 0

    e7 = (a7 + b7) * c7 / d7  # ( 30 * 15 ) / 5
    print("(a7 + b7) * c7 / d7 运算结果为：", e7)

    e7 = ((a7 + b7) * c7) / d7  # (30 * 15 ) / 5
    print("((a7 + b7) * c7) / d7 运算结果为：", e7)

    e7 = (a7 + b7) * (c7 / d7)  # (30) * (15/5)
    print("(a7 + b7) * (c7 / d7) 运算结果为：", e7)

    e7 = a7 + (b7 * c7) / d7  # 20 + (150/5)
    print("a7 + (b7 * c7) / d7 运算结果为：", e7)

    print("and 拥有更高优先级-------:")
    x = True
    y = False
    z = False

    if x or y and z:
        print("yes")
    else:
        print("no")

def ListTest():
    list = ['abcd', 786, 2.23, 'runoob', 70.2]
    tinylist = [123, 'runoob']

    print("输出完整列表:",list)  # 输出完整列表
    print("输出列表第一个元素",list[0])  # 输出列表第一个元素
    print("从第二个开始输出到第三个元素",list[1:3])  # 从第二个开始输出到第三个元素
    print("输出从第三个元素开始的所有元素",list[2:])  # 输出从第三个元素开始的所有元素
    print("输出两次列表",tinylist * 2)  # 输出两次列表
    print("连接列表:",list + tinylist)  # 连接列表

    a = [66.25, 333, 333, 1, 1234.5]
    print(a.count(333), a.count(66.25), a.count('x'))
    a.insert(2, -1)
    a.append(333)
    print("插入元素后的列表a为: ",a)
    a.index(333)
    a.remove(333)
    print("删除元素后的列表a为:",a)
    a.reverse()
    print("逆序后的列表a为:",a)
    a.sort()
    print("经过排序后的列表a为:",a)

    print("将列表当做堆栈使用:")
    stack = [3, 4, 5]
    stack.append(6)
    stack.append(7)
    print("进入两个元素后的栈stack为:",stack)
    stack.pop()
    print("栈顶元素出栈后，此时的栈stack为:",stack)
    stack.pop()
    stack.pop()
    print("再次经历过两次栈顶元素出栈，此时的栈stack为:",stack)

    print("将列表当作队列使用:")
    queue = deque(["Eric", "John", "Michael"])
    queue.append("Terry")  # Terry arrives
    queue.append("Graham")  # Graham arrives
    queue.popleft()  # The first to arrive now leaves
    queue.popleft()  # The second to arrive now leaves
    print("此时的队列queue为:",queue)

def reverseWords(input):
    # 通过空格将字符串分隔符，把各个单词分隔为列表
    inputWords = input.split(" ")
    # 翻转字符串
    # 假设列表 list = [1,2,3,4],
    # list[0]=1, list[1]=2 ，而 -1 表示最后一个元素 list[-1]=4 ( 与 list[3]=4 一样)
    # inputWords[-1::-1] 有三个参数
    # 第一个参数 -1 表示最后一个元素
    # 第二个参数为空，表示移动到列表末尾
    # 第三个参数为步长，-1 表示逆向
    inputWords = inputWords[-1::-1]
    # 重新组合字符串
    output = ' '.join(inputWords)
    return output

def tupleTest():
    tuple = ('abcd', 786, 2.23, 'runoob', 70.2)
    tinytuple = (123, 'runoob')

    print("输出完整元组:",tuple)  # 输出完整元组
    print("输出元组的第一个元素:",tuple[0])  # 输出元组的第一个元素
    print("输出从第二个元素开始到第三个元素:",tuple[1:3])  # 输出从第二个元素开始到第三个元素
    print("输出从第三个元素开始的所有元素:",tuple[2:])  # 输出从第三个元素开始的所有元素
    print("输出两次元组:",tinytuple * 2)  # 输出两次元组
    print("连接元组:",tuple + tinytuple)  # 连接元组

def setTest():
    sites = {'Google', 'Taobao', 'Runoob', 'Facebook', 'Zhihu', 'Baidu'}
    print("输出集合，重复的元素被自动去掉:",sites)  # 输出集合，重复的元素被自动去掉
    # 成员测试
    if 'Runoob' in sites:
        print('Runoob 在集合中')
    else:
        print('Runoob 不在集合中')

    # set可以进行集合运算
    a = set('abracadabra')
    b = set('alacazam')
    print("集合a:",a)
    print("a 和 b 的差集:",a - b)  # a 和 b 的差集
    print("a 和 b 的并集:",a | b)  # a 和 b 的并集
    print("a 和 b 的交集:",a & b)  # a 和 b 的交集
    print("a 和 b 中不同时存在的元素:",a ^ b)  # a 和 b 中不同时存在的元素

def dictionaryTest():
    dict = {}
    dict['one'] = "1 - 菜鸟教程"
    dict[2] = "2 - 菜鸟工具"

    tinydict = {'name': 'runoob', 'code': 1, 'site': 'www.runoob.com'}
    print("输出键为 'one' 的值:",dict['one'])  # 输出键为 'one' 的值
    print("输出键为 2 的值:",dict[2])  # 输出键为 2 的值
    print("输出完整的字典:",tinydict)  # 输出完整的字典
    print("输出所有键:",tinydict.keys())  # 输出所有键
    print("输出所有值:",tinydict.values())  # 输出所有值
    for k, v in tinydict.items():
        print(k, v)

def ifConditionTest():
    age = 25
    print("IF条件语句测试结果------:")
    if age <= 0:
        print("你是在逗我吧!")
    elif age == 1:
        print("相当于 14 岁的人。")
    elif age == 2:
        print("相当于 22 岁的人。")
    elif age > 2:
        human = 22 + (age - 2) * 5
        print("对应人类年龄: ", human)

def LoopTest():
    print("while循环测试-----:")
    count = 0
    while count < 5:
        print(count, " 小于 5")
        count = count + 1
    else:
        print(count, " 大于或等于 5")

    print("for循环测试------:")
    sites = ["Baidu", "Google", "Runoob", "Taobao"]
    for site in sites:
        if site == "Runoob":
            print("菜鸟教程!")
            break
        print("循环数据 " + site)
    else:
        print("没有循环数据!")
    print("完成循环!")

    print("range()函数来遍历数字序列-------:")
    a = ['Google', 'Baidu', 'Runoob', 'Taobao', 'QQ']
    for i in range(len(a)):
        print(i, a[i])

# 创建迭代器
class MyNumbers:
    def __iter__(self):
        self.a = 1
        return self

    def __next__(self):
        x = self.a
        self.a += 1
        return x

def fibonacci(n): # 生成器函数 - 斐波那契
    a, b, counter = 0, 1, 0
    while True:
        if (counter > n):
            return
        yield a
        a, b = b, a + b
        counter += 1

# 定义函数
def printme(str):
    # 打印任何传入的字符串
    print(str)
    return

#可写函数说明
def printinfo( name, age = 35 ):
   "打印任何传入的字符串"
   print ("名字: ", name)
   print ("年龄: ", age)
   return

#二分查找算法(binarySearch函数)
def binarySearch(arr, l, r, x):
    # 基本判断
    if r >= l:
        mid = int(l + (r - l) / 2)
        # 元素正好的中间位置
        if arr[mid] == x:
            return mid
        # 元素小于中间位置的元素，只需要再比较左边的元素
        elif arr[mid] > x:
            return binarySearch(arr, l, mid - 1, x)
        # 元素大于中间位置的元素，只需要再比较右边的元素
        else:
            return binarySearch(arr, mid + 1, r, x)
    else:
        # 不存在
        return -1

# 快速排序算法(partition函数和quickSort函数)
def partition(arr, low, high):
    i = (low - 1)  # 最小元素索引
    pivot = arr[high]

    for j in range(low, high):
        # 当前元素小于或等于 pivot
        if arr[j] <= pivot:
            i = i + 1
            arr[i], arr[j] = arr[j], arr[i]

    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return (i + 1)

def quickSort(arr, low, high):
    if low < high:
        pi = partition(arr, low, high)
        quickSort(arr, low, pi - 1)
        quickSort(arr, pi + 1, high)

# 归并排序算法(包括merge函数和mergeSort函数)
def merge(arr, l, m, r):
    n1 = m - l + 1
    n2 = r - m

    # 创建临时数组
    L = [0] * (n1)
    R = [0] * (n2)

    # 拷贝数据到临时数组 arrays L[] 和 R[]
    for i in range(0, n1):
        L[i] = arr[l + i]

    for j in range(0, n2):
        R[j] = arr[m + 1 + j]

    # 归并临时数组到 arr[l..r]
    i = 0  # 初始化第一个子数组的索引
    j = 0  # 初始化第二个子数组的索引
    k = l  # 初始归并子数组的索引

    while i < n1 and j < n2:
        if L[i] <= R[j]:
            arr[k] = L[i]
            i += 1
        else:
            arr[k] = R[j]
            j += 1
        k += 1

    # 拷贝 L[] 的保留元素
    while i < n1:
        arr[k] = L[i]
        i += 1
        k += 1

    # 拷贝 R[] 的保留元素
    while j < n2:
        arr[k] = R[j]
        j += 1
        k += 1

def mergeSort(arr, l, r):
    if l < r:
        m = int((l + (r - 1)) / 2)

        mergeSort(arr, l, m)
        mergeSort(arr, m + 1, r)
        merge(arr, l, m, r)

if __name__ == "__main__":
    start = time.time()
    for i in range(0,10000):
        print_helloWorld()
        stringTest()
        basic_dataType()
        ListTest()
        input = 'I like runoob'
        rw = reverseWords(input)
        print("翻转后的字符串为:",rw)
        tupleTest()
        setTest()
        dictionaryTest()
        ifConditionTest()
        LoopTest()
        print("创建迭代器-------:")
        myclass = MyNumbers()
        myiter = iter(myclass)
        print(next(myiter))
        print(next(myiter))
        print("生成器-------:")
        f = fibonacci(10)  # f 是一个迭代器，由生成器返回生成
        printme("我要调用用户自定义函数!")
        # 调用printinfo函数
        printinfo(age=50, name="runoob")
        print("------------------------")
        printinfo(name="runoob")
        # 测试数组
        print("二分查找算法测试-----:")
        arr = [2, 3, 4, 10, 40]
        x = 10
        # 函数调用
        result = binarySearch(arr, 0, len(arr) - 1, x)
        if result != -1:
            print("元素在数组中的索引为 %d" % result)
        else:
            print("元素不在数组中")

        print("快速排序算法测试------:")
        arr = [10, 7, 8, 9, 1, 5]
        n = len(arr)
        quickSort(arr, 0, n - 1)
        print("排序后的数组:")
        for i in range(n):
            print("%d" % arr[i])

        print("归并排序算法测试------:")
        arr = [12, 11, 13, 5, 6, 7]
        n = len(arr)
        print("给定的数组")
        for i in range(n):
            print("%d" % arr[i]),
        mergeSort(arr, 0, n - 1)
        print("排序后的数组")
        for i in range(n):
            print("%d" % arr[i])
    end = time.time()
    print(end - start)  # 程序开始到程序结束的运行时间
