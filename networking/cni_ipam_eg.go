/*
Program to clone and build CNI & CNI plugins projects, install go
compiler and run an example to demonstrate the use of IPAM host-local
module for ip address management
Author: Arun Sriraman <reachme[at]arunsriraman[dot]com>

Copyright (c) 2017 Arun Sriraman.

This program is free software: you can redistribute it and/or modify  
it under the terms of the GNU General Public License as published by  
the Free Software Foundation, version 3.

This program is distributed in the hope that it will be useful, but 
WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU 
General Public License for more details.

You should have received a copy of the GNU General Public License 
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
    "fmt"
    "bufio"
    "os/exec"
    "os"
    "sync"
)

var reader = bufio.NewReader(os.Stdin)
var wg sync.WaitGroup

func exec_cmd(cmd string, help_text string) {
    defer wg.Done()
    fmt.Println(help_text)
    fmt.Println("Executing:", cmd)
    out,_ := exec.Command("/bin/bash", "-c", cmd).Output()
    // All the printing is thread unsafe (fmt.print uses write() which
    // does not lock the buffer :( )
    fmt.Print(string(out))
}

func wait_for_user_input() {
    for true {
        inp,_ := reader.ReadString('\n')
        if (inp == "\n") {
            break
        }
    }
}

func install_prereqs(){
    // Parallely execute setup
    wg.Add(3)
    go exec_cmd(
        "git clone https://github.com/containernetworking/cni.git",
        "Cloning CNI repository",
    )
    go exec_cmd(
        "git clone https://github.com/containernetworking/plugins.git",
        "Cloning CNI plugins repository",
        )
    /* Ideally go should already be installed
    go_installer := "sudo add-apt-repository ppa:longsleep/golang-backports && sudo apt-get update && sudo apt-get install golang-go"
    go exec_cmd(
        go_installer,
        "Installing go compiler",
        )
    */
    go exec_cmd("go version", "Check if go installed")
    wg.Wait()
    // Build binaries for the CNI tool and the plugins
    wg.Add(2)
    go exec_cmd("cd cni && ./build.sh", "Build CNI tool")
    go exec_cmd("cd plugins && ./build.sh", "Build CNI plugin binaries")
    wg.Wait()
}

func run_demo_cni_ipam() {
    exec_cmd("sudo ip netns add cni_ipam_eg",
        "Add example namespace cni_ipam_eg",
        )
    exec_cmd("sudo ip netns", "Check if namespace is created")
}

func main() {
    fmt.Println("CNI with IPAM driver tutorial")

    install_prereqs()
    run_demo_cni_ipam()


    fmt.Println("Thank you!")
}
