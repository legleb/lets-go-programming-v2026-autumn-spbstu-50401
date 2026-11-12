package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
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

	xmlData, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("ошибка чтения XML: %v", err))
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(charset) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Sprintf("ошибка декодирования XML: %v", err))
	}

	fmt.Printf("Успешно прочитано валют: %d\n", len(valCurs.Valutes))
}
