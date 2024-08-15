package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 在一个新的UTS命名空间中执行一个shell命令。
// 它通过设置特定的系统进程属性来实现这一点，这些属性允许它创建一个与当前进程隔离的环境。
func main() {
	// 创建一个sh命令的执行对象。
	cmd := exec.Command("sh")

	// 设置进程属性，以在新的UTS命名空间中运行。
	// 这允许进程使用不同的主机名和域名。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	// 将标准输入、输出和错误流重定向到当前进程的相应流。
	// 这样可以在shell命令执行时直接访问和输出数据。
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行命令，并检查是否有错误发生。
	// 如果有错误，记录错误信息并退出程序。
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
