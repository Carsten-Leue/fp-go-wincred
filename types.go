package main

import (
	"github.com/IBM/fp-go/v2/context/readerioresult"
	"github.com/IBM/fp-go/v2/effect"
	"github.com/IBM/fp-go/v2/optics/codec"
	"github.com/IBM/fp-go/v2/optics/codec/decode"
	"github.com/IBM/fp-go/v2/optics/prism"
	"github.com/IBM/fp-go/v2/option"
	"github.com/IBM/fp-go/v2/result"
)

type (
	// ReaderIOResult[A] is a lazy computation that depends on a [context.Context],
	// may perform side effects, can fail with an [error], and produces a value of
	// type A on success. It expands to:
	//
	//	func(context.Context) func() Either[error, A]
	//
	// Effects are composed using [readerioresult.Map] and [readerioresult.Chain],
	// and executed by calling the function with a context, then invoking the result:
	//
	//	result := myEffect(ctx)()
	ReaderIOResult[A any] = readerioresult.ReaderIOResult[A]

	// Result[A] represents a value that is either a success (Right[A]) or an
	// error (Left[error]). It is an alias for Either[error, A].
	Result[A any] = result.Result[A]

	// Effect[C, A] is the general form of [ReaderIOResult], parameterised over
	// both the context type C and the success type A. [ReaderIOResult] is the
	// specialisation where C is fixed to [context.Context].
	Effect[C, A any] = effect.Effect[C, A]

	// Prism[S, A] is an optic for accessing a value of type A inside a structure
	// of type S, where the value may not always be present. It provides two
	// operations:
	//
	//   - GetOption(S) Option[A]  – extract the focused value, returning None if absent
	//   - ReverseGet(A) S         – construct an S with only the focused field set
	//
	// String prisms return None for empty strings; byte-slice prisms return None
	// for nil or empty slices.
	Prism[S, A any] = prism.Prism[S, A]

	// Option[A] represents an optional value: either Some(A) when a value is
	// present, or None when it is absent. Use [option.Of] to construct Some and
	// [option.None] to construct None.
	Option[A any] = option.Option[A]

	// Type[A, O, I] is a bidirectional codec that combines decoding (I → A with
	// validation) and encoding (A → O). It is used by [UTF16LEType] to represent
	// the UTF-16 LE ↔ string transformation as a validated codec.
	//
	// Type parameters:
	//   - A: the in-memory Go type (e.g. string)
	//   - O: the encoded output type (e.g. []byte)
	//   - I: the raw input type for decoding (e.g. []byte)
	Type[A, O, I any] = codec.Type[A, O, I]

	// Decode[I, A] is a context-aware validation function that transforms an
	// input of type I into a validated value of type A, returning detailed
	// validation errors when the input is invalid.
	Decode[I, A any] = decode.Decode[I, A]
)
