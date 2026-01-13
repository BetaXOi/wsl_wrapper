package main

import (
    "flag"
    "fmt"
    "net/url"
    "os/exec"
)

// "C:\Program Files (x86)\NetSarang\Xshell 8\Xshell.exe" -url ssh://aifort:2003014114539995136@21.40.83.21:50300 -newtab 10.16.63.5:22

func main() {
    // 定义变量，用于接收命令行参数值
    var urlStr string
    var newTab string

    // 注册命令行参数
    flag.StringVar(&urlStr, "url", "", "SSH URL，格式为 ssh://user:password@host:port")
    flag.StringVar(&newTab, "newtab", "", "新标签页的地址，格式为 host:port")

    // 解析命令行参数
    flag.Parse()

    // 解析 SSH URL
    if urlStr != "" {
        sshURL, err := url.Parse(urlStr)
        if err == nil {
            fmt.Printf("SSH URL: %s\n", urlStr)
            fmt.Printf("SSH 协议: %s\n", sshURL.Scheme)
            fmt.Printf("SSH 用户名: %s\n", sshURL.User.Username())
            password, _ := sshURL.User.Password()
            fmt.Printf("SSH 密码: %s\n", password)
            fmt.Printf("SSH 主机: %s\n", sshURL.Hostname())
            fmt.Printf("SSH 端口: %s\n", sshURL.Port())
        } else {
            fmt.Printf("解析 SSH URL 失败: %v\n", err)
        }
    }

    bin := `C:\Program Files\WindowsApps\Microsoft.WindowsTerminal_1.23.12811.0_x64__8wekyb3d8bbwe\wt.exe`
    cmd := exec.Command(bin)
    script := "/home/nb/wsl_wrapper.sh " + urlStr
    cmd.Args = append(cmd.Args, "-w", "0", "nt", "--title", newTab, "bash", "-c", script)
    fmt.Println(cmd.Args)
    cmd.Start()
}
