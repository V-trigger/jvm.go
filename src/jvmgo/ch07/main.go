package main
import (
    "fmt"
    "strings"
    "jvmgo/ch07/classpath"
    "jvmgo/ch07/rtda/heap"
)

func main()  {
    cmd := parseCmd()
    if cmd.versionFlag {
        fmt.Println("version 0.0.1")
    } else if cmd.helpFlag || cmd.class == "" {
        printUsage()
    } else {
        startJVM(cmd)
    }
}

func startJVM(cmd *Cmd) {
    cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
    //获取一个加载器
    classLoader := heap.NewClassLoader(cp)
    className := strings.Replace(cmd.class, ".", "/", -1)
    //加载主类
    mainClass := classLoader.LoadClass(className)
    //找到main方法
    mainMethod := mainClass.GetMainMethod()
    if mainMethod != nil {
        interpret(mainMethod)
    } else {
        fmt.Printf("Main method not found in class %s\n", cmd.class)
    }
}