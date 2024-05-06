package main

type Command interface {
    GetDescription() string
}

type SimpleCommand struct {
    Description string
}

func (c SimpleCommand) GetDescription() string {
    return c.Description
}

func InitializeCommands() map[string]Command {
    return map[string]Command{
        "add":   GitAdd{},
        "commit": GitCommit{},
        "push":   GitPush{},
        "pull":   GitPull{},
        "clone":  GitClone{},
        "checkout": GitCheckout{},
    }
}