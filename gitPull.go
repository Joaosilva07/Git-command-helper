package main

type GitPull struct{}

func (g GitPull) GetDescription() string {
    return "Busca e integra com outra referência ou repositório"
}

