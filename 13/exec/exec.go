package exec

import (
	"fmt"
	"os"
	"os/exec"
)

func Exct() {
	cmd := exec.Command("cmd.exe", "/k", "dir")
	//cmd := exec.Command("cmd.exe")
	cmd.Dir = `D:\Code`

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func startprocess() {
	//要启动的可执行文件路径
	exe := os.Getenv("ComSpec")
	if exe == "" {
		exe = `C:\Windows\System32\cmd.exe`
	}
	//exe = `D:\Tool\notepad++\notepad++.exe`

	// 工作目录、环境变量、标准输入输出等属性
	procAttr := &os.ProcAttr{
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   "D:/Code", //工作目录
	}
	//C：执行命令后退出
	//K：执行命令后保留窗口，不退出
	process, err := os.StartProcess(exe, []string{"cmd.exe", "/k", "dir"}, procAttr)
	//process, err := os.StartProcess(exe, []string{"notepad++.exe"}, procAttr)
	if err != nil {
		fmt.Println("启动失败", err)
		return
	}
	fmt.Println("进程PID", process.Pid)

	// if err := process.Release(); err != nil {
	// 	fmt.Println("释放进程资源失败：", err)
	// }
	state, err := process.Wait()
	if err != nil {
		fmt.Println("等待进程结束失败", err)
		return
	}
	fmt.Println("进程结束状态", state)
	fmt.Println("退出码", state.ExitCode())
}
