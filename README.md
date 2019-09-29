# Galactic Merchant Guide

You decided to leave Earth after the last financial collapse left 99.99% of the population with 0.01% of the wealth. Fortunately, with the small amount of money left in your account, you can rent a spaceship, leave Earth, and travel around the Galaxy to sell common metals and dust (which is apparently very expensive).

To buy and sell around the Galaxy you need to convert numbers and units, so you decided to write a program that helps you.

The numbers used for intergalactic transactions follow a convention similar to that of Roman numerals and with much effort you have collected the appropriate translation between them.
### Installation and Run

Galactic Merchant Guide requires [Go](https://golang.org) to run.

```sh
$ cd galactic_merchant
$ go run main.go
```

or...

```sh
$ cd galactic_merchant
$ go build main.go
$ .\main
```

### Instrucctions



| Types of Input | example |
| ------ | ------ |
| ROMAN_ASING| glob is I, prok is V |
| VALUE_ASING | glob glob Silver is 34 Credits |
| QUESTION | how much is pish tegj glob glob ? |
| EXIT | exit |
| UNKNOW | how much wood could a woodchuck chuck if a woodchuck could chuck wood |


### Example Inputs
- glob is I
- prok is V
- pish is X
- tegj is L
- glob glob Silver is 34 Credits
- glob prok Gold is 57800 Credits
- pish pish Iron is 3910 Credits
- how much is pish tegj glob glob ?
- how many Credits is glob prok Silver ?
- how many Credits is glob prok Gold ?
- how many Credits is glob prok Iron ?
- how much wood could a woodchuck chuck if a woodchuck could chuck wood ?

### Example Outputs
- pish tegj glob glob is 42
- glob prok Silver is 68 Credits
- glob prok Gold is 57800 Credits
- glob prok Iron is 782 Credits
- I have no idea what you are talking about


