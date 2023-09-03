package main

import (
	"fmt"
	"math"
	"strings"
)

const b94chars = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

// ToBase94 converts a decimal number to its base-94 representation.
func ToBase94(n int64) string {
	if n == 0 {
		return "!"
	}

	var s string
	for n > 0 {
		remainder := n % 94
		s = string(b94chars[remainder]) + s
		n /= 94
	}

	return s
}

// FromBase94 converts a base-94 number (as a string) to its decimal representation.
func FromBase94(s string) int64 {
	var n int64
	for _, char := range s {
		n = n*94 + int64(strings.Index(b94chars, string(char)))
	}

	return n
}

///// b64

const b64chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"

// ToBase64 converts a decimal number to its base-64 representation.
func ToBase64(n int64) string {
	if n == 0 {
		return "!"
	}

	var s string
	for n > 0 {
		remainder := n % 64
		s = string(b64chars[remainder]) + s
		n /= 64
	}

	return s
}

// FromBase64 converts a base-64 number (as a string) to its decimal representation.
func FromBase64(s string) int64 {
	var n int64
	for _, char := range s {
		n = n*64 + int64(strings.Index(b64chars, string(char)))
	}

	return n
}

///// b58

const b58chars = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

// ToBase58 converts a decimal number to its base-58 representation.
func ToBase58(n int64) string {
	if n == 0 {
		return "!"
	}

	var s string
	for n > 0 {
		remainder := n % 58
		s = string(b58chars[remainder]) + s
		n /= 58
	}

	return s
}

// FromBase58 converts a base-58 number (as a string) to its decimal representation.
func FromBase58(s string) int64 {
	var n int64
	for _, char := range s {
		n = n*58 + int64(strings.Index(b58chars, string(char)))
	}

	return n
}

func main() {
	fmt.Printf("MaxInt: %d\n", math.MaxInt64)

	fmt.Printf("Base58: %s\n", ToBase58(math.MaxInt64))
	fmt.Printf("Base64: %s\n", ToBase64(math.MaxInt64))
	fmt.Printf("Base94: %s\n", ToBase94(math.MaxInt64))
}
