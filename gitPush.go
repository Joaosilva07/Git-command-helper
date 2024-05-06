package main

type GitPush struct{}

func (g GitPush) GetDescription() string {
    return "Envia as alterações para o repositório remoto"
}

