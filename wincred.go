package main

import (
	"github.com/IBM/fp-go/v2/context/readerioresult"
	"github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/ioresult"
	"github.com/danieljoos/wincred"
)

var (
	GetGenericCredential = function.Flow2(
		ioresult.Eitherize1(wincred.GetGenericCredential),
		readerioresult.FromIOResult,
	)

	ListCredentials = function.Pipe1(
		ioresult.Eitherize0(wincred.List)(),
		readerioresult.FromIOResult,
	)
)
