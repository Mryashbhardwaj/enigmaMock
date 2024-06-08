package models

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/eiannone/keyboard"
)

// Enigma Order
// -<- reflector <- rotor1 <- rotor2 <- rotor3 <- input
// |
// ->- reflector -> rotor1 -> rotor2 -> rotor3 -> output
type Enigma struct {
	rotor1 Rotor
	rotor2 Rotor
	rotor3 Rotor

	reflector Reflector
	plugBoard PlugBoard
}

func (e *Enigma) incrementRotors() {
	if e.rotor3.increment() {
		if e.rotor2.increment() {
			e.rotor1.increment()
		}
	}
}

func AtoI(char string) int {
	runes := []rune(char)
	return int(runes[0] - 'A')
}

func ItoA(input int) string {
	return string(rune(input + 'A'))
}

func (e *Enigma) Encrypt(input string) string {
	var buffer bytes.Buffer
	for _, char := range strings.Split(input, "") {
		if char == " " {
			buffer.WriteString(char)
			continue
		}
		buffer.WriteString(e.process(strings.ToUpper(char)))
	}
	return buffer.String()
}

func (e *Enigma) process(char string) string {
	char = e.plugBoard.Process(char)
	val := AtoI(char)
	val = e.rotor3.crunch(val, DirectionRightToLeft)
	val = e.rotor2.crunch(val, DirectionRightToLeft)
	val = e.rotor1.crunch(val, DirectionRightToLeft)
	val = e.reflector.process(val)
	val = e.rotor1.crunch(val, DirectionLeftToRight)
	val = e.rotor2.crunch(val, DirectionLeftToRight)
	val = e.rotor3.crunch(val, DirectionLeftToRight)
	e.incrementRotors()
	char = ItoA(val)
	return e.plugBoard.Process(char)
}

func (e *Enigma) Welcome() {
	fmt.Println("[INFO] The machine is initialized with the set config")
}

func (e *Enigma) Prompt() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press any key, press ESC to quit")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch key {
		case keyboard.KeyEsc:
			fmt.Println("Exiting...")
			return
		default:
			if unicode.IsLetter(char) {
				upperChar := string(unicode.ToUpper(char))
				output := e.process(upperChar)
				fmt.Printf("> %s -> %s\n", upperChar, output)
			}
		}
	}
}

func GetNewMachine(rotor1, rotor2, rotor3 Rotor, reflector Reflector, plugBoard PlugBoard) Enigma {
	return Enigma{
		rotor1:    rotor1,
		rotor2:    rotor2,
		rotor3:    rotor3,
		reflector: reflector,
		plugBoard: plugBoard,
	}
}
