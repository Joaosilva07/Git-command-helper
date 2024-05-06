package main

type GitClone struct{}

func (g GitClone) GetDescription() string {
    return "Clona um repositório remoto"
}

