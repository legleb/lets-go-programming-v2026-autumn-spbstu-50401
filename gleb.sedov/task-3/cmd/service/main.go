package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value string `xml:"Value"`
}

type Currency struct {
	NumCode int `json:"num_code"`
	CharCode string `json:"char_code"`
	Value float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to the YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("the mandatory -config flag was not provided")
	}

	yamlData, err := os.ReadFile(*configPath)

	if err != nil {
		panic(fmt.Sprintf("error reading config: %v", err))
	}

	var cfg Config

	if err := yaml.Unmarshal(yamlData, &cfg); err != nil {
		panic(fmt.Sprintf("YAML decoding error: %v", err))
	}

	xmlData, err := os.ReadFile(cfg.InputFile)

	if err != nil {
		panic(fmt.Sprintf("XML reading error: %v", err))
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charsetReader
	var valCurs ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Sprintf("XML decoding error: %v", err))
	}

	currencies := make([]Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		codeInt := 0

		if s := strings.TrimSpace(valute.NumCode); s != "" {
			codeInt, _ = strconv.Atoi(s)
		}

		valStr := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")
		valFloat, err := strconv.ParseFloat(valStr, 64)

		if err != nil {
			panic(fmt.Sprintf("incorrect value %q: %v", valute.Value, err))
		}

		currencies = append(currencies, Currency{
			NumCode: codeInt,
			CharCode: strings.TrimSpace(valute.CharCode),
			Value: valFloat,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	jsonData, err := json.MarshalIndent(currencies, "", "    ")

	if err != nil {
		panic(fmt.Sprintf("JSON serialization error: %v", err))
	}

	dir := filepath.Dir(cfg.OutputFile)

	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("error creating directory: %v", err))
	}

	if err := os.WriteFile(cfg.OutputFile, jsonData, 0600); err != nil {
		panic(fmt.Sprintf("file write error: %v", err))
	}
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	if strings.ToLower(charset) == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}
