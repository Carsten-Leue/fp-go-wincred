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
	ReaderIOResult[A any] = readerioresult.ReaderIOResult[A]
	Result[A any]         = result.Result[A]
	Effect[C, A any]      = effect.Effect[C, A]
	Prism[S, A any]       = prism.Prism[S, A]
	Option[A any]         = option.Option[A]
	Type[A, O, I any]     = codec.Type[A, O, I]
	Decode[I, A any]      = decode.Decode[I, A]
)
