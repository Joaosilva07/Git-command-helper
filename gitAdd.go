package main

type GitAdd struct{}

func (g GitAdd) GetDescription() string {
    return "Adiciona um arquivo ao repositório git"
}
