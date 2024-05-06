package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func init() {
   GitCommands = InitializeCommands()
}

var GitCommands map[string]Command


func main() {

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Insira um comando git ou 'help' para ver todos os comandos: ")
    command, _ := reader.ReadString('\n')
    command = strings.TrimSpace(command)

    if command == "help" {
        for cmd, desc := range GitCommands {
            fmt.Printf("%s: %s\n", cmd, desc.GetDescription())
        }
    } else if cmd, ok := GitCommands[command]; ok {
        fmt.Println(cmd.GetDescription())
    } else {
        fmt.Println("Comando git não encontrado")
    }
}