package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

	currencies := make([]Currency, 0, len(valCurs.Valutes))

	for _, v := range valCurs.Valutes {
		codeInt := 0
		if s := strings.TrimSpace(v.NumCode); s != "" {
			codeInt, _ = strconv.Atoi(s)
		}

		valStr := strings.ReplaceAll(strings.TrimSpace(v.Value), ",", ".")
		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			panic(fmt.Sprintf("некорректное значение %q: %v", v.Value, err))
		}

		currencies = append(currencies, Currency{
			NumCode:  codeInt,
			CharCode: strings.TrimSpace(v.CharCode),
			Value:    valFloat,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	fmt.Printf("Отсортировано: %d. Первый: %s (%.4f)\n", len(currencies), currencies[0].CharCode, currencies[0].Value)
}
