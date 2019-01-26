package main

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Base base ''", args{""}, ""},
		{"Base base 'a'", args{"a"}, "a"},
		{"Example 1", args{"babad"}, "bab"},
		{"Example 2", args{"cbbd"}, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Base case empty string", args{""}, true},
		{"Base case single character string", args{"a"}, true},
		{"Base case two character string", args{"ab"}, false},
		{"Odd length palindrome", args{"aba"}, true},
		{"Even length palindrome", args{"abba"}, true},
		{"Odd length nonpalindrome", args{"abcbd"}, false},
		{"Even length nonpalindrome", args{"dabbac"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
