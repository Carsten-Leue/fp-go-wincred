// Package main provides functional programming wrappers around the Windows Credential Manager API.
// It uses the fp-go library to provide a purely functional interface for credential operations.
package main

import (
	"github.com/IBM/fp-go/v2/context/readerioresult"
	"github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/ioresult"
	"github.com/danieljoos/wincred"
)

// GetGenericCredential returns an Effect that retrieves a generic credential from the
// Windows Credential Manager by its target name.
//
// The returned Effect takes a target name string and produces a ReaderIOResult containing
// either the credential or an error if the credential cannot be found or accessed.
func GetGenericCredential() Effect[string, *wincred.GenericCredential] {
	return function.Flow2(
		ioresult.Eitherize1(wincred.GetGenericCredential),
		readerioresult.FromIOResult,
	)
}

// NewGenericCredential creates a new generic credential in the Windows Credential Manager.
// It takes a credential template and returns an Effect that, when given a target name,
// writes a copy of the credential with that target name to the credential store.
//
// The original credential is not modified; a copy is made with the specified target name.
// Returns the written credential on success or an error if the write operation fails.
func NewGenericCredential(cred *wincred.GenericCredential) Effect[string, *wincred.GenericCredential] {
	return func(targetName string) ReaderIOResult[*wincred.GenericCredential] {
		return function.Pipe1(
			ioresult.TryCatchError(func() (*wincred.GenericCredential, error) {
				cpy := *cred
				cpy.TargetName = targetName
				return &cpy, cpy.Write()
			}),
			readerioresult.FromIOResult,
		)
	}
}

// DeleteGenericCredential returns an Effect that deletes a generic credential from the
// Windows Credential Manager.
//
// The Effect takes a credential and removes it from the credential store.
// Returns the deleted credential on success or an error if the deletion fails.
func DeleteGenericCredential() Effect[*wincred.GenericCredential, *wincred.GenericCredential] {
	return func(cred *wincred.GenericCredential) ReaderIOResult[*wincred.GenericCredential] {
		return function.Pipe1(
			ioresult.TryCatchError(func() (*wincred.GenericCredential, error) {
				cpy := *cred
				return &cpy, cpy.Delete()
			}),
			readerioresult.FromIOResult,
		)
	}
}

// ListCredentials returns a ReaderIOResult that lists all credentials stored in the
// Windows Credential Manager.
//
// Returns a slice of all credentials on success or an error if the list operation fails.
func ListCredentials() ReaderIOResult[[]*wincred.Credential] {
	return function.Pipe1(
		ioresult.Eitherize0(wincred.List)(),
		readerioresult.FromIOResult,
	)
}
