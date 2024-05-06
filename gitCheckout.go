package main

type GitCheckout struct{}

func (g GitCheckout) GetDescription() string {
    return "Muda para um branch ou commit diferente"
}

