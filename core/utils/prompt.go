package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/eiannone/keyboard"

	"main/core"
)

func VerifyConfig(config core.Config) bool {
	fmt.Printf("Please review the config\n")
	fmt.Println("Slot 1 :")
	fmt.Printf("\t Rotor Selected : %s\n", config.SelectedRotor[0])
	fmt.Printf("\t Rotor Position : %d\n", config.Slot1Pos)

	fmt.Println("Slot 2 :")
	fmt.Printf("\t Rotor Selected : %s\n", config.SelectedRotor[1])
	fmt.Printf("\t Rotor Position : %d\n", config.Slot2Pos)

	fmt.Println("Slot 3 :")
	fmt.Printf("\t Rotor Selected : %s\n", config.SelectedRotor[2])
	fmt.Printf("\t Rotor Position : %d\n", config.Slot3Pos)

	fmt.Println("Plug Map Parings:")
	for key, val := range config.PlugMap {
		fmt.Printf("\t %s <-> %s\n", strings.ToUpper(key), strings.ToUpper(val))
	}
	fmt.Println("\n> Press Return/Enter to accept.")
	fmt.Println("> Press any other key should you wish to halt.")

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()
	_, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}
	if key == keyboard.KeyEnter {
		return true
	}
	return false

}
