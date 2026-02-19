# 🔐 fp-go-wincred

A functional programming wrapper around the Windows Credential Manager API, built with [fp-go](https://github.com/IBM/fp-go) and [wincred](https://github.com/danieljoos/wincred).

## ✨ Features

- 🎯 **Effect-based API** - All credential operations use the `Effect` type for composable, pure functional programming
- 🔍 **Optics (Lenses & Prisms)** - Type-safe accessors for reading and modifying credential fields
- 🔄 **UTF-16 LE Encoding** - Built-in prism for converting between Go strings and Windows UTF-16 LE byte arrays
- 🛡️ **Type Safety** - Leverages Go generics for compile-time type checking

## 📦 Installation

```bash
go get github.com/carst/fp-go-wincred
```

## 🚀 Usage

### Managing Credentials with Effects

All credential operations return an `Effect` type, which represents a computation that:
- Takes a context parameter
- May produce a result or an error
- Is lazily evaluated (only runs when invoked)

```go
import (
    "context"
    "fmt"
    "github.com/IBM/fp-go/v2/array"
    "github.com/IBM/fp-go/v2/context/readerioresult"
    "github.com/IBM/fp-go/v2/either"
    "github.com/IBM/fp-go/v2/function"
)

// 🔍 Create lenses and prisms for accessing credential fields
lenses := MakeGenericCredentialRefLenses()
utf16 := UTF16LEString()
ctx := context.Background()

// 📖 Compose an effect to get username, then execute at the boundary
getUsernameEffect := function.Pipe3(
    GetGenericCredential()("my-app/credentials"),
    readerioresult.Map(lenses.UserName.Get),
    readerioresult.GetOrElse(function.Constant1[error]("unknown")),
)
username := getUsernameEffect(ctx)()
fmt.Printf("Username: %s\n", username)

// 📖 Compose an effect to get and decode password from UTF-16 LE
getPasswordEffect := function.Pipe3(
    GetGenericCredential()("my-app/credentials"),
    readerioresult.Map(lenses.CredentialBlob.Get),
    readerioresult.Map(utf16.GetOption),  // Returns Option[string]
)
// Execute at the boundary
passwordResult := getPasswordEffect(ctx)()
function.Pipe1(
    passwordResult,
    either.Fold(
        func(err error) { fmt.Printf("Error: %v\n", err) },
        func(password Option[string]) { fmt.Printf("Password: %v\n", password) },
    ),
)

// ✏️ Create a new credential using lenses and UTF-16 LE encoding
newCred := function.Pipe1(
    wincred.NewGenericCredential("my-app/credentials"),
    function.Flow2(
        lenses.UserName.Set("user@example.com"),
        lenses.CredentialBlob.Set(utf16.ReverseGet("secret-password")),
    ),
)

// Compose the create effect, then execute at the boundary
createEffect := NewGenericCredential(newCred)("my-app/credentials")
function.Pipe1(
    createEffect(ctx)(),
    either.Fold(
        func(err error) { fmt.Printf("Failed to create: %v\n", err) },
        func(_ *wincred.GenericCredential) { fmt.Println("Created successfully!") },
    ),
)

// 🗑️ Compose the delete effect, then execute at the boundary
deleteEffect := DeleteGenericCredential()(existingCred)
function.Pipe1(
    deleteEffect(ctx)(),
    either.Fold(
        func(err error) { fmt.Printf("Failed to delete: %v\n", err) },
        func(_ *wincred.GenericCredential) { fmt.Println("Deleted successfully!") },
    ),
)

// 📋 Compose an effect to count credentials, then execute at the boundary
countEffect := function.Pipe3(
    ListCredentials(),
    readerioresult.Map(array.Size[*wincred.Credential]),
    readerioresult.GetOrElse(function.Constant1[error](0)),
)
count := countEffect(ctx)()
fmt.Printf("Found %d credentials\n", count)
```

### 🔍 Using Lenses for Field Access

Lenses provide a composable way to get and set fields on credential structures:

```go
// Create lenses for GenericCredential
lenses := MakeGenericCredentialLenses()

// 📖 Get a field value
username := lenses.UserName.Get(credential)

// ✏️ Set a field value (returns a new credential)
updated := lenses.UserName.Set("new-user")(credential)

// 🔗 Chain multiple updates
updated := lenses.UserName.Set("user")(
    lenses.Comment.Set("My credential")(
        lenses.TargetName.Set("app/cred")(credential),
    ),
)

// 🔍 Optional lenses return Option for zero values
maybeUsername := lenses.UserNameO.Get(credential)  // Option[string]
```

### 🔮 Using Prisms for Partial Access

Prisms are used when a value may or may not be present:

```go
// Create prisms for GenericCredential
prisms := MakeGenericCredentialPrisms()

// 📖 Get a field (returns Option - None for empty/zero values)
maybeTarget := prisms.TargetName.GetOption(credential)  // None if empty string
maybeBlob := prisms.CredentialBlob.GetOption(credential)  // None if empty []byte

// ✏️ Create a struct from a single field value
cred := prisms.UserName.ReverseGet("my-user")
```

### 🔄 UTF-16 LE String Conversion

Windows APIs often use UTF-16 Little Endian encoding. Use the `UTF16LEString` prism to convert:

```go
utf16 := UTF16LEString()
lenses := MakeGenericCredentialRefLenses()

// 🔄 Encode: Go string → UTF-16 LE bytes
utf16Bytes := utf16.ReverseGet("Hello, 世界! 👋")

// 🔄 Decode: UTF-16 LE bytes → Go string
maybeString := utf16.GetOption(utf16Bytes)  // Option[string]

// ✏️ Set credential password with UTF-16 LE encoding
cred := wincred.NewGenericCredential("my-app/credentials")
cred = lenses.CredentialBlob.Set(utf16.ReverseGet("my-secret-password"))(cred)

// 📖 Read and decode credential password
blob := lenses.CredentialBlob.Get(cred)
password := utf16.GetOption(blob)  // Option[string]
```

## 📚 API Reference

### Effect Types

| Type | Description |
|------|-------------|
| `Effect[C, A]` | A computation taking context `C` and producing `A` or error |
| `ReaderIOResult[A]` | An effect that reads from context and produces `A` or error |

### Credential Operations

| Function | Description |
|----------|-------------|
| `GetGenericCredential()` | Retrieve a credential by target name |
| `NewGenericCredential(cred)` | Create/update a credential |
| `DeleteGenericCredential()` | Delete a credential |
| `ListCredentials()` | List all stored credentials |

### Optics Factories

| Function | Description |
|----------|-------------|
| `MakeCredentialLenses()` | Lenses for `wincred.Credential` |
| `MakeCredentialRefLenses()` | Lenses for `*wincred.Credential` |
| `MakeCredentialPrisms()` | Prisms for `wincred.Credential` |
| `MakeGenericCredentialLenses()` | Lenses for `wincred.GenericCredential` |
| `MakeGenericCredentialRefLenses()` | Lenses for `*wincred.GenericCredential` |
| `MakeGenericCredentialPrisms()` | Prisms for `wincred.GenericCredential` |
| `UTF16LEString()` | Prism for UTF-16 LE ↔ string conversion |

## 🔗 Dependencies

- [IBM/fp-go](https://github.com/IBM/fp-go) - Functional programming library for Go
- [danieljoos/wincred](https://github.com/danieljoos/wincred) - Windows Credential Manager bindings

## 📄 License

MIT License

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
