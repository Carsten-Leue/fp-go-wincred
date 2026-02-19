# AGENTS.md

Instructions for AI agents working on this codebase.

## Project Overview

This is a functional programming wrapper around the Windows Credential Manager API, built with [fp-go](https://github.com/IBM/fp-go) and [wincred](https://github.com/danieljoos/wincred).

## Code Style and Patterns

### Functional Programming with fp-go

This codebase uses the IBM fp-go library (v2) for functional programming patterns. Key concepts:

- **Effect type**: All credential operations return `Effect[C, A]` which represents a lazy, composable computation
- **Either monad**: Use `either.Fold`, `either.Map`, `either.GetOrElse` for error handling - avoid `IsRight`/`IsLeft` checks followed by extraction
- **Option monad**: Use `option.Of`, `option.None`, `option.Map`, `option.GetOrElse` for optional values
- **Function composition**: Use `function.Pipe1`, `function.Pipe2`, `function.Flow2`, `function.Flow3` for composing operations

### Lenses and Prisms

Always use lenses and prisms from `lenses.go` for field access instead of direct struct field access:

```go
// CORRECT: Use lenses
lenses := MakeGenericCredentialRefLenses()
username := lenses.UserName.Get(cred)
updated := lenses.UserName.Set("new-user")(cred)

// INCORRECT: Direct field access
username := cred.UserName
cred.UserName = "new-user"
```

Available lens factories:
- `MakeGenericCredentialLenses()` - for `wincred.GenericCredential` values
- `MakeGenericCredentialRefLenses()` - for `*wincred.GenericCredential` pointers
- `MakeCredentialLenses()` - for `wincred.Credential` values
- `MakeCredentialRefLenses()` - for `*wincred.Credential` pointers
- `MakeCredentialAttributeLenses()` - for `wincred.CredentialAttribute` values
- `MakeCredentialAttributeRefLenses()` - for `*wincred.CredentialAttribute` pointers

Prism behavior:
- String prisms return `None` for empty strings
- Byte slice prisms return `None` for empty or nil slices
- Use `prism.GetOption()` to get an `Option[A]`
- Use `prism.ReverseGet()` to create a struct from a field value

### Either/Option Assertions in Tests

Prefer direct option/either comparison over boolean checks for final assertions:

```go
// CORRECT: Direct comparison for assertions
assert.Equal(t, option.Of("expected"), prism.GetOption(value))
assert.Equal(t, option.None[string](), prism.GetOption(emptyValue))
assert.Equal(t, either.Of[error]("expected"), getUsernameEffect(ctx)())

// ACCEPTABLE: Boolean checks for preconditions/setup
require.True(t, either.IsRight(createResult), "Failed to create credential")

// INCORRECT: Boolean checks followed by extraction for final assertions
require.True(t, option.IsSome(result))
actual := option.GetOrElse(func() string { return "" })(result)
assert.Equal(t, "expected", actual)
```

### Effect Pattern for Credential Operations

Compose effects first using `readerioresult.Map`, `readerioresult.Chain`, etc., then execute at the boundary:

```go
// Compose effects - no execution yet
getUsernameEffect := function.Pipe1(
    GetGenericCredential()("target-name"),
    readerioresult.Map(lenses.UserName.Get),
)

// Execute at the boundary
result := getUsernameEffect(ctx)()

// Handle the result with either.Fold
either.Fold(
    func(err error) { /* handle error */ },
    func(username string) { /* handle success */ },
)(result)
```

Key functions for effect composition:
- `readerioresult.Map(f)` - Transform the success value
- `readerioresult.Chain(f)` - Sequence effects (flatMap)

After executing at the boundary, use `either` functions to handle the result:
- `either.Fold(onLeft, onRight)` - Handle both error and success cases
- `either.GetOrElse(f)` - Extract value with default on error
- `either.Map(f)` - Transform success value

## Testing

### Platform-Specific Tests

Tests that interact with the Windows Credential Manager must have the build tag:

```go
//go:build windows
```

Tests for platform-independent code (like `encoding_test.go`, `lenses_test.go`) should NOT have this tag.

### Test Resource Management

Use the `withCredential` helper for tests that create credentials:

```go
withCredential(t, testCred, targetName, func(created *wincred.GenericCredential) {
    // Test code here
    // Cleanup is automatic
})
```

### Creating Test Credentials

Use `wincred.NewGenericCredential()` to create credentials with proper defaults, then use lenses to set fields. Always encode passwords using the `UTF16LEString` prism:

```go
utf16 := UTF16LEString()
lenses := MakeGenericCredentialRefLenses()

cred := wincred.NewGenericCredential(targetName)
cred = lenses.UserName.Set(username)(cred)
cred = lenses.CredentialBlob.Set(utf16.ReverseGet(password))(cred)
```

### UTF-16 LE Encoding for Credential Blobs

Windows Credential Manager stores credential blobs as UTF-16 LE encoded bytes. Always use the `UTF16LEString` prism for encoding/decoding:

```go
utf16 := UTF16LEString()

// Encoding: string → []byte
encoded := utf16.ReverseGet("my-password")

// Decoding: []byte → Option[string]
decoded := utf16.GetOption(blob)
```

## File Structure

- `wincred.go` - Core credential operations (Get, New, Delete, List)
- `types.go` - Type aliases for fp-go types
- `lenses.go` - Optics (lenses and prisms) for credential types
- `encoding.go` - UTF-16 LE string conversion prism
- `*_test.go` - Unit tests

## Dependencies

- `github.com/IBM/fp-go/v2` - Functional programming library
- `github.com/danieljoos/wincred` - Windows Credential Manager bindings
- `golang.org/x/text` - UTF-16 encoding support
- `github.com/stretchr/testify` - Testing assertions

## CI/CD

- GitHub Actions runs tests on Go 1.24.x, 1.25.x, and 1.26.x
- Semantic release generates versions based on conventional commits
- Renovate updates dependencies (excluding Go version in go.mod)
