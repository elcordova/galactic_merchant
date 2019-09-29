package main

import (
	"testing"
)

func TestConvertRomanNumeralTo1(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("I")

	if n != 1 {
		t.Errorf("I should return 1, but %v", n)
	}
}

func TestConvertRomanNumeralTo2(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("I I")

	if n != 2 {
		t.Errorf("II should return 2, but %v", n)
	}
}

func TestConvertRomanNumeralTo3(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("I I I")

	if n != 3 {
		t.Errorf("III should return 3, but %v", n)
	}
}

func TestConvertRomanNumeralTo4(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("I V")

	if n != 4 {
		t.Errorf("IV should return 4, but %v", n)
	}
}

func TestConvertRomanNumeralTo1996(t *testing.T) {
	r := NewRoman()

	n := r.ToNumber("M C M X C V I")

	if n != 1996 {
		t.Errorf("MCMXCVI should return 1996, but %v", n)
	}
}

func TestInputDecisionIsQuestion(t *testing.T) {
	q := getTypeInput("how much is pish tegj glob glob ?")
	if q != QUESTION {
		t.Errorf("-how much is pish tegj glob glob ? -  should return QUESTION, but %s ", q)
	}
}

func TestInputDecisionIsRomanAssing(t *testing.T) {
	q := getTypeInput("prok is V")
	if q != ROMAN_ASING {
		t.Errorf("prok is V - should return ROMAN_ASING, but %s ", q)
	}
}

func TestInputDecisionIsValueAssing(t *testing.T) {
	q := getTypeInput("glob glob Silver is 34 Credits")
	if q != VALUE_ASING {
		t.Errorf("ish pish Iron is 3910 Credits - should return VALUE_ASING, but %s ", q)
	}
}

func TestInputDecisionIsExit(t *testing.T) {
	q := getTypeInput("exit")
	if q != EXIT {
		t.Errorf("exit should return EXIT, but %s ", q)
	}
}

func TestInputDecisionIsUnknow(t *testing.T) {
	q := getTypeInput("unknow")
	if q != UNKNOW {
		t.Errorf("unknow should return UNKNOW, but %s ", q)
	}
}

func TestRomanAssingProcess(t *testing.T) {
	romanAssingProcess("prok is V")
	q := num["V"].Word
	if q != "prok" {
		t.Errorf("word of V should return prok, but %s ", q)
	}
}

func TestGetGalacticalNumbersToRoman(t *testing.T) {
	v := num["V"]
	v.Word = "kjhg"
	i := num["I"]
	i.Word = "asas"
	num["V"] = v
	num["I"] = i
	var p = []string{"V", "I"}
	q := getGalacticalNumbersToRoman("kjhg asas Dfdfdf")
	if q != "V I" {
		t.Errorf("kjhg asas should return %s, but %s ", p, q)
	}
}

func TestGetMetal(t *testing.T) {
	q := getMetal("glob prok Gold is 57800 Credits")
	if q != "Gold" {
		t.Errorf("glob prok Gold is 57800 Credits should return Gold, but %s ", q)
	}
}

func TestGetValueAssing(t *testing.T) {
	q := getNumberAssigment("glob prok Gold is 57800 Credits")
	if q != 57800 {
		t.Errorf("glob prok Gold is 57800 Credits should return 57800, but %v ", q)
	}
}

func TestValueAssingProcess(t *testing.T) {
	v := num["V"]
	v.Word = "glob"
	i := num["I"]
	i.Word = "prok"
	num["V"] = v
	num["I"] = i
	valueAssingProcess("glob prok Gold is 57800 Credits")
	if metals["Gold"] != 9633 {
		t.Errorf("product of Gold should return 9633, but %v", metals["Gold"])
	}
}

func TestFindCoin(t *testing.T) {
	coin := getCoin("how many Credits is glob prok Silver ?")
	if "Credits" != coin {
		t.Errorf("how many Credits should return Credits, but %v", coin)
	}
}

func TestFunctional1(t *testing.T) {
	validateInput("glob is I")
	validateInput("prok is V")
	validateInput("pish is X")
	validateInput("tegj is L")
	validateInput("glob glob Silver is 34 Credits")
	validateInput("glob prok Gold is 57800 Credits")
	validateInput("pish pish Iron is 3910 Credits")
	strResult := questionProcess("how many Credits is glob prok Silver ?")
	if "glob prok Silver is 68 Credits" != strResult {
		t.Errorf("product of Gold should return 9633, but %v", metals["Gold"])
	}

}
