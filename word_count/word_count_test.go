package main

import "testing"

func TestWordCounterOneWord(t *testing.T) {
	got := WordCounter("hello")
	expected := 1
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterTwoWords(t *testing.T) {
	got := WordCounter("hello world")
	expected := 2
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterMultitpleWords(t *testing.T) {
	got := WordCounter("hello world this is a test")
	expected := 6
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterEmptyString(t *testing.T) {
	got := WordCounter("")
	expected := 0
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterNoSplitOneWord(t *testing.T) {
	got := WordCounter("hello")
	expected := 1
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterNoSplitTwoWords(t *testing.T) {
	got := WordCounter("hello world")
	expected := 2
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterNoSplitMultitpleWords(t *testing.T) {
	got := WordCounter("hello world this is a test")
	expected := 6
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestWordCounterNoSplitEmptyString(t *testing.T) {
	got := WordCounter("")
	expected := 0
	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
