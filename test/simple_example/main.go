package main

import (
	"fmt"

	chain "github.com/djunigari/golang-chain"
)

// Defina o tipo para o contexto extra (se necessário)
type MyExtraContext struct {
	Message string
}

func main() {
	// Crie algumas ações
	action1 := chain.NewAction[MyExtraContext]("Action1").Function(func(ctx *chain.Context[MyExtraContext]) {
		fmt.Println("Executing Action 1")
		fmt.Println("Message from extra context:", ctx.Extra.Message)
	})

	action2 := chain.NewAction[MyExtraContext]("Action2").Function(func(ctx *chain.Context[MyExtraContext]) {
		fmt.Println("Executing Action 2")
	})

	// Adicione as ações em uma lista
	actions := chain.Actions[MyExtraContext]{action1, action2}

	// Crie um Processor com as ações e habilite o log
	processor := chain.New(&actions, true)

	// Execute o processor com o contexto extra
	extraContext := MyExtraContext{
		Message: "Hello from extra context!",
	}

	processor.Run(&extraContext)
}
