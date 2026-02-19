package main

import (
	"github.com/IBM/fp-go/v2/context/readerioresult"
	"github.com/IBM/fp-go/v2/effect"
	"github.com/IBM/fp-go/v2/optics/prism"
	"github.com/IBM/fp-go/v2/option"
)

type (
	ReaderIOResult[A any] = readerioresult.ReaderIOResult[A]
	Effect[C, A any]      = effect.Effect[C, A]
	Prism[S, A any]       = prism.Prism[S, A]
	Option[A any]         = option.Option[A]
)
