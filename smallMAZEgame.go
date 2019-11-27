package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, desc string, nextNode *storyNode) {
	newChoice := &choice{cmd, desc, nextNode}
	node.choices = append(node.choices, newChoice)
}

func (node *storyNode) logCmds() {
	fmt.Println(node.text)

	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Printf("\t Command: %s  Description: %s \n", choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}

	fmt.Println("\tI couldn't find the matching command")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.logCmds()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	startRoom := storyNode{text: `
	You are in large room, You see 3 passages. You need the pick a direction!
	`}

	darkRoom := storyNode{text: `
	You are in dark room, You see a 2 shiny doors. You need a pick a door!
	`}

	girlsRoom := storyNode{text: `
	You are in room that has full of beautiful girls, You see that one coming and offering you a gift. Do you want to take and open it?
	`}

	nowayRoom := storyNode{text: `
	You are in room that has nothing to do and explore, You eventually get bored here and think to go back where you come from.
	`}

	leftDoorRoom := storyNode{text: `
	You are in room where ground gets sticky and pullingg you down. You have nothing to grasp and you eventually drown then die!
	`}

	rightDoorRoom := storyNode{text: `
	Congratz! you found the exit and came back to the daily normal stupid life where you belonged
	`}

	trapRoom := storyNode{text: `
	While not looking on your steps, you jumped into a room and trigged spikes on the floor as you entered. Spikes killed you!
	`}

	startRoom.addChoice("N", "Go north", &darkRoom)
	startRoom.addChoice("S", "Go south", &girlsRoom)
	startRoom.addChoice("W", "Go west", &trapRoom)
	startRoom.addChoice("E", "Go east", &nowayRoom)

	darkRoom.addChoice("L", "Pick left door", &leftDoorRoom)
	darkRoom.addChoice("R", "Pick right door", &rightDoorRoom)

	girlsRoom.addChoice("Y", "Take and open the gift", &trapRoom)
	girlsRoom.addChoice("O", "Take but don't open it and walk away slowly", &startRoom)
	girlsRoom.addChoice("N", "Don't take and walk away fast", &startRoom)

	nowayRoom.addChoice("B", "Go back to the door you used", &startRoom)

	startRoom.play()

	fmt.Printf("\n\nThe End\n")
}
