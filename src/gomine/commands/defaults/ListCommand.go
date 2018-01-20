package defaults

import (
	"gomine/commands"
	"gomine/interfaces"
	"gomine/utils"
	"strconv"
)

type ListCommand struct {
	*commands.Command
	server interfaces.IServer
}

func NewList(server interfaces.IServer) ListCommand {
	var list = ListCommand{commands.NewCommand("list", "Lists all players online", "gomine.list", []string{}), server}
	list.ExemptFromPermissionCheck(true)
	return list
}

func (list ListCommand) Execute(sender interfaces.ICommandSender) {
	var playerList = utils.BrightGreen + "-----" + utils.White + " Player List (" + strconv.Itoa(len(list.server.GetPlayerFactory().GetPlayers())) + " Players) " + utils.BrightGreen + "-----\n"
	for name, player := range list.server.GetPlayerFactory().GetPlayers() {
		playerList += utils.BrightGreen + name + ": " + utils.Yellow + utils.Bold + strconv.Itoa(int(player.GetPing())) + "ms\n"
	}
	sender.SendMessage(playerList)
}