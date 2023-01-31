package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	//"strings"
	"sync"

	"github.com/hpifu/go-kit/hflag"
)

// "C:\Program Files\PuTTY\kitty.exe" -t -ssh -load bastion_2018 -pw ****** ningbo@127.0.0.1 -P 53377
// "C:\Program Files (x86)\SSH Communications Security\SSH Secure Shell\SshClient.exe" -h localhost -p 39464 -u ningbo ningbo@10．246．157．248(SSH).ssh2

func main() {
	// putty
	//hflag.AddFlag("tty", "request tty", hflag.Shorthand("t"), hflag.Type("bool"))
	//hflag.AddFlag("ssh", "protocol", hflag.Type("bool"))
	//hflag.AddFlag("load", "load session")
	//hflag.AddFlag("pw", "password")
	//hflag.AddFlag("port", "ssh port", hflag.Shorthand("P"))
	//hflag.AddPosFlag("pos", "site")

	// ssh secure shell
	hflag.AddFlag("host", "ssh remote host", hflag.Shorthand("h"))
	hflag.AddFlag("port", "ssh port", hflag.Shorthand("p"))
	hflag.AddFlag("user", "username", hflag.Shorthand("u"))
	hflag.AddPosFlag("pos", "site")


	hflag.Parse()

	//site := strings.Split(hflag.GetString("pos"), "@")
	//user := site[0]
	user := hflag.GetString("user")
	remotePort := hflag.GetString("port")

	incoming, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}

	_, port, _ := net.SplitHostPort(incoming.Addr().String())
	fmt.Printf("Listening on port 0.0.0.0:%s\n", port)

	var first bool = true
	serving := make(chan int)
	client := sync.WaitGroup{}

	go func() {
		for {
			conn, err := incoming.Accept()
			if err != nil {
				log.Fatal("could not accept client connection", err)
			}

			target, err := net.Dial("tcp", "127.0.0.1:"+remotePort)
			if err != nil {
				log.Fatal("could not connect to target", err)
			}
			fmt.Printf("connection to server %v established!\n", target.RemoteAddr())

			fmt.Printf("client '%v' connected!\n", conn.RemoteAddr())

			client.Add(1)
			if first {
				serving <- 1
				first = false
			}

			fmt.Printf("%s <--> %s\n", port, hflag.GetString("port"))
			go tunnelIO(&client, conn, target)

		}
	}()

	bin := "cmd.exe"
	cmd := exec.Command(bin)
	script := "/home/nb/wsl_wrapper.sh " + port + " " + user
	cmd.Args = append(cmd.Args, "/C", "start", "wt", "-w", "0", "nt", "bash", "-c", script)
	//bin := `C:\Program Files\WindowsApps\Microsoft.WindowsTerminal_1.15.3466.0_x64__8wekyb3d8bbwe\wt.exe`
	//cmd := exec.Command(bin)
	//script := "/home/nb/wsl_wrapper.sh " + port + " " + user
	//cmd.Args = append(cmd.Args, "-w", "0", "nt", "bash", "-c", script)
	fmt.Println(cmd.Args)
	cmd.Start()

	fmt.Println("waiting for first incoming connection...")
	<-serving

	fmt.Println("waiting for client disconnecting...")
	client.Wait()
}

func tunnelIO(client *sync.WaitGroup, src, dst net.Conn) {
	defer func() {
		fmt.Println("closing tunning...")
		src.Close()
		dst.Close()
		client.Done()
	}()

	fmt.Println("tunning...")
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func(){
		io.Copy(src, dst)
		wg.Done()
	}()
	go func(){
		io.Copy(dst, src)
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("tunning finished")
}
