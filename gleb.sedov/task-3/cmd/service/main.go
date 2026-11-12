package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "", "путь до YAML-конфига")
	flag.Parse()

	if *configPath == "" {
		panic("не передан обязательный флаг -config")
	}

	yamlData, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Sprintf("ошибка чтения конфига: %v", err))
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlData, &cfg); err != nil {
		panic(fmt.Sprintf("ошибка декодирования YAML: %v", err))
	}

	fmt.Printf("Input: %s, Output: %s\n", cfg.InputFile, cfg.OutputFile)
}
