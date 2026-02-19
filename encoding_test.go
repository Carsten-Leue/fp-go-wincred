package main

import (
	"testing"

	"github.com/IBM/fp-go/v2/option"
	"github.com/stretchr/testify/assert"
)

func TestUTF16LEString_RoundTrip(t *testing.T) {
	prism := UTF16LEString()

	testCases := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"simple ascii", "hello"},
		{"with spaces", "hello world"},
		{"numbers", "12345"},
		{"special chars", "hello@world.com"},
		{"unicode", "héllo wörld"},
		{"emoji", "hello 👋 world"},
		{"chinese", "你好世界"},
		{"japanese", "こんにちは"},
		{"mixed", "Hello 世界 🌍"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encode string to UTF-16 LE bytes
			encoded := prism.ReverseGet(tc.input)

			// Decode back to string
			decoded := prism.GetOption(encoded)

			assert.Equal(t, option.Of(tc.input), decoded)
		})
	}
}

func TestUTF16LEString_Encode(t *testing.T) {
	prism := UTF16LEString()

	// "AB" in UTF-16 LE is: 0x41 0x00 0x42 0x00
	encoded := prism.ReverseGet("AB")
	expected := []byte{0x41, 0x00, 0x42, 0x00}

	assert.Equal(t, expected, encoded)
}

func TestUTF16LEString_Decode(t *testing.T) {
	prism := UTF16LEString()

	// "AB" in UTF-16 LE is: 0x41 0x00 0x42 0x00
	input := []byte{0x41, 0x00, 0x42, 0x00}
	decoded := prism.GetOption(input)

	assert.Equal(t, option.Of("AB"), decoded)
}

func TestUTF16LEString_DecodeUnicode(t *testing.T) {
	prism := UTF16LEString()

	// "é" (U+00E9) in UTF-16 LE is: 0xE9 0x00
	input := []byte{0xE9, 0x00}
	decoded := prism.GetOption(input)

	assert.Equal(t, option.Of("é"), decoded)
}

func TestUTF16LEString_DecodeSurrogatePair(t *testing.T) {
	prism := UTF16LEString()

	// "😀" (U+1F600) in UTF-16 LE is a surrogate pair: 0x3D 0xD8 0x00 0xDE
	input := []byte{0x3D, 0xD8, 0x00, 0xDE}
	decoded := prism.GetOption(input)

	assert.Equal(t, option.Of("😀"), decoded)
}

func TestUTF16LEString_DecodeEmpty(t *testing.T) {
	prism := UTF16LEString()

	decoded := prism.GetOption([]byte{})

	assert.Equal(t, option.Of(""), decoded)
}

func TestUTF16LEString_DecodeOddBytes(t *testing.T) {
	prism := UTF16LEString()

	// Odd number of bytes is invalid UTF-16
	input := []byte{0x41, 0x00, 0x42}
	decoded := prism.GetOption(input)

	// The decoder may handle this gracefully or return None
	// depending on the implementation - we just verify it doesn't panic
	if option.IsSome(decoded) {
		// If it succeeds, verify we get some result
		result := option.GetOrElse(func() string { return "" })(decoded)
		assert.NotEmpty(t, result)
	}
}

func TestUTF16LEString_EncodeEmpty(t *testing.T) {
	prism := UTF16LEString()

	encoded := prism.ReverseGet("")

	assert.Empty(t, encoded)
}

func TestUTF16LEString_PrismLaws(t *testing.T) {
	prism := UTF16LEString()

	t.Run("ReverseGet then GetOption returns original", func(t *testing.T) {
		original := "test string"
		encoded := prism.ReverseGet(original)
		decoded := prism.GetOption(encoded)

		assert.Equal(t, option.Of(original), decoded)
	})

	t.Run("GetOption then ReverseGet preserves bytes for valid input", func(t *testing.T) {
		// Valid UTF-16 LE for "test"
		original := []byte{0x74, 0x00, 0x65, 0x00, 0x73, 0x00, 0x74, 0x00}
		decoded := prism.GetOption(original)

		assert.Equal(t, option.Of("test"), decoded)

		str := option.GetOrElse(func() string { return "" })(decoded)
		reencoded := prism.ReverseGet(str)

		assert.Equal(t, original, reencoded)
	})
}
