package main

type GitStatus struct{}

func (g GitStatus) GetDescription() string {
    return "Exibe o estado do repositório git"
}
