package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	ROMAN_ASING = "ROMAN_ASSING"
	VALUE_ASING = "VALUE_ASING"
	QUESTION    = "QUESTION"
	EXIT        = "EXIT"
	UNKNOW      = "UNKNOW"
)

type UniversalLenguage struct {
	Word  string
	Value int
}

var num = map[string]UniversalLenguage{
	"I": {"", 1},
	"V": {"", 5},
	"X": {"", 10},
	"L": {"", 50},
	"C": {"", 100},
	"D": {"", 500},
	"M": {"", 1000},
}

var metals = make(map[string]int)

type Roman struct{}

func NewRoman() *Roman {
	return &Roman{}
}

func (r *Roman) ToNumber(n string) int {
	out := 0
	splitRomand := strings.Fields(n)
	ln := len(splitRomand)
	for i := 0; i < ln; i++ {
		c := string(splitRomand[i])
		vc := num[c].Value
		if i < ln-1 {
			cnext := string(splitRomand[i+1])
			vcnext := num[cnext].Value
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}

func main() {
	recursiveApp()
}

func recursiveApp() {
	validateInput(getInput())
	recursiveApp()
}

func validateInput(inputString string) {
	typeInput := getTypeInput(inputString)
	fmt.Println(typeInput)
	fmt.Printf("%v, %v", num, metals)
	switch typeInput {
	case ROMAN_ASING:
		romanAssingProcess(inputString)
	case QUESTION:
		fmt.Println(questionProcess(inputString))
	case VALUE_ASING:
		valueAssingProcess(inputString)
	case EXIT:
		os.Exit(3)
	default:
		fmt.Println("I have no idea what you are talking about")
	}

}

func getTypeInput(input string) string {
	splitWords := strings.Fields(input)
	if len(splitWords) == 3 {
		return ROMAN_ASING
	}
	if strings.HasSuffix(input, "?") {
		return QUESTION
	}

	if strings.HasSuffix(input, "Credits") {
		return VALUE_ASING
	}
	if strings.EqualFold(input, "exit") {
		return EXIT
	}
	return UNKNOW
}

func romanAssingProcess(romanAssingInput string) {
	splitInput := strings.Fields(romanAssingInput)
	v, ok := num[splitInput[2]]
	if ok {
		v.Word = splitInput[0]
		num[splitInput[2]] = v
	} else {
		fmt.Printf("can not assing %s to %s  \n", splitInput[0], splitInput[2])
	}
}

func getGalacticalNumbersToRoman(galacticalString string) string {
	galacticalSplit := strings.Fields(galacticalString)
	var romanNumbers []string
	for i := 0; i < len(galacticalSplit); i++ {
		lastValue := getRomanNumber(galacticalSplit[i])
		if lastValue != "" {
			romanNumbers = append(romanNumbers, lastValue)
		}
	}
	return strings.Join(romanNumbers[:], " ")
}

func getMetal(galacticalString string) string {
	galacticalSplit := strings.Fields(galacticalString)
	for i := 0; i < len(galacticalSplit); i++ {
		runeChar := rune(galacticalSplit[i][0])
		if unicode.IsUpper(runeChar) {
			return galacticalSplit[i]
		}
	}
	return ""
}

func getNumberAssigment(galacticalString string) int {
	galacticalSplit := strings.Fields(galacticalString)
	numericalAssing := 0
	for i := 0; i < len(galacticalSplit); i++ {
		if v, err := strconv.Atoi(galacticalSplit[i]); err == nil {
			numericalAssing = v
		}
	}
	return numericalAssing
}

func getRomanNumber(galacticalWord string) string {
	for k, v := range num {
		if v.Word == galacticalWord {
			return k
		}
	}
	return ""
}

func questionProcess(questionInput string) string {
	var strs []string
	splitSentence := strings.Split(questionInput, "is")
	stringRoman := getGalacticalNumbersToRoman(splitSentence[1])
	metal := getMetal(splitSentence[1])
	coin := getCoin(questionInput)
	total := NewRoman().ToNumber(stringRoman) * metals[metal]
	for _, romanNumber := range strings.Fields(stringRoman) {
		strs = append(strs, num[romanNumber].Word)
	}
	if metal != "" {
		strs = append(strs, metal)
	}
	strs = append(strs, "is", strconv.Itoa(total))

	if coin != "" {
		strs = append(strs, coin)
	}
	return strings.Join(strs, " ")
}

func getCoin(questionInput string) string {
	galacticalSplit := strings.Fields(questionInput)
	for ind, word := range galacticalSplit {
		if word == "is" {
			return galacticalSplit[ind-1]
		}
	}
	return ""
}

func valueAssingProcess(valueAssingInput string) {
	stringRoman := getGalacticalNumbersToRoman(valueAssingInput)
	valueMetal := getNumberAssigment(valueAssingInput)
	valuefromRoman := NewRoman().ToNumber(stringRoman)
	valueProductMetal := valueMetal / valuefromRoman
	metal := getMetal(valueAssingInput)
	metals[metal] = valueProductMetal
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)
	return strings.TrimSuffix(text, "\n")
}
