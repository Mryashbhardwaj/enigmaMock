package main

import (
	"flag"
	"fmt"

	"main/core"
	"main/core/models"
	"main/core/utils"
)

func initRotors(cfg core.Config, factSheet core.FactSheet) (models.Rotor, models.Rotor, models.Rotor) {
	rotorSlot1, err := models.GetNewRotor(
		factSheet.RotorSheet[cfg.SelectedRotor[0]],
		cfg.Slot1Pos,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	rotorSlot2, err := models.GetNewRotor(
		factSheet.RotorSheet[cfg.SelectedRotor[1]],
		cfg.Slot2Pos,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	rotorSlot3, err := models.GetNewRotor(
		factSheet.RotorSheet[cfg.SelectedRotor[2]],
		cfg.Slot3Pos,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rotorSlot1, rotorSlot2, rotorSlot3
}

func initialiseEnigmaMachine(cfg core.Config, factSheet core.FactSheet) models.Enigma {
	rot1, rot2, rot3 := initRotors(cfg, factSheet)
	plugBoard := models.GetNewPegboard(cfg.PlugMap)
	reflector := models.GetNewReflector(factSheet.Reflector)
	return models.GetNewMachine(rot1, rot2, rot3, reflector, plugBoard)
}

func main() {
	file := flag.String("file", "", "--file=./input.txt")
	flag.Parse()

	cfg, factSheet := core.GetConfigAndRotorMap()
	if !utils.VerifyConfig(cfg) {
		fmt.Println("> Exiting as the init config needs to be assessed again")
		return
	}
	machine := initialiseEnigmaMachine(cfg, factSheet)
	machine.Welcome()

	if *file != "" {
		//file mode
		fileContent, err := utils.ReadFile(*file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(machine.Encrypt(fileContent))
	} else {
		machine.Prompt()
	}
}
