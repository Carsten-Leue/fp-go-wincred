package main

import (
	"testing"
	"time"

	"github.com/IBM/fp-go/v2/function"
	"github.com/IBM/fp-go/v2/option"
	"github.com/danieljoos/wincred"
	"github.com/stretchr/testify/assert"
)

// TestCredentialAttributeLenses_GetSet verifies lens laws for CredentialAttribute
func TestCredentialAttributeLenses_GetSet(t *testing.T) {
	lenses := MakeCredentialAttributeLenses()

	t.Run("Keyword lens", func(t *testing.T) {
		attr := wincred.CredentialAttribute{Keyword: "test-key", Value: []byte("test-value")}

		// Get-Set law: Setting what you get changes nothing
		got := lenses.Keyword.Get(attr)
		updated := lenses.Keyword.Set(got)(attr)
		assert.Equal(t, attr, updated)

		// Set-Get law: Getting what you set returns what you set
		newKeyword := "new-key"
		updated = lenses.Keyword.Set(newKeyword)(attr)
		assert.Equal(t, newKeyword, lenses.Keyword.Get(updated))

		// Set-Set law: Setting twice is the same as setting once
		updated1 := lenses.Keyword.Set("first")(attr)
		updated2 := lenses.Keyword.Set("second")(updated1)
		updated3 := lenses.Keyword.Set("second")(attr)
		assert.Equal(t, lenses.Keyword.Get(updated2), lenses.Keyword.Get(updated3))
	})

	t.Run("Value lens", func(t *testing.T) {
		attr := wincred.CredentialAttribute{Keyword: "test-key", Value: []byte("test-value")}

		// Get-Set law
		got := lenses.Value.Get(attr)
		updated := lenses.Value.Set(got)(attr)
		assert.Equal(t, attr.Keyword, updated.Keyword)
		assert.Equal(t, attr.Value, updated.Value)

		// Set-Get law
		newValue := []byte("new-value")
		updated = lenses.Value.Set(newValue)(attr)
		assert.Equal(t, newValue, lenses.Value.Get(updated))
	})
}

// TestCredentialAttributeRefLenses_GetSet verifies lens laws for CredentialAttribute references
func TestCredentialAttributeRefLenses_GetSet(t *testing.T) {
	lenses := MakeCredentialAttributeRefLenses()

	t.Run("Keyword ref lens", func(t *testing.T) {
		attr := &wincred.CredentialAttribute{Keyword: "test-key", Value: []byte("test-value")}

		// Set-Get law
		newKeyword := "new-key"
		updated := lenses.Keyword.Set(newKeyword)(attr)
		assert.Equal(t, newKeyword, lenses.Keyword.Get(updated))
	})

	t.Run("Value ref lens", func(t *testing.T) {
		attr := &wincred.CredentialAttribute{Keyword: "test-key", Value: []byte("test-value")}

		// Set-Get law
		newValue := []byte("new-value")
		updated := lenses.Value.Set(newValue)(attr)
		assert.Equal(t, newValue, lenses.Value.Get(updated))
	})
}

// TestCredentialAttributePrisms verifies prism behavior for CredentialAttribute
func TestCredentialAttributePrisms(t *testing.T) {
	prisms := MakeCredentialAttributePrisms()

	t.Run("Keyword prism returns None for empty string", func(t *testing.T) {
		attr := wincred.CredentialAttribute{Keyword: "", Value: []byte("value")}
		result := prisms.Keyword.GetOption(attr)
		assert.Equal(t, option.None[string](), result)
	})

	t.Run("Keyword prism returns Some for non-empty string", func(t *testing.T) {
		attr := wincred.CredentialAttribute{Keyword: "key", Value: []byte("value")}
		result := prisms.Keyword.GetOption(attr)
		assert.Equal(t, option.Of("key"), result)
	})

	t.Run("Value prism returns Some for non-empty", func(t *testing.T) {
		attr := wincred.CredentialAttribute{Keyword: "key", Value: []byte("value")}
		result := prisms.Value.GetOption(attr)
		assert.Equal(t, option.Of([]byte("value")), result)
	})

	t.Run("Value prism returns None for empty", func(t *testing.T) {
		attrEmpty := wincred.CredentialAttribute{Keyword: "key", Value: nil}
		resultEmpty := prisms.Value.GetOption(attrEmpty)
		assert.Equal(t, option.None[[]byte](), resultEmpty)

		attrEmptySlice := wincred.CredentialAttribute{Keyword: "key", Value: []byte{}}
		resultEmptySlice := prisms.Value.GetOption(attrEmptySlice)
		assert.Equal(t, option.None[[]byte](), resultEmptySlice)
	})

	t.Run("ReverseGet creates struct with field", func(t *testing.T) {
		created := prisms.Keyword.ReverseGet("my-keyword")
		assert.Equal(t, "my-keyword", created.Keyword)
	})
}

// TestGenericCredentialLenses_GetSet verifies lens laws for GenericCredential
func TestGenericCredentialLenses_GetSet(t *testing.T) {
	lenses := MakeGenericCredentialLenses()

	cred := wincred.GenericCredential{
		Credential: wincred.Credential{
			TargetName:     "target",
			Comment:        "comment",
			CredentialBlob: []byte("secret"),
			UserName:       "user",
		},
	}

	t.Run("TargetName lens", func(t *testing.T) {
		// Get-Set law
		got := lenses.TargetName.Get(cred)
		updated := lenses.TargetName.Set(got)(cred)
		assert.Equal(t, cred.TargetName, updated.TargetName)

		// Set-Get law
		newTarget := "new-target"
		updated = lenses.TargetName.Set(newTarget)(cred)
		assert.Equal(t, newTarget, lenses.TargetName.Get(updated))
	})

	t.Run("UserName lens", func(t *testing.T) {
		// Set-Get law
		newUser := "new-user"
		updated := lenses.UserName.Set(newUser)(cred)
		assert.Equal(t, newUser, lenses.UserName.Get(updated))
	})

	t.Run("CredentialBlob lens", func(t *testing.T) {
		// Set-Get law
		newBlob := []byte("new-secret")
		updated := lenses.CredentialBlob.Set(newBlob)(cred)
		assert.Equal(t, newBlob, lenses.CredentialBlob.Get(updated))
	})

	t.Run("LastWritten lens", func(t *testing.T) {
		now := time.Now()
		updated := lenses.LastWritten.Set(now)(cred)
		assert.Equal(t, now, lenses.LastWritten.Get(updated))
	})

	t.Run("Attributes lens", func(t *testing.T) {
		attrs := []wincred.CredentialAttribute{
			{Keyword: "key1", Value: []byte("val1")},
			{Keyword: "key2", Value: []byte("val2")},
		}
		updated := lenses.Attributes.Set(attrs)(cred)
		assert.Equal(t, attrs, lenses.Attributes.Get(updated))
	})
}

// TestGenericCredentialRefLenses_GetSet verifies lens laws for GenericCredential references
func TestGenericCredentialRefLenses_GetSet(t *testing.T) {
	lenses := MakeGenericCredentialRefLenses()

	cred := &wincred.GenericCredential{
		Credential: wincred.Credential{
			TargetName:     "target",
			Comment:        "comment",
			CredentialBlob: []byte("secret"),
			UserName:       "user",
		},
	}

	t.Run("TargetName ref lens", func(t *testing.T) {
		newTarget := "new-target"
		updated := lenses.TargetName.Set(newTarget)(cred)
		assert.Equal(t, newTarget, lenses.TargetName.Get(updated))
	})

	t.Run("CredentialBlob ref lens", func(t *testing.T) {
		newBlob := []byte("new-secret")
		updated := lenses.CredentialBlob.Set(newBlob)(cred)
		assert.Equal(t, newBlob, lenses.CredentialBlob.Get(updated))
	})
}

// TestGenericCredentialPrisms verifies prism behavior for GenericCredential
func TestGenericCredentialPrisms(t *testing.T) {
	prisms := MakeGenericCredentialPrisms()

	t.Run("TargetName prism returns None for empty", func(t *testing.T) {
		cred := wincred.GenericCredential{}
		result := prisms.TargetName.GetOption(cred)
		assert.Equal(t, option.None[string](), result)
	})

	t.Run("TargetName prism returns Some for non-empty", func(t *testing.T) {
		cred := wincred.GenericCredential{
			Credential: wincred.Credential{TargetName: "my-target"},
		}
		result := prisms.TargetName.GetOption(cred)
		assert.Equal(t, option.Of("my-target"), result)
	})

	t.Run("CredentialBlob prism returns Some for non-empty", func(t *testing.T) {
		cred := wincred.GenericCredential{
			Credential: wincred.Credential{CredentialBlob: []byte("secret")},
		}
		result := prisms.CredentialBlob.GetOption(cred)
		assert.Equal(t, option.Of([]byte("secret")), result)
	})

	t.Run("CredentialBlob prism returns None for empty", func(t *testing.T) {
		cred := wincred.GenericCredential{}
		result := prisms.CredentialBlob.GetOption(cred)
		assert.Equal(t, option.None[[]byte](), result)
	})

	t.Run("ReverseGet creates struct with field", func(t *testing.T) {
		created := prisms.UserName.ReverseGet("test-user")
		assert.Equal(t, "test-user", created.UserName)
	})
}

// TestCredentialLenses_GetSet verifies lens laws for Credential
func TestCredentialLenses_GetSet(t *testing.T) {
	lenses := MakeCredentialLenses()

	cred := wincred.Credential{
		TargetName:     "target",
		Comment:        "comment",
		CredentialBlob: []byte("secret"),
		UserName:       "user",
		Persist:        wincred.PersistLocalMachine,
	}

	t.Run("TargetName lens", func(t *testing.T) {
		newTarget := "new-target"
		updated := lenses.TargetName.Set(newTarget)(cred)
		assert.Equal(t, newTarget, lenses.TargetName.Get(updated))
	})

	t.Run("Persist lens", func(t *testing.T) {
		newPersist := wincred.PersistEnterprise
		updated := lenses.Persist.Set(newPersist)(cred)
		assert.Equal(t, newPersist, lenses.Persist.Get(updated))
	})
}

// TestOptionalLenses verifies optional lens behavior
func TestOptionalLenses(t *testing.T) {
	lenses := MakeCredentialLenses()

	t.Run("TargetNameO returns None for empty", func(t *testing.T) {
		cred := wincred.Credential{}
		result := lenses.TargetNameO.Get(cred)
		assert.Equal(t, option.None[string](), result)
	})

	t.Run("TargetNameO returns Some for non-empty", func(t *testing.T) {
		cred := wincred.Credential{TargetName: "target"}
		result := lenses.TargetNameO.Get(cred)
		assert.Equal(t, option.Of("target"), result)
	})

	t.Run("LastWrittenO returns None for zero time", func(t *testing.T) {
		cred := wincred.Credential{}
		result := lenses.LastWrittenO.Get(cred)
		assert.Equal(t, option.None[time.Time](), result)
	})

	t.Run("LastWrittenO returns Some for non-zero time", func(t *testing.T) {
		now := time.Now()
		cred := wincred.Credential{LastWritten: now}
		result := lenses.LastWrittenO.Get(cred)
		assert.Equal(t, option.Of(now), result)
	})
}

// TestLensComposition verifies that lenses can be composed
func TestLensComposition(t *testing.T) {
	credLenses := MakeCredentialLenses()

	// Create a credential and modify multiple fields using functional composition
	cred := function.Pipe1(
		wincred.Credential{},
		function.Flow4(
			credLenses.TargetName.Set("target"),
			credLenses.UserName.Set("user"),
			credLenses.Comment.Set("comment"),
			credLenses.CredentialBlob.Set([]byte("secret")),
		),
	)

	assert.Equal(t, "target", credLenses.TargetName.Get(cred))
	assert.Equal(t, "user", credLenses.UserName.Get(cred))
	assert.Equal(t, "comment", credLenses.Comment.Get(cred))
	assert.Equal(t, []byte("secret"), credLenses.CredentialBlob.Get(cred))
}
