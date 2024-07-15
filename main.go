package main
import (
	"fmt"
	"os"
	"syscall"
	"os/exec"
)

/*
	run: calls main.go with arguments "child" 
		 similar to docker run <image>
*/
func run(){
	fmt.Println("Running", os.Args[2])

	//command to run same program by itself
	// this command will call same main.go(self) program with child as second args 
	// and other arguments
	// goprogram calling itself
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)	
	
	//running in a new namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err:= cmd.Run(); err!=nil {
		fmt.Println(err)
	}
}

/*
	Child: Called By main.go, creates a process with Alpine file System
			runs the alpine filesystem as a process in root	
*/
func child(){
	fmt.Println("Running in child", os.Args[2])

	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))
	must(syscall.Sethostname([]byte("container")))


	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

/*
 must: a function to handle errors
 @args: error
*/
func must(err error) {
	if err != nil {
		panic(err)
	}
}

/*
	main: Entry Point To GO Program
*/
func main(){
	switch os.Args[1] {
		case "run": 
			run()
		//same program called with this args
		case "child":
			child()
		default: 
			fmt.Println("Invalid Paramaters")
	}
}