//go:build windows

package main

import (
	"context"
	"testing"

	"github.com/IBM/fp-go/v2/either"
	"github.com/danieljoos/wincred"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTargetName = "fp-go-wincred-test-credential"

// newTestCredential creates a new test credential with proper defaults.
func newTestCredential(username string, password []byte) *wincred.GenericCredential {
	cred := wincred.NewGenericCredential(testTargetName)
	cred.UserName = username
	cred.CredentialBlob = password
	return cred
}

// withCredential creates a credential, runs the test function, and ensures cleanup.
// This implements the WithResource pattern for safe credential lifecycle management.
func withCredential(t *testing.T, cred *wincred.GenericCredential, targetName string, testFn func(*wincred.GenericCredential)) {
	ctx := context.Background()

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
	testCred := newTestCredential("test-user", []byte("test-password"))

	withCredential(t, testCred, testTargetName, func(created *wincred.GenericCredential) {
		ctx := context.Background()

		// Retrieve the credential
		getEffect := GetGenericCredential()(testTargetName)
		getResult := getEffect(ctx)()

		require.True(t, either.IsRight(getResult), "Failed to get credential")

		retrieved := either.Fold(
			func(err error) *wincred.GenericCredential { return nil },
			func(c *wincred.GenericCredential) *wincred.GenericCredential { return c },
		)(getResult)

		require.NotNil(t, retrieved, "Retrieved credential should not be nil")
		assert.Equal(t, testTargetName, retrieved.TargetName)
		assert.Equal(t, "test-user", retrieved.UserName)
		assert.Equal(t, []byte("test-password"), retrieved.CredentialBlob)
	})
}

func TestGetGenericCredential_NotFound(t *testing.T) {
	ctx := context.Background()

	// Try to get a non-existent credential
	getEffect := GetGenericCredential()("non-existent-credential-12345")
	getResult := getEffect(ctx)()

	assert.True(t, either.IsLeft(getResult), "Expected error for non-existent credential")
}

func TestNewGenericCredential(t *testing.T) {
	ctx := context.Background()
	targetName := testTargetName + "-new"

	testCred := newTestCredential("new-user", []byte("new-password"))

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
	assert.Equal(t, targetName, created.TargetName)
	assert.Equal(t, "new-user", created.UserName)
}

func TestDeleteGenericCredential(t *testing.T) {
	ctx := context.Background()
	targetName := testTargetName + "-delete"

	testCred := newTestCredential("delete-user", []byte("delete-password"))

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
	testCred := newTestCredential("list-user", []byte("list-password"))

	withCredential(t, testCred, testTargetName+"-list", func(created *wincred.GenericCredential) {
		ctx := context.Background()

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
		found := false
		for _, cred := range credentials {
			if cred.TargetName == testTargetName+"-list" {
				found = true
				break
			}
		}
		assert.True(t, found, "Test credential should be in the list")
	})
}

func TestCredentialRoundTrip(t *testing.T) {
	ctx := context.Background()
	targetName := testTargetName + "-roundtrip"

	// Test data
	username := "roundtrip-user"
	password := []byte("roundtrip-secret-password-123!")

	testCred := newTestCredential(username, password)

	// Create the credential
	createEffect := NewGenericCredential(testCred)(targetName)
	createResult := createEffect(ctx)()
	require.True(t, either.IsRight(createResult), "Failed to create credential")

	// Get the credential
	getEffect := GetGenericCredential()(targetName)
	getResult := getEffect(ctx)()
	require.True(t, either.IsRight(getResult), "Failed to get credential")

	retrieved := either.Fold(
		func(err error) *wincred.GenericCredential { return nil },
		func(c *wincred.GenericCredential) *wincred.GenericCredential { return c },
	)(getResult)

	// Cleanup
	defer func() {
		deleteEffect := DeleteGenericCredential()(retrieved)
		deleteEffect(ctx)()
	}()

	require.NotNil(t, retrieved, "Retrieved credential should not be nil")
	assert.Equal(t, targetName, retrieved.TargetName)
	assert.Equal(t, username, retrieved.UserName)
	assert.Equal(t, password, retrieved.CredentialBlob)
}
