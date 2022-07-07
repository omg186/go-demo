# 学习

## 应用程序入口

1. package main 是必须的
2. 必须是main方法：`func main() {}`
3. 文件名不一定是main.go

## 入口参数

使用`os.Args`直接获取所用参数

执行 `go run hello_world.go hello_world`

`os.Args` 返回的是个数组`[执行路径,hello_world]`

## panic

1. panic用于不可以恢复的错误
2. panic退出钱会执行defer指定的内容

## os.Exit

1. 退出时不会掉用defer指定的函数
2. 退出时不输出当前调用栈信息

## recover

> 当心形成僵尸服务进程，导致health check 失效
> 往往是我们恢复不确定性错误的最好方法

## package

1. 基本复用模块单元

    1. 以首字母来表明可被包外访问
2. 代码的package 可以和所在的目录不一致
3. 同一目录里的Go代码的package要保持一致


## 协层

> Groutine 是M:N的
> 
> 如果协层发生堵塞，go在启动的时候会有一个守护线程会做计数，每个Processor完成的协层数量，如果一段时间没有发生变化的时候，它就插入一个标记，协层运行非内联函数的时候会发现这个标记，就会切换到别的协层继续运行
>> Processor如果发生等待了，Processor会自动将自己移动到别的协层
