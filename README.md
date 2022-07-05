

# ch1

## 应用程序入口

1. package main 是必须的
2. 必须是main方法：`func main() {}`
3. 文件名不一定是main.go

## 入口参数

使用`os.Args`直接获取所用参数

执行 `go run hello_world.go hello_world` 

`os.Args` 返回的是个数组`[执行路径,hello_world]`
