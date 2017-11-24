package defaults

import (
	"gomine/commands"
	"gomine/interfaces"
	"gomine/commands/arguments"
	"fmt"
	"gomine/vectorMath"
)

type TestCommand struct {
	*commands.Command
	server interfaces.IServer
}

func NewTest(server interfaces.IServer) TestCommand {
	var test = TestCommand{commands.NewCommand("test", "Tests the command parser", "gomine.stop", []string{"test"}), server}

	var intArg = arguments.NewFloatArg("test", false)
	intArg.SetInputAmount(1)
	test.AppendArgument(intArg)

	var stringArg = arguments.NewStringArg("anotherTest", true)
	stringArg.SetInputAmount(2)
	test.AppendArgument(stringArg)

	test.AppendArgument(arguments.NewStringEnum("testEnum", true, []string{"option", "test_option", "test"}))
	return test
}

func (command TestCommand) Execute(sender interfaces.ICommandSender, arguments []interfaces.ICommandArgument) bool {
	fmt.Println(arguments[0].GetOutput())
	fmt.Println(arguments[1].GetOutput())
	fmt.Println(arguments[2].GetOutput())

	vectorMath.NewTripleVector(arguments[0].GetOutput().(float32), arguments[0].GetOutput().(float32), arguments[0].GetOutput())
	return true
}