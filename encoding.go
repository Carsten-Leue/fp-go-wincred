package main

import (
	B "github.com/IBM/fp-go/v2/bytes"
	"github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/optics/iso"
	"github.com/IBM/fp-go/v2/optics/prism"
	"github.com/IBM/fp-go/v2/option"
	"github.com/IBM/fp-go/v2/result"
	"golang.org/x/text/encoding/unicode"
)

// UTF16LEString returns a Prism that converts between UTF-16 Little Endian encoded
// byte slices and UTF-8 strings.
//
// The Prism provides bidirectional conversion:
//   - GetOption: Decodes UTF-16 LE bytes to a UTF-8 string, returning None if decoding fails
//   - ReverseGet: Encodes a UTF-8 string to UTF-16 LE bytes, returning empty bytes on failure
//
// This is useful for interoperating with Windows APIs that use UTF-16 LE encoding,
// such as the Windows Credential Manager.
func UTF16LEString() Prism[[]byte, string] {
	utf8 := iso.UTF8String()
	encoding := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	encoder := result.Eitherize1(encoding.NewEncoder().Bytes)
	decoder := result.Eitherize1(encoding.NewDecoder().Bytes)

	return prism.MakePrismWithName(
		function.Flow3(
			decoder,
			result.ToOption,
			option.Map(utf8.Get),
		),
		function.Flow3(
			utf8.ReverseGet,
			encoder,
			result.GetOrElse(function.Ignore1of1[error](B.Empty)),
		),
		"UTF16LE",
	)
}
