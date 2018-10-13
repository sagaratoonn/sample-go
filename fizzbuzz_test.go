package main

import (
	"testing"
)
func TestToFizzbuzzNormalNumber(t *testing.T) {
	if got, want := Fizzbuzz(1), "1"; got != want {
		t.Errorf("FizzBuzz: got %v, want %v", got, want)
	}

	if got, want := Fizzbuzz(2), "2"; got != want {
		t.Errorf("FizzBuzz(2): got %v, want %v", got, want)
	}
}

func TestToFizzbuzzMultiplesOf3(t *testing.T) {
	if got, want := Fizzbuzz(3), "fizz"; got != want {
		t.Errorf("FizzBuzz(3): got %v, want %v", got, want)
	}
}

func TestToFizzbuzzMultiplesOf5(t *testing.T) {
	if got, want := Fizzbuzz(5), "buzz"; got != want {
		t.Errorf("FizzBuzz(5): got %v, want %v", got, want)
	}
}

func TestToFizzbuzzMultiplesOf3and5(t *testing.T) {
	if got, want := Fizzbuzz(15), "fizzbuzz"; got != want {
		t.Errorf("FizzBuzz(15): got %v, want %v", got, want)
	}
}