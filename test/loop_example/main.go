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
	// Crie uma ação para inicializar uma lista e armazenar no contexto
	createListAction := chain.NewAction[MyExtraContext]("CreateListAction").Function(func(ctx *chain.Context[MyExtraContext]) {
		fmt.Println("Creating list and storing in context")
		list := []string{"item1", "item2", "item3", "item4"}
		ctx.Additional["MyList"] = list
	})

	getNextElement := chain.NewAction[MyExtraContext]("PrintItemAction").Function(func(ctx *chain.Context[MyExtraContext]) {
		myList, ok := ctx.Additional["MyList"].([]string)
		if !ok {
			ctx.SetErr(chain.ErrInvalidVariableType)
			return
		}

		index, ok := ctx.Additional["MyList_index"].(int)
		if !ok {
			ctx.SetErr(chain.ErrInvalidVariableType)
			return
		}

		if index+1 < len(myList) {
			ctx.LoopAction["LoopPrintList"] = true
			ctx.Additional["MyList_index"] = index + 1
			ctx.Additional["item"] = myList[index+1]
			return
		}
		ctx.LoopAction["LoopPrintList"] = false
	})

	// Crie uma ação para imprimir a lista armazenada no contexto
	printItemAction := chain.NewAction[MyExtraContext]("PrintItemAction").Function(func(ctx *chain.Context[MyExtraContext]) {
		if item, ok := ctx.Additional["item"].(string); ok {
			fmt.Println(item)
		} else {
			fmt.Println("List not found in context or has invalid type")
		}
	})

	// // Crie uma ação para imprimir a lista armazenada no contexto
	// endLoop := chain.NewAction[MyExtraContext]("endLoop").Function(func(ctx *chain.Context[MyExtraContext]) {
	// 	len, ok := ctx.Additional["Len(MyList)"].(int)
	// 	if !ok {
	// 		ctx.SetErr(chain.ErrInvalidVariableType)
	// 		return
	// 	}

	// 	index, ok := ctx.Additional["MyList_index"].(int)
	// 	if !ok {
	// 		ctx.SetErr(chain.ErrInvalidVariableType)
	// 		return
	// 	}

	// 	if index < len {
	// 		ctx.LoopAction["LoopPrintList"] = true
	// 	} else {
	// 		ctx.LoopAction["LoopPrintList"] = false
	// 	}
	// })

	printListAction := chain.NewAction[MyExtraContext]("LoopPrintList").
		Type(chain.LoopAction).
		Loop(printItemAction, getNextElement).
		Function(func(ctx *chain.Context[MyExtraContext]) {
			myList, ok := ctx.Additional["MyList"].([]string)
			if !ok {
				ctx.SetErr(chain.ErrInvalidVariableType)
				return
			}

			ctx.Additional["MyList_index"] = 0
			ctx.Additional["item"] = myList[0]
		})

	// Adicione as ações em uma lista
	actions := chain.Actions[MyExtraContext]{createListAction, printListAction}

	// Crie um Processor com as ações e habilite o log
	processor := chain.New("loop_example_test", &actions, false)

	// Execute o processor com o contexto extra
	extraContext := MyExtraContext{
		Message: "Hello from extra context!",
	}

	processor.Run(&extraContext)
}
