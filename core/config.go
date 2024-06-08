package core

import (
	"encoding/json"
	"os"
)

type FactSheet struct {
	RotorSheet map[string]map[string]int `json:"rotors"`
	Reflector  map[string]int            `json:"reflector"`
}

type Config struct {
	SelectedRotor []string          `json:"selected_rotors"`
	Slot1Pos      int               `json:"rotor_slot_1_init_pos"`
	Slot2Pos      int               `json:"rotor_slot_2_init_pos"`
	Slot3Pos      int               `json:"rotor_slot_3_init_pos"`
	PlugMap       map[string]string `json:"plug_map"`
}

func readConfig[format Config | FactSheet](path string) (format, error) {
	var cfg format
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func GetConfigAndRotorMap() (Config, FactSheet) {
	cfg, err := readConfig[Config]("./config.json")
	if err != nil {
		panic("error reading config " + err.Error())
	}
	rotorMap, err := readConfig[FactSheet]("./assets/rotor_sheet.json")
	if err != nil {
		panic("error reading config " + err.Error())
	}
	return cfg, rotorMap
}
