package main

type GitCommit struct{}

func (g GitCommit) GetDescription() string {
    return "Cria um novo commit com as mudanças rastreadas"
}

