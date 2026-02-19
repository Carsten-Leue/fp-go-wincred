//go:build windows

package main

import (
	"testing"

	"github.com/IBM/fp-go/v2/context/readerioresult"
	"github.com/IBM/fp-go/v2/either"
	"github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/option"
	"github.com/danieljoos/wincred"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTargetName = "fp-go-wincred-test-credential"

// refLenses provides lenses for accessing fields on *wincred.GenericCredential.
var refLenses = MakeGenericCredentialRefLenses()

// utf16Prism provides UTF-16 LE encoding/decoding for credential blobs.
var utf16Prism = UTF16LEString()

// newTestCredential creates a new test credential with proper defaults.
// The password is encoded as UTF-16 LE in the credential blob.
func newTestCredential(username string, password string) *wincred.GenericCredential {
	cred := wincred.NewGenericCredential(testTargetName)
	cred = refLenses.UserName.Set(username)(cred)
	cred = refLenses.CredentialBlob.Set(utf16Prism.ReverseGet(password))(cred)
	return cred
}

// withCredential creates a credential, runs the test function, and ensures cleanup.
// This implements the WithResource pattern for safe credential lifecycle management.
func withCredential(t *testing.T, cred *wincred.GenericCredential, targetName string, testFn func(*wincred.GenericCredential)) {
	ctx := t.Context()

	// Create the credential
	createEffect := NewGenericCredential(cred)(targetName)
	createResult := createEffect(ctx)()

	require.True(t, either.IsRight(createResult), "Failed to create test credential")

	createdCred := either.Fold(
		func(err error) *wincred.GenericCredential { return nil },
		func(c *wincred.GenericCredential) *wincred.GenericCredential { return c },
	)(createResult)

	require.NotNil(t, createdCred, "Created credential should not be nil")

	// Ensure cleanup happens regardless of test outcome
	defer func() {
		deleteEffect := DeleteGenericCredential()(createdCred)
		deleteResult := deleteEffect(ctx)()
		if either.IsLeft(deleteResult) {
			t.Logf("Warning: failed to delete test credential")
		}
	}()

	// Run the test function
	testFn(createdCred)
}

func TestGetGenericCredential(t *testing.T) {
	testCred := newTestCredential("test-user", "test-password")

	withCredential(t, testCred, testTargetName, func(created *wincred.GenericCredential) {
		ctx := t.Context()

		// Compose effects to extract fields
		getTargetEffect := function.Pipe1(
			GetGenericCredential()(testTargetName),
			readerioresult.Map(refLenses.TargetName.Get),
		)
		getUserEffect := function.Pipe1(
			GetGenericCredential()(testTargetName),
			readerioresult.Map(refLenses.UserName.Get),
		)
		getBlobEffect := function.Pipe1(
			GetGenericCredential()(testTargetName),
			readerioresult.Map(refLenses.CredentialBlob.Get),
		)

		// Execute at the boundary and assert
		assert.Equal(t, either.Of[error](testTargetName), getTargetEffect(ctx)())
		assert.Equal(t, either.Of[error]("test-user"), getUserEffect(ctx)())
		assert.Equal(t, either.Of[error](utf16Prism.ReverseGet("test-password")), getBlobEffect(ctx)())
	})
}

func TestGetGenericCredential_NotFound(t *testing.T) {
	ctx := t.Context()

	// Try to get a non-existent credential
	getEffect := GetGenericCredential()("non-existent-credential-12345")
	getResult := getEffect(ctx)()

	assert.True(t, either.IsLeft(getResult), "Expected error for non-existent credential")
}

func TestNewGenericCredential(t *testing.T) {
	ctx := t.Context()
	targetName := testTargetName + "-new"

	testCred := newTestCredential("new-user", "new-password")

	// Create credential
	createEffect := NewGenericCredential(testCred)(targetName)
	createResult := createEffect(ctx)()

	require.True(t, either.IsRight(createResult), "Failed to create credential")

	created := either.Fold(
		func(err error) *wincred.GenericCredential { return nil },
		func(c *wincred.GenericCredential) *wincred.GenericCredential { return c },
	)(createResult)

	// Cleanup
	defer func() {
		deleteEffect := DeleteGenericCredential()(created)
		deleteEffect(ctx)()
	}()

	require.NotNil(t, created, "Created credential should not be nil")
	assert.Equal(t, targetName, refLenses.TargetName.Get(created))
	assert.Equal(t, "new-user", refLenses.UserName.Get(created))
}

func TestDeleteGenericCredential(t *testing.T) {
	ctx := t.Context()
	targetName := testTargetName + "-delete"

	testCred := newTestCredential("delete-user", "delete-password")

	// Create credential first
	createEffect := NewGenericCredential(testCred)(targetName)
	createResult := createEffect(ctx)()
	require.True(t, either.IsRight(createResult), "Failed to create credential for deletion test")

	created := either.Fold(
		func(err error) *wincred.GenericCredential { return nil },
		func(c *wincred.GenericCredential) *wincred.GenericCredential { return c },
	)(createResult)

	// Delete the credential
	deleteEffect := DeleteGenericCredential()(created)
	deleteResult := deleteEffect(ctx)()

	require.True(t, either.IsRight(deleteResult), "Failed to delete credential")

	// Verify it's gone
	getEffect := GetGenericCredential()(targetName)
	getResult := getEffect(ctx)()

	assert.True(t, either.IsLeft(getResult), "Credential should not exist after deletion")
}

func TestListCredentials(t *testing.T) {
	testCred := newTestCredential("list-user", "list-password")

	withCredential(t, testCred, testTargetName+"-list", func(created *wincred.GenericCredential) {
		ctx := t.Context()

		// List all credentials
		listEffect := ListCredentials()
		listResult := listEffect(ctx)()

		require.True(t, either.IsRight(listResult), "Failed to list credentials")

		credentials := either.Fold(
			func(err error) []*wincred.Credential { return nil },
			func(c []*wincred.Credential) []*wincred.Credential { return c },
		)(listResult)

		require.NotNil(t, credentials, "Credential list should not be nil")
		assert.NotEmpty(t, credentials, "Credential list should not be empty")

		// Verify our test credential is in the list
		credLenses := MakeCredentialRefLenses()
		found := false
		for _, cred := range credentials {
			if credLenses.TargetName.Get(cred) == testTargetName+"-list" {
				found = true
				break
			}
		}
		assert.True(t, found, "Test credential should be in the list")
	})
}

func TestCredentialRoundTrip(t *testing.T) {
	ctx := t.Context()
	targetName := testTargetName + "-roundtrip"

	// Test data
	username := "roundtrip-user"
	password := "roundtrip-secret-password-123!"

	testCred := newTestCredential(username, password)

	// Compose create effect and execute at the boundary
	createEffect := NewGenericCredential(testCred)(targetName)
	createResult := createEffect(ctx)()
	require.True(t, either.IsRight(createResult), "Failed to create credential")

	// Cleanup - compose delete effect
	defer func() {
		deleteEffect := function.Pipe1(
			GetGenericCredential()(targetName),
			readerioresult.Chain(DeleteGenericCredential()),
		)
		deleteEffect(ctx)()
	}()

	// Compose effects to verify round-trip using readerioresult.Map
	getUsernameEffect := function.Pipe1(
		GetGenericCredential()(targetName),
		readerioresult.Map(refLenses.UserName.Get),
	)
	getPasswordEffect := function.Pipe1(
		GetGenericCredential()(targetName),
		readerioresult.Map(function.Flow2(
			refLenses.CredentialBlob.Get,
			utf16Prism.GetOption,
		)),
	)

	// Execute at the boundary and assert
	assert.Equal(t, either.Of[error](username), getUsernameEffect(ctx)())
	assert.Equal(t, either.Of[error](option.Of(password)), getPasswordEffect(ctx)())
}
