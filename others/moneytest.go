package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Options struct {
	DecimalDigits     int8   `json:"decimalDigits"`
	ThousandSeparator string `json:"thousandSeparator"`
	DecimalSeparator  string `json:"decimalSeparator"`
	CurrencySymbol    string `json:"currencySymbol"`
}

//Money represents currency value as Integer
type Money struct {
	value   int
	options Options
}

func MakeFromString(s string, o Options) (Money, error) {
	isNegativeInt := strings.Contains(s, "-")
	m := Money{options: o}
	parts := strings.Split(s, o.DecimalSeparator)
	log.Println("parts", parts)
	var intstr, decstr string
	if len(parts) == 1 {
	} else if len(parts) == 2 {
		intstr = parts[0]
		decstr = parts[1]
	} else {
		return Money{}, fmt.Errorf("Cannot convert %s to money", s)
	}

	log.Println("intstr", intstr)
	log.Println("decstr", decstr)
	intstr = strings.Replace(intstr, o.ThousandSeparator, "", -1)
	intPart, err := strconv.Atoi(intstr)
	if err != nil {
		return Money{}, fmt.Errorf("Cannot convert %s to money", s)
	}

	if intPart < 0 {
		intPart = -1 * intPart
	}
	m.value = intPart * int(math.Pow10(int(o.DecimalDigits)))
	log.Println(int(math.Pow10(int(o.DecimalDigits))))
	decval, err := makeDecPart(decstr, o.DecimalDigits)
	log.Println(" decval", decval)
	if err != nil {
		return Money{}, err
	}
	m.value += decval

	if isNegativeInt {
		m.value = -1 * m.value
	}
	return m, nil
}
func makeDecPart(s string, digits int8) (int, error) {

	var val int
	ds := strings.TrimSpace(s)
	log.Println("trim", ds)
	if len(ds) == 0 {
		return 0, nil
	}
	for i := 0; i < len(ds); i++ {
		if !unicode.IsDigit(rune(ds[i])) {
			return 0, fmt.Errorf("It is not unicode", ds)
		}
		log.Println("ds", i, ds[i])
		digitstr := ds[i : i+1]
		dig, _ := strconv.Atoi(digitstr)
		log.Println("dig", dig)
		log.Println(int(digits - 1 - int8(i)))
		val += dig * int(math.Pow10(int(digits-1-int8(i))))
		log.Println("val", val)
		if digits-1-int8(i) == 0 {
			//if next digit available then round and return else return val
			if i < len(ds)-1 {
				nextDig, err := strconv.Atoi(ds[i+1 : i+2])
				if err != nil {
					return 0, fmt.Errorf("Cannot convert %s to decimal part of money", ds)
				}
				if nextDig >= 5 {
					val++
				}
			}
			return val, nil
		}
	}
	return val, nil
}

func main() {
	money := "100,000.5555"
	a := 10
	log.Println(rune(a))
	options := Options{DecimalDigits: 2, ThousandSeparator: ",", DecimalSeparator: ".", CurrencySymbol: "$"}
	data, err := MakeFromString(money, options)

	log.Println(data.value)
	//log.Println(data)
	log.Println(data)
	log.Println(err)
}
