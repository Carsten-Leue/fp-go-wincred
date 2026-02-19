//go:build windows

package main

import (
	"testing"

	"github.com/IBM/fp-go/v2/array"
	"github.com/IBM/fp-go/v2/context/readerioresult"
	"github.com/IBM/fp-go/v2/either"
	"github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/option"
	"github.com/danieljoos/wincred"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sampleTargetName = "fp-go-wincred-sample-credential"

// TestSample_GetUsername demonstrates composing an effect to get a username
// from a credential and executing at the boundary.
func TestSample_GetUsername(t *testing.T) {
	ctx := t.Context()
	lenses := MakeGenericCredentialRefLenses()
	utf16 := UTF16LEString()

	// Create a test credential first
	testCred := function.Pipe1(
		wincred.NewGenericCredential(sampleTargetName),
		function.Flow2(
			lenses.UserName.Set("sample-user@example.com"),
			lenses.CredentialBlob.Set(utf16.ReverseGet("sample-password")),
		),
	)

	// Create and cleanup
	createEffect := NewGenericCredential(testCred)(sampleTargetName)
	createResult := createEffect(ctx)()
	require.True(t, either.IsRight(createResult), "Failed to create sample credential")

	defer func() {
		deleteEffect := function.Pipe1(
			GetGenericCredential()(sampleTargetName),
			readerioresult.Chain(DeleteGenericCredential()),
		)
		deleteEffect(ctx)()
	}()

	// 📖 Compose an effect to get username, then execute at the boundary
	getUsernameEffect := function.Pipe1(
		GetGenericCredential()(sampleTargetName),
		readerioresult.Map(lenses.UserName.Get),
	)

	// Execute at boundary and extract result
	usernameResult := getUsernameEffect(ctx)()
	username := either.GetOrElse(function.Constant1[error]("unknown"))(usernameResult)

	assert.Equal(t, "sample-user@example.com", username)
}

// TestSample_GetPassword demonstrates composing an effect to get and decode
// a password from UTF-16 LE encoding.
func TestSample_GetPassword(t *testing.T) {
	ctx := t.Context()
	lenses := MakeGenericCredentialRefLenses()
	utf16 := UTF16LEString()

	// Create a test credential first
	testCred := function.Pipe1(
		wincred.NewGenericCredential(sampleTargetName),
		function.Flow2(
			lenses.UserName.Set("password-test-user"),
			lenses.CredentialBlob.Set(utf16.ReverseGet("secret-password-123")),
		),
	)

	createEffect := NewGenericCredential(testCred)(sampleTargetName)
	createResult := createEffect(ctx)()
	require.True(t, either.IsRight(createResult), "Failed to create sample credential")

	defer func() {
		deleteEffect := function.Pipe1(
			GetGenericCredential()(sampleTargetName),
			readerioresult.Chain(DeleteGenericCredential()),
		)
		deleteEffect(ctx)()
	}()

	// 📖 Compose an effect to get and decode password from UTF-16 LE
	getPasswordEffect := function.Pipe2(
		GetGenericCredential()(sampleTargetName),
		readerioresult.Map(lenses.CredentialBlob.Get),
		readerioresult.Map(utf16.GetOption), // Returns Option[string]
	)

	// Execute at the boundary
	passwordResult := getPasswordEffect(ctx)()

	// Handle the result with either.Fold
	password := either.Fold(
		func(err error) Option[string] {
			t.Fatalf("Error: %v", err)
			return option.None[string]()
		},
		func(p Option[string]) Option[string] { return p },
	)(passwordResult)

	assert.Equal(t, option.Of("secret-password-123"), password)
}

// TestSample_CreateCredential demonstrates creating a new credential
// using lenses and UTF-16 LE encoding.
func TestSample_CreateCredential(t *testing.T) {
	ctx := t.Context()
	lenses := MakeGenericCredentialRefLenses()
	utf16 := UTF16LEString()

	// ✏️ Create a new credential using lenses and UTF-16 LE encoding
	newCred := function.Pipe1(
		wincred.NewGenericCredential(sampleTargetName),
		function.Flow2(
			lenses.UserName.Set("user@example.com"),
			lenses.CredentialBlob.Set(utf16.ReverseGet("secret-password")),
		),
	)

	// Compose the create effect, then execute at the boundary
	createEffect := NewGenericCredential(newCred)(sampleTargetName)

	created := either.Fold(
		func(err error) bool {
			t.Fatalf("Failed to create: %v", err)
			return false
		},
		func(_ *wincred.GenericCredential) bool { return true },
	)(createEffect(ctx)())

	assert.True(t, created, "Credential should be created")

	// Cleanup
	deleteEffect := function.Pipe1(
		GetGenericCredential()(sampleTargetName),
		readerioresult.Chain(DeleteGenericCredential()),
	)
	deleteEffect(ctx)()
}

// TestSample_DeleteCredential demonstrates deleting a credential.
func TestSample_DeleteCredential(t *testing.T) {
	ctx := t.Context()
	lenses := MakeGenericCredentialRefLenses()
	utf16 := UTF16LEString()

	// Create a credential to delete
	testCred := function.Pipe1(
		wincred.NewGenericCredential(sampleTargetName),
		function.Flow2(
			lenses.UserName.Set("delete-test-user"),
			lenses.CredentialBlob.Set(utf16.ReverseGet("delete-password")),
		),
	)

	// 🗑️ Compose create-then-delete as a single functional pipeline, then execute at the boundary
	deleteEffect := function.Pipe1(
		NewGenericCredential(testCred)(sampleTargetName),
		readerioresult.Chain(DeleteGenericCredential()),
	)

	deleted := either.Fold(
		func(err error) bool {
			t.Fatalf("Failed to delete: %v", err)
			return false
		},
		func(_ *wincred.GenericCredential) bool { return true },
	)(deleteEffect(ctx)())

	assert.True(t, deleted, "Credential should be deleted")

	// Verify it's gone
	getEffect := GetGenericCredential()(sampleTargetName)
	getResult := getEffect(ctx)()
	assert.True(t, either.IsLeft(getResult), "Credential should not exist after deletion")
}

// TestSample_ListCredentials demonstrates listing and counting credentials.
func TestSample_ListCredentials(t *testing.T) {
	ctx := t.Context()
	lenses := MakeGenericCredentialRefLenses()
	utf16 := UTF16LEString()

	// Create a test credential to ensure we have at least one
	testCred := function.Pipe1(
		wincred.NewGenericCredential(sampleTargetName),
		function.Flow2(
			lenses.UserName.Set("list-test-user"),
			lenses.CredentialBlob.Set(utf16.ReverseGet("list-password")),
		),
	)

	createEffect := NewGenericCredential(testCred)(sampleTargetName)
	createResult := createEffect(ctx)()
	require.True(t, either.IsRight(createResult), "Failed to create sample credential")

	defer func() {
		deleteEffect := function.Pipe1(
			GetGenericCredential()(sampleTargetName),
			readerioresult.Chain(DeleteGenericCredential()),
		)
		deleteEffect(ctx)()
	}()

	// 📋 Compose an effect to count credentials, then execute at the boundary
	countEffect := function.Pipe1(
		ListCredentials(),
		readerioresult.Map(array.Size[*wincred.Credential]),
	)

	// Execute and extract result
	countResult := countEffect(ctx)()
	count := either.GetOrElse(function.Constant1[error](0))(countResult)

	assert.Greater(t, count, 0, "Should have at least one credential")
}

// TestSample_Lenses demonstrates using lenses for field access.
func TestSample_Lenses(t *testing.T) {
	// Create lenses for GenericCredential
	lenses := MakeGenericCredentialLenses()

	credential := wincred.GenericCredential{
		Credential: wincred.Credential{
			TargetName: "original-target",
			UserName:   "original-user",
			Comment:    "original-comment",
		},
	}

	// 📖 Get a field value
	username := lenses.UserName.Get(credential)
	assert.Equal(t, "original-user", username)

	// ✏️ Set a field value (returns a new credential)
	updated := lenses.UserName.Set("new-user")(credential)
	assert.Equal(t, "new-user", updated.UserName)
	assert.Equal(t, "original-user", credential.UserName) // Original unchanged

	// 🔗 Chain multiple updates
	chained := lenses.UserName.Set("user")(
		lenses.Comment.Set("My credential")(
			lenses.TargetName.Set("app/cred")(credential),
		),
	)
	assert.Equal(t, "app/cred", chained.TargetName)
	assert.Equal(t, "My credential", chained.Comment)
	assert.Equal(t, "user", chained.UserName)

	// 🔍 Optional lenses return Option for zero values
	emptyCredential := wincred.GenericCredential{}
	maybeUsername := lenses.UserNameO.Get(emptyCredential)
	assert.Equal(t, option.None[string](), maybeUsername)

	nonEmptyCredential := wincred.GenericCredential{
		Credential: wincred.Credential{UserName: "test-user"},
	}
	maybeUsername2 := lenses.UserNameO.Get(nonEmptyCredential)
	assert.Equal(t, option.Of("test-user"), maybeUsername2)
}

// TestSample_Prisms demonstrates using prisms for partial access.
func TestSample_Prisms(t *testing.T) {
	// Create prisms for GenericCredential
	prisms := MakeGenericCredentialPrisms()

	credential := wincred.GenericCredential{
		Credential: wincred.Credential{
			TargetName:     "my-target",
			CredentialBlob: []byte("secret"),
		},
	}

	// 📖 Get a field (returns Option - None for empty/zero values)
	maybeTarget := prisms.TargetName.GetOption(credential)
	assert.Equal(t, option.Of("my-target"), maybeTarget)

	maybeBlob := prisms.CredentialBlob.GetOption(credential)
	assert.Equal(t, option.Of([]byte("secret")), maybeBlob)

	// Empty values return None
	emptyCredential := wincred.GenericCredential{}
	maybeTargetEmpty := prisms.TargetName.GetOption(emptyCredential)
	assert.Equal(t, option.None[string](), maybeTargetEmpty)

	maybeBlobEmpty := prisms.CredentialBlob.GetOption(emptyCredential)
	assert.Equal(t, option.None[[]byte](), maybeBlobEmpty)

	// ✏️ Create a struct from a single field value
	cred := prisms.UserName.ReverseGet("my-user")
	assert.Equal(t, "my-user", cred.UserName)
}

// TestSample_UTF16LEString demonstrates UTF-16 LE string conversion.
func TestSample_UTF16LEString(t *testing.T) {
	utf16 := UTF16LEString()

	// 🔄 Encode: Go string → UTF-16 LE bytes
	utf16Bytes := utf16.ReverseGet("Hello, 世界! 👋")
	assert.NotEmpty(t, utf16Bytes)

	// 🔄 Decode: UTF-16 LE bytes → Go string
	maybeString := utf16.GetOption(utf16Bytes)
	assert.Equal(t, option.Of("Hello, 世界! 👋"), maybeString)

	// Round-trip test
	original := "Test password with émojis 🔐"
	encoded := utf16.ReverseGet(original)
	decoded := utf16.GetOption(encoded)
	assert.Equal(t, option.Of(original), decoded)
}

// TestSample_UTF16LEWithCredential demonstrates using UTF-16 LE encoding with credentials.
func TestSample_UTF16LEWithCredential(t *testing.T) {
	utf16 := UTF16LEString()
	lenses := MakeGenericCredentialRefLenses()

	// ✏️ Set credential password with UTF-16 LE encoding
	cred := function.Pipe1(
		wincred.NewGenericCredential("my-app/credentials"),
		lenses.CredentialBlob.Set(utf16.ReverseGet("my-secret-password")),
	)

	// 📖 Read and decode credential password
	blob := lenses.CredentialBlob.Get(cred)
	password := utf16.GetOption(blob)

	assert.Equal(t, option.Of("my-secret-password"), password)
}
