// lenses.go provides functional optics (lenses and prisms) for accessing and modifying
// fields of Windows Credential Manager types in a composable, type-safe manner.
//
// This file contains:
//   - Lenses: Bidirectional accessors that can both get and set a field value
//   - Prisms: Partial accessors that may fail to extract a value (returns Option)
//   - Value and Reference variants: Operate on values or pointers respectively
//   - Optional lenses (LensO): Return None for zero values
//
// The optics follow standard functional programming laws:
//   - Lens Laws: Get-Set, Set-Get, Set-Set
//   - Prism Laws: Preview-Review, Review-Preview
//
// Types covered:
//   - [wincred.CredentialAttribute]: Simple key-value attribute
//   - [wincred.Credential]: Base credential with all common fields
//   - [wincred.GenericCredential]: Generic credential (embeds Credential)
package main

import (
	time "time"

	__iso_option "github.com/IBM/fp-go/v2/optics/iso/option"
	__lens "github.com/IBM/fp-go/v2/optics/lens"
	__lens_option "github.com/IBM/fp-go/v2/optics/lens/option"
	__prism "github.com/IBM/fp-go/v2/optics/prism"
	__option "github.com/IBM/fp-go/v2/option"
	"github.com/danieljoos/wincred"
)

// CredentialAttributeLenses provides [lenses] for accessing fields of [wincred.CredentialAttribute]
//
// [lenses]: __lens.Lens
type CredentialAttributeLenses struct {
	// mandatory fields
	Keyword __lens.Lens[wincred.CredentialAttribute, string]
	Value   __lens.Lens[wincred.CredentialAttribute, []byte]
	// optional fields
	KeywordO __lens_option.LensO[wincred.CredentialAttribute, string]
}

// CredentialAttributeRefLenses provides [lenses] for accessing fields of [wincred.CredentialAttribute] via a reference to [wincred.CredentialAttribute]
//
// [lenses]: __lens.Lens
type CredentialAttributeRefLenses struct {
	// mandatory fields
	Keyword __lens.Lens[*wincred.CredentialAttribute, string]
	Value   __lens.Lens[*wincred.CredentialAttribute, []byte]
	// optional fields
	KeywordO __lens_option.LensO[*wincred.CredentialAttribute, string]
}

// CredentialAttributePrisms provides [prisms] for accessing fields of [wincred.CredentialAttribute]
//
// [prisms]: __prism.Prism
type CredentialAttributePrisms struct {
	Keyword __prism.Prism[wincred.CredentialAttribute, string]
	Value   __prism.Prism[wincred.CredentialAttribute, []byte]
}

// CredentialAttributeRefPrisms provides [prisms] for accessing fields of [wincred.CredentialAttribute] via a reference to [wincred.CredentialAttribute]
//
// [prisms]: __prism.Prism
type CredentialAttributeRefPrisms struct {
	Keyword __prism.Prism[*wincred.CredentialAttribute, string]
	Value   __prism.Prism[*wincred.CredentialAttribute, []byte]
}

// MakeCredentialAttributeLenses creates a new [CredentialAttributeLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func MakeCredentialAttributeLenses() CredentialAttributeLenses {
	// mandatory lenses
	lensKeyword := __lens.MakeLensWithName(
		func(s wincred.CredentialAttribute) string { return s.Keyword },
		func(s wincred.CredentialAttribute, v string) wincred.CredentialAttribute { s.Keyword = v; return s },
		"wincred.CredentialAttribute.Keyword",
	)
	lensValue := __lens.MakeLensWithName(
		func(s wincred.CredentialAttribute) []byte { return s.Value },
		func(s wincred.CredentialAttribute, v []byte) wincred.CredentialAttribute { s.Value = v; return s },
		"wincred.CredentialAttribute.Value",
	)
	// optional lenses
	lensKeywordO := __lens_option.FromIso[wincred.CredentialAttribute](__iso_option.FromZero[string]())(lensKeyword)
	return CredentialAttributeLenses{
		// mandatory lenses
		Keyword: lensKeyword,
		Value:   lensValue,
		// optional lenses
		KeywordO: lensKeywordO,
	}
}

// MakeCredentialAttributeRefLenses creates a new [CredentialAttributeRefLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func MakeCredentialAttributeRefLenses() CredentialAttributeRefLenses {
	// mandatory lenses
	lensKeyword := __lens.MakeLensStrictWithName(
		func(s *wincred.CredentialAttribute) string { return s.Keyword },
		func(s *wincred.CredentialAttribute, v string) *wincred.CredentialAttribute { s.Keyword = v; return s },
		"(*wincred.CredentialAttribute).Keyword",
	)
	lensValue := __lens.MakeLensRefWithName(
		func(s *wincred.CredentialAttribute) []byte { return s.Value },
		func(s *wincred.CredentialAttribute, v []byte) *wincred.CredentialAttribute { s.Value = v; return s },
		"(*wincred.CredentialAttribute).Value",
	)
	// optional lenses
	lensKeywordO := __lens_option.FromIso[*wincred.CredentialAttribute](__iso_option.FromZero[string]())(lensKeyword)
	return CredentialAttributeRefLenses{
		// mandatory lenses
		Keyword: lensKeyword,
		Value:   lensValue,
		// optional lenses
		KeywordO: lensKeywordO,
	}
}

// MakeCredentialAttributePrisms creates a new [CredentialAttributePrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func MakeCredentialAttributePrisms() CredentialAttributePrisms {
	_fromNonZeroKeyword := __option.FromNonZero[string]()
	_prismKeyword := __prism.MakePrismWithName(
		func(s wincred.CredentialAttribute) __option.Option[string] { return _fromNonZeroKeyword(s.Keyword) },
		func(v string) wincred.CredentialAttribute {
			return wincred.CredentialAttribute{Keyword: v}
		},
		"wincred.CredentialAttribute.Keyword",
	)
	_prismValue := __prism.MakePrismWithName(
		func(s wincred.CredentialAttribute) __option.Option[[]byte] { return __option.Some(s.Value) },
		func(v []byte) wincred.CredentialAttribute {
			return wincred.CredentialAttribute{Value: v}
		},
		"wincred.CredentialAttribute.Value",
	)
	return CredentialAttributePrisms{
		Keyword: _prismKeyword,
		Value:   _prismValue,
	}
}

// MakeCredentialAttributeRefPrisms creates a new [CredentialAttributeRefPrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func MakeCredentialAttributeRefPrisms() CredentialAttributeRefPrisms {
	_fromNonZeroKeyword := __option.FromNonZero[string]()
	_prismKeyword := __prism.MakePrismWithName(
		func(s *wincred.CredentialAttribute) __option.Option[string] { return _fromNonZeroKeyword(s.Keyword) },
		func(v string) *wincred.CredentialAttribute {
			return &wincred.CredentialAttribute{Keyword: v}
		},
		"wincred.CredentialAttribute.Keyword",
	)
	_prismValue := __prism.MakePrismWithName(
		func(s *wincred.CredentialAttribute) __option.Option[[]byte] { return __option.Some(s.Value) },
		func(v []byte) *wincred.CredentialAttribute {
			return &wincred.CredentialAttribute{Value: v}
		},
		"wincred.CredentialAttribute.Value",
	)
	return CredentialAttributeRefPrisms{
		Keyword: _prismKeyword,
		Value:   _prismValue,
	}
}

// GenericCredentialLenses provides [lenses] for accessing fields of [wincred.GenericCredential]
//
// [lenses]: __lens.Lens
type GenericCredentialLenses struct {
	// mandatory fields
	TargetName     __lens.Lens[wincred.GenericCredential, string]
	Comment        __lens.Lens[wincred.GenericCredential, string]
	LastWritten    __lens.Lens[wincred.GenericCredential, time.Time]
	CredentialBlob __lens.Lens[wincred.GenericCredential, []byte]
	Attributes     __lens.Lens[wincred.GenericCredential, []wincred.CredentialAttribute]
	TargetAlias    __lens.Lens[wincred.GenericCredential, string]
	UserName       __lens.Lens[wincred.GenericCredential, string]
	Persist        __lens.Lens[wincred.GenericCredential, wincred.CredentialPersistence]
	// optional fields
	TargetNameO  __lens_option.LensO[wincred.GenericCredential, string]
	CommentO     __lens_option.LensO[wincred.GenericCredential, string]
	LastWrittenO __lens_option.LensO[wincred.GenericCredential, time.Time]
	TargetAliasO __lens_option.LensO[wincred.GenericCredential, string]
	UserNameO    __lens_option.LensO[wincred.GenericCredential, string]
	PersistO     __lens_option.LensO[wincred.GenericCredential, wincred.CredentialPersistence]
}

// GenericCredentialRefLenses provides [lenses] for accessing fields of [wincred.GenericCredential] via a reference to [wincred.GenericCredential]
//
// [lenses]: __lens.Lens
type GenericCredentialRefLenses struct {
	// mandatory fields
	TargetName     __lens.Lens[*wincred.GenericCredential, string]
	Comment        __lens.Lens[*wincred.GenericCredential, string]
	LastWritten    __lens.Lens[*wincred.GenericCredential, time.Time]
	CredentialBlob __lens.Lens[*wincred.GenericCredential, []byte]
	Attributes     __lens.Lens[*wincred.GenericCredential, []wincred.CredentialAttribute]
	TargetAlias    __lens.Lens[*wincred.GenericCredential, string]
	UserName       __lens.Lens[*wincred.GenericCredential, string]
	Persist        __lens.Lens[*wincred.GenericCredential, wincred.CredentialPersistence]
	// optional fields
	TargetNameO  __lens_option.LensO[*wincred.GenericCredential, string]
	CommentO     __lens_option.LensO[*wincred.GenericCredential, string]
	LastWrittenO __lens_option.LensO[*wincred.GenericCredential, time.Time]
	TargetAliasO __lens_option.LensO[*wincred.GenericCredential, string]
	UserNameO    __lens_option.LensO[*wincred.GenericCredential, string]
	PersistO     __lens_option.LensO[*wincred.GenericCredential, wincred.CredentialPersistence]
}

// GenericCredentialPrisms provides [prisms] for accessing fields of [wincred.GenericCredential]
//
// [prisms]: __prism.Prism
type GenericCredentialPrisms struct {
	TargetName     __prism.Prism[wincred.GenericCredential, string]
	Comment        __prism.Prism[wincred.GenericCredential, string]
	LastWritten    __prism.Prism[wincred.GenericCredential, time.Time]
	CredentialBlob __prism.Prism[wincred.GenericCredential, []byte]
	Attributes     __prism.Prism[wincred.GenericCredential, []wincred.CredentialAttribute]
	TargetAlias    __prism.Prism[wincred.GenericCredential, string]
	UserName       __prism.Prism[wincred.GenericCredential, string]
	Persist        __prism.Prism[wincred.GenericCredential, wincred.CredentialPersistence]
}

// GenericCredentialRefPrisms provides [prisms] for accessing fields of [wincred.GenericCredential] via a reference to [wincred.GenericCredential]
//
// [prisms]: __prism.Prism
type GenericCredentialRefPrisms struct {
	TargetName     __prism.Prism[*wincred.GenericCredential, string]
	Comment        __prism.Prism[*wincred.GenericCredential, string]
	LastWritten    __prism.Prism[*wincred.GenericCredential, time.Time]
	CredentialBlob __prism.Prism[*wincred.GenericCredential, []byte]
	Attributes     __prism.Prism[*wincred.GenericCredential, []wincred.CredentialAttribute]
	TargetAlias    __prism.Prism[*wincred.GenericCredential, string]
	UserName       __prism.Prism[*wincred.GenericCredential, string]
	Persist        __prism.Prism[*wincred.GenericCredential, wincred.CredentialPersistence]
}

// MakeGenericCredentialLenses creates a new [GenericCredentialLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func MakeGenericCredentialLenses() GenericCredentialLenses {
	// mandatory lenses
	lensTargetName := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) string { return s.TargetName },
		func(s wincred.GenericCredential, v string) wincred.GenericCredential { s.TargetName = v; return s },
		"wincred.GenericCredential.TargetName",
	)
	lensComment := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) string { return s.Comment },
		func(s wincred.GenericCredential, v string) wincred.GenericCredential { s.Comment = v; return s },
		"wincred.GenericCredential.Comment",
	)
	lensLastWritten := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) time.Time { return s.LastWritten },
		func(s wincred.GenericCredential, v time.Time) wincred.GenericCredential { s.LastWritten = v; return s },
		"wincred.GenericCredential.LastWritten",
	)
	lensCredentialBlob := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) []byte { return s.CredentialBlob },
		func(s wincred.GenericCredential, v []byte) wincred.GenericCredential { s.CredentialBlob = v; return s },
		"wincred.GenericCredential.CredentialBlob",
	)
	lensAttributes := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) []wincred.CredentialAttribute { return s.Attributes },
		func(s wincred.GenericCredential, v []wincred.CredentialAttribute) wincred.GenericCredential {
			s.Attributes = v
			return s
		},
		"wincred.GenericCredential.Attributes",
	)
	lensTargetAlias := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) string { return s.TargetAlias },
		func(s wincred.GenericCredential, v string) wincred.GenericCredential { s.TargetAlias = v; return s },
		"wincred.GenericCredential.TargetAlias",
	)
	lensUserName := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) string { return s.UserName },
		func(s wincred.GenericCredential, v string) wincred.GenericCredential { s.UserName = v; return s },
		"wincred.GenericCredential.UserName",
	)
	lensPersist := __lens.MakeLensWithName(
		func(s wincred.GenericCredential) wincred.CredentialPersistence { return s.Persist },
		func(s wincred.GenericCredential, v wincred.CredentialPersistence) wincred.GenericCredential {
			s.Persist = v
			return s
		},
		"wincred.GenericCredential.Persist",
	)
	// optional lenses
	lensTargetNameO := __lens_option.FromIso[wincred.GenericCredential](__iso_option.FromZero[string]())(lensTargetName)
	lensCommentO := __lens_option.FromIso[wincred.GenericCredential](__iso_option.FromZero[string]())(lensComment)
	lensLastWrittenO := __lens_option.FromIso[wincred.GenericCredential](__iso_option.FromZero[time.Time]())(lensLastWritten)
	lensTargetAliasO := __lens_option.FromIso[wincred.GenericCredential](__iso_option.FromZero[string]())(lensTargetAlias)
	lensUserNameO := __lens_option.FromIso[wincred.GenericCredential](__iso_option.FromZero[string]())(lensUserName)
	lensPersistO := __lens_option.FromIso[wincred.GenericCredential](__iso_option.FromZero[wincred.CredentialPersistence]())(lensPersist)
	return GenericCredentialLenses{
		// mandatory lenses
		TargetName:     lensTargetName,
		Comment:        lensComment,
		LastWritten:    lensLastWritten,
		CredentialBlob: lensCredentialBlob,
		Attributes:     lensAttributes,
		TargetAlias:    lensTargetAlias,
		UserName:       lensUserName,
		Persist:        lensPersist,
		// optional lenses
		TargetNameO:  lensTargetNameO,
		CommentO:     lensCommentO,
		LastWrittenO: lensLastWrittenO,
		TargetAliasO: lensTargetAliasO,
		UserNameO:    lensUserNameO,
		PersistO:     lensPersistO,
	}
}

// MakeGenericCredentialRefLenses creates a new [GenericCredentialRefLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func MakeGenericCredentialRefLenses() GenericCredentialRefLenses {
	// mandatory lenses
	lensTargetName := __lens.MakeLensStrictWithName(
		func(s *wincred.GenericCredential) string { return s.TargetName },
		func(s *wincred.GenericCredential, v string) *wincred.GenericCredential { s.TargetName = v; return s },
		"(*wincred.GenericCredential).TargetName",
	)
	lensComment := __lens.MakeLensStrictWithName(
		func(s *wincred.GenericCredential) string { return s.Comment },
		func(s *wincred.GenericCredential, v string) *wincred.GenericCredential { s.Comment = v; return s },
		"(*wincred.GenericCredential).Comment",
	)
	lensLastWritten := __lens.MakeLensStrictWithName(
		func(s *wincred.GenericCredential) time.Time { return s.LastWritten },
		func(s *wincred.GenericCredential, v time.Time) *wincred.GenericCredential {
			s.LastWritten = v
			return s
		},
		"(*wincred.GenericCredential).LastWritten",
	)
	lensCredentialBlob := __lens.MakeLensRefWithName(
		func(s *wincred.GenericCredential) []byte { return s.CredentialBlob },
		func(s *wincred.GenericCredential, v []byte) *wincred.GenericCredential {
			s.CredentialBlob = v
			return s
		},
		"(*wincred.GenericCredential).CredentialBlob",
	)
	lensAttributes := __lens.MakeLensRefWithName(
		func(s *wincred.GenericCredential) []wincred.CredentialAttribute { return s.Attributes },
		func(s *wincred.GenericCredential, v []wincred.CredentialAttribute) *wincred.GenericCredential {
			s.Attributes = v
			return s
		},
		"(*wincred.GenericCredential).Attributes",
	)
	lensTargetAlias := __lens.MakeLensStrictWithName(
		func(s *wincred.GenericCredential) string { return s.TargetAlias },
		func(s *wincred.GenericCredential, v string) *wincred.GenericCredential { s.TargetAlias = v; return s },
		"(*wincred.GenericCredential).TargetAlias",
	)
	lensUserName := __lens.MakeLensStrictWithName(
		func(s *wincred.GenericCredential) string { return s.UserName },
		func(s *wincred.GenericCredential, v string) *wincred.GenericCredential { s.UserName = v; return s },
		"(*wincred.GenericCredential).UserName",
	)
	lensPersist := __lens.MakeLensStrictWithName(
		func(s *wincred.GenericCredential) wincred.CredentialPersistence { return s.Persist },
		func(s *wincred.GenericCredential, v wincred.CredentialPersistence) *wincred.GenericCredential {
			s.Persist = v
			return s
		},
		"(*wincred.GenericCredential).Persist",
	)
	// optional lenses
	lensTargetNameO := __lens_option.FromIso[*wincred.GenericCredential](__iso_option.FromZero[string]())(lensTargetName)
	lensCommentO := __lens_option.FromIso[*wincred.GenericCredential](__iso_option.FromZero[string]())(lensComment)
	lensLastWrittenO := __lens_option.FromIso[*wincred.GenericCredential](__iso_option.FromZero[time.Time]())(lensLastWritten)
	lensTargetAliasO := __lens_option.FromIso[*wincred.GenericCredential](__iso_option.FromZero[string]())(lensTargetAlias)
	lensUserNameO := __lens_option.FromIso[*wincred.GenericCredential](__iso_option.FromZero[string]())(lensUserName)
	lensPersistO := __lens_option.FromIso[*wincred.GenericCredential](__iso_option.FromZero[wincred.CredentialPersistence]())(lensPersist)
	return GenericCredentialRefLenses{
		// mandatory lenses
		TargetName:     lensTargetName,
		Comment:        lensComment,
		LastWritten:    lensLastWritten,
		CredentialBlob: lensCredentialBlob,
		Attributes:     lensAttributes,
		TargetAlias:    lensTargetAlias,
		UserName:       lensUserName,
		Persist:        lensPersist,
		// optional lenses
		TargetNameO:  lensTargetNameO,
		CommentO:     lensCommentO,
		LastWrittenO: lensLastWrittenO,
		TargetAliasO: lensTargetAliasO,
		UserNameO:    lensUserNameO,
		PersistO:     lensPersistO,
	}
}

// MakeGenericCredentialPrisms creates a new [GenericCredentialPrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func MakeGenericCredentialPrisms() GenericCredentialPrisms {
	_fromNonZeroTargetName := __option.FromNonZero[string]()
	_prismTargetName := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[string] { return _fromNonZeroTargetName(s.TargetName) },
		func(v string) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{TargetName: v}}
		},
		"wincred.GenericCredential.TargetName",
	)
	_fromNonZeroComment := __option.FromNonZero[string]()
	_prismComment := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[string] { return _fromNonZeroComment(s.Comment) },
		func(v string) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{Comment: v}}
		},
		"wincred.GenericCredential.Comment",
	)
	_fromNonZeroLastWritten := __option.FromNonZero[time.Time]()
	_prismLastWritten := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[time.Time] {
			return _fromNonZeroLastWritten(s.LastWritten)
		},
		func(v time.Time) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{LastWritten: v}}
		},
		"wincred.GenericCredential.LastWritten",
	)
	_prismCredentialBlob := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[[]byte] { return __option.Some(s.CredentialBlob) },
		func(v []byte) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{CredentialBlob: v}}
		},
		"wincred.GenericCredential.CredentialBlob",
	)
	_prismAttributes := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[[]wincred.CredentialAttribute] {
			return __option.Some(s.Attributes)
		},
		func(v []wincred.CredentialAttribute) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{Attributes: v}}
		},
		"wincred.GenericCredential.Attributes",
	)
	_fromNonZeroTargetAlias := __option.FromNonZero[string]()
	_prismTargetAlias := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[string] {
			return _fromNonZeroTargetAlias(s.TargetAlias)
		},
		func(v string) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{TargetAlias: v}}
		},
		"wincred.GenericCredential.TargetAlias",
	)
	_fromNonZeroUserName := __option.FromNonZero[string]()
	_prismUserName := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[string] { return _fromNonZeroUserName(s.UserName) },
		func(v string) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{UserName: v}}
		},
		"wincred.GenericCredential.UserName",
	)
	_fromNonZeroPersist := __option.FromNonZero[wincred.CredentialPersistence]()
	_prismPersist := __prism.MakePrismWithName(
		func(s wincred.GenericCredential) __option.Option[wincred.CredentialPersistence] {
			return _fromNonZeroPersist(s.Persist)
		},
		func(v wincred.CredentialPersistence) wincred.GenericCredential {
			return wincred.GenericCredential{Credential: wincred.Credential{Persist: v}}
		},
		"wincred.GenericCredential.Persist",
	)
	return GenericCredentialPrisms{
		TargetName:     _prismTargetName,
		Comment:        _prismComment,
		LastWritten:    _prismLastWritten,
		CredentialBlob: _prismCredentialBlob,
		Attributes:     _prismAttributes,
		TargetAlias:    _prismTargetAlias,
		UserName:       _prismUserName,
		Persist:        _prismPersist,
	}
}

// MakeGenericCredentialRefPrisms creates a new [GenericCredentialRefPrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func MakeGenericCredentialRefPrisms() GenericCredentialRefPrisms {
	_fromNonZeroTargetName := __option.FromNonZero[string]()
	_prismTargetName := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[string] {
			return _fromNonZeroTargetName(s.TargetName)
		},
		func(v string) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{TargetName: v}}
		},
		"wincred.GenericCredential.TargetName",
	)
	_fromNonZeroComment := __option.FromNonZero[string]()
	_prismComment := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[string] { return _fromNonZeroComment(s.Comment) },
		func(v string) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{Comment: v}}
		},
		"wincred.GenericCredential.Comment",
	)
	_fromNonZeroLastWritten := __option.FromNonZero[time.Time]()
	_prismLastWritten := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[time.Time] {
			return _fromNonZeroLastWritten(s.LastWritten)
		},
		func(v time.Time) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{LastWritten: v}}
		},
		"wincred.GenericCredential.LastWritten",
	)
	_prismCredentialBlob := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[[]byte] { return __option.Some(s.CredentialBlob) },
		func(v []byte) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{CredentialBlob: v}}
		},
		"wincred.GenericCredential.CredentialBlob",
	)
	_prismAttributes := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[[]wincred.CredentialAttribute] {
			return __option.Some(s.Attributes)
		},
		func(v []wincred.CredentialAttribute) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{Attributes: v}}
		},
		"wincred.GenericCredential.Attributes",
	)
	_fromNonZeroTargetAlias := __option.FromNonZero[string]()
	_prismTargetAlias := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[string] {
			return _fromNonZeroTargetAlias(s.TargetAlias)
		},
		func(v string) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{TargetAlias: v}}
		},
		"wincred.GenericCredential.TargetAlias",
	)
	_fromNonZeroUserName := __option.FromNonZero[string]()
	_prismUserName := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[string] { return _fromNonZeroUserName(s.UserName) },
		func(v string) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{UserName: v}}
		},
		"wincred.GenericCredential.UserName",
	)
	_fromNonZeroPersist := __option.FromNonZero[wincred.CredentialPersistence]()
	_prismPersist := __prism.MakePrismWithName(
		func(s *wincred.GenericCredential) __option.Option[wincred.CredentialPersistence] {
			return _fromNonZeroPersist(s.Persist)
		},
		func(v wincred.CredentialPersistence) *wincred.GenericCredential {
			return &wincred.GenericCredential{Credential: wincred.Credential{Persist: v}}
		},
		"wincred.GenericCredential.Persist",
	)
	return GenericCredentialRefPrisms{
		TargetName:     _prismTargetName,
		Comment:        _prismComment,
		LastWritten:    _prismLastWritten,
		CredentialBlob: _prismCredentialBlob,
		Attributes:     _prismAttributes,
		TargetAlias:    _prismTargetAlias,
		UserName:       _prismUserName,
		Persist:        _prismPersist,
	}
}

// CredentialLenses provides [lenses] for accessing fields of [wincred.Credential]
//
// [lenses]: __lens.Lens
type CredentialLenses struct {
	// mandatory fields
	TargetName     __lens.Lens[wincred.Credential, string]
	Comment        __lens.Lens[wincred.Credential, string]
	LastWritten    __lens.Lens[wincred.Credential, time.Time]
	CredentialBlob __lens.Lens[wincred.Credential, []byte]
	Attributes     __lens.Lens[wincred.Credential, []wincred.CredentialAttribute]
	TargetAlias    __lens.Lens[wincred.Credential, string]
	UserName       __lens.Lens[wincred.Credential, string]
	Persist        __lens.Lens[wincred.Credential, wincred.CredentialPersistence]
	// optional fields
	TargetNameO  __lens_option.LensO[wincred.Credential, string]
	CommentO     __lens_option.LensO[wincred.Credential, string]
	LastWrittenO __lens_option.LensO[wincred.Credential, time.Time]
	TargetAliasO __lens_option.LensO[wincred.Credential, string]
	UserNameO    __lens_option.LensO[wincred.Credential, string]
	PersistO     __lens_option.LensO[wincred.Credential, wincred.CredentialPersistence]
}

// CredentialRefLenses provides [lenses] for accessing fields of [wincred.Credential] via a reference to [wincred.Credential]
//
// [lenses]: __lens.Lens
type CredentialRefLenses struct {
	// mandatory fields
	TargetName     __lens.Lens[*wincred.Credential, string]
	Comment        __lens.Lens[*wincred.Credential, string]
	LastWritten    __lens.Lens[*wincred.Credential, time.Time]
	CredentialBlob __lens.Lens[*wincred.Credential, []byte]
	Attributes     __lens.Lens[*wincred.Credential, []wincred.CredentialAttribute]
	TargetAlias    __lens.Lens[*wincred.Credential, string]
	UserName       __lens.Lens[*wincred.Credential, string]
	Persist        __lens.Lens[*wincred.Credential, wincred.CredentialPersistence]
	// optional fields
	TargetNameO  __lens_option.LensO[*wincred.Credential, string]
	CommentO     __lens_option.LensO[*wincred.Credential, string]
	LastWrittenO __lens_option.LensO[*wincred.Credential, time.Time]
	TargetAliasO __lens_option.LensO[*wincred.Credential, string]
	UserNameO    __lens_option.LensO[*wincred.Credential, string]
	PersistO     __lens_option.LensO[*wincred.Credential, wincred.CredentialPersistence]
}

// CredentialPrisms provides [prisms] for accessing fields of [wincred.Credential]
//
// [prisms]: __prism.Prism
type CredentialPrisms struct {
	TargetName     __prism.Prism[wincred.Credential, string]
	Comment        __prism.Prism[wincred.Credential, string]
	LastWritten    __prism.Prism[wincred.Credential, time.Time]
	CredentialBlob __prism.Prism[wincred.Credential, []byte]
	Attributes     __prism.Prism[wincred.Credential, []wincred.CredentialAttribute]
	TargetAlias    __prism.Prism[wincred.Credential, string]
	UserName       __prism.Prism[wincred.Credential, string]
	Persist        __prism.Prism[wincred.Credential, wincred.CredentialPersistence]
}

// CredentialRefPrisms provides [prisms] for accessing fields of [wincred.Credential] via a reference to [wincred.Credential]
//
// [prisms]: __prism.Prism
type CredentialRefPrisms struct {
	TargetName     __prism.Prism[*wincred.Credential, string]
	Comment        __prism.Prism[*wincred.Credential, string]
	LastWritten    __prism.Prism[*wincred.Credential, time.Time]
	CredentialBlob __prism.Prism[*wincred.Credential, []byte]
	Attributes     __prism.Prism[*wincred.Credential, []wincred.CredentialAttribute]
	TargetAlias    __prism.Prism[*wincred.Credential, string]
	UserName       __prism.Prism[*wincred.Credential, string]
	Persist        __prism.Prism[*wincred.Credential, wincred.CredentialPersistence]
}

// MakeCredentialLenses creates a new [CredentialLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func MakeCredentialLenses() CredentialLenses {
	// mandatory lenses
	lensTargetName := __lens.MakeLensWithName(
		func(s wincred.Credential) string { return s.TargetName },
		func(s wincred.Credential, v string) wincred.Credential { s.TargetName = v; return s },
		"wincred.Credential.TargetName",
	)
	lensComment := __lens.MakeLensWithName(
		func(s wincred.Credential) string { return s.Comment },
		func(s wincred.Credential, v string) wincred.Credential { s.Comment = v; return s },
		"wincred.Credential.Comment",
	)
	lensLastWritten := __lens.MakeLensWithName(
		func(s wincred.Credential) time.Time { return s.LastWritten },
		func(s wincred.Credential, v time.Time) wincred.Credential { s.LastWritten = v; return s },
		"wincred.Credential.LastWritten",
	)
	lensCredentialBlob := __lens.MakeLensWithName(
		func(s wincred.Credential) []byte { return s.CredentialBlob },
		func(s wincred.Credential, v []byte) wincred.Credential { s.CredentialBlob = v; return s },
		"wincred.Credential.CredentialBlob",
	)
	lensAttributes := __lens.MakeLensWithName(
		func(s wincred.Credential) []wincred.CredentialAttribute { return s.Attributes },
		func(s wincred.Credential, v []wincred.CredentialAttribute) wincred.Credential {
			s.Attributes = v
			return s
		},
		"wincred.Credential.Attributes",
	)
	lensTargetAlias := __lens.MakeLensWithName(
		func(s wincred.Credential) string { return s.TargetAlias },
		func(s wincred.Credential, v string) wincred.Credential { s.TargetAlias = v; return s },
		"wincred.Credential.TargetAlias",
	)
	lensUserName := __lens.MakeLensWithName(
		func(s wincred.Credential) string { return s.UserName },
		func(s wincred.Credential, v string) wincred.Credential { s.UserName = v; return s },
		"wincred.Credential.UserName",
	)
	lensPersist := __lens.MakeLensWithName(
		func(s wincred.Credential) wincred.CredentialPersistence { return s.Persist },
		func(s wincred.Credential, v wincred.CredentialPersistence) wincred.Credential {
			s.Persist = v
			return s
		},
		"wincred.Credential.Persist",
	)
	// optional lenses
	lensTargetNameO := __lens_option.FromIso[wincred.Credential](__iso_option.FromZero[string]())(lensTargetName)
	lensCommentO := __lens_option.FromIso[wincred.Credential](__iso_option.FromZero[string]())(lensComment)
	lensLastWrittenO := __lens_option.FromIso[wincred.Credential](__iso_option.FromZero[time.Time]())(lensLastWritten)
	lensTargetAliasO := __lens_option.FromIso[wincred.Credential](__iso_option.FromZero[string]())(lensTargetAlias)
	lensUserNameO := __lens_option.FromIso[wincred.Credential](__iso_option.FromZero[string]())(lensUserName)
	lensPersistO := __lens_option.FromIso[wincred.Credential](__iso_option.FromZero[wincred.CredentialPersistence]())(lensPersist)
	return CredentialLenses{
		// mandatory lenses
		TargetName:     lensTargetName,
		Comment:        lensComment,
		LastWritten:    lensLastWritten,
		CredentialBlob: lensCredentialBlob,
		Attributes:     lensAttributes,
		TargetAlias:    lensTargetAlias,
		UserName:       lensUserName,
		Persist:        lensPersist,
		// optional lenses
		TargetNameO:  lensTargetNameO,
		CommentO:     lensCommentO,
		LastWrittenO: lensLastWrittenO,
		TargetAliasO: lensTargetAliasO,
		UserNameO:    lensUserNameO,
		PersistO:     lensPersistO,
	}
}

// MakeCredentialRefLenses creates a new [CredentialRefLenses] with [lenses] for all fields
//
// [lenses]:__lens.Lens
func MakeCredentialRefLenses() CredentialRefLenses {
	// mandatory lenses
	lensTargetName := __lens.MakeLensStrictWithName(
		func(s *wincred.Credential) string { return s.TargetName },
		func(s *wincred.Credential, v string) *wincred.Credential { s.TargetName = v; return s },
		"(*wincred.Credential).TargetName",
	)
	lensComment := __lens.MakeLensStrictWithName(
		func(s *wincred.Credential) string { return s.Comment },
		func(s *wincred.Credential, v string) *wincred.Credential { s.Comment = v; return s },
		"(*wincred.Credential).Comment",
	)
	lensLastWritten := __lens.MakeLensStrictWithName(
		func(s *wincred.Credential) time.Time { return s.LastWritten },
		func(s *wincred.Credential, v time.Time) *wincred.Credential { s.LastWritten = v; return s },
		"(*wincred.Credential).LastWritten",
	)
	lensCredentialBlob := __lens.MakeLensRefWithName(
		func(s *wincred.Credential) []byte { return s.CredentialBlob },
		func(s *wincred.Credential, v []byte) *wincred.Credential { s.CredentialBlob = v; return s },
		"(*wincred.Credential).CredentialBlob",
	)
	lensAttributes := __lens.MakeLensRefWithName(
		func(s *wincred.Credential) []wincred.CredentialAttribute { return s.Attributes },
		func(s *wincred.Credential, v []wincred.CredentialAttribute) *wincred.Credential {
			s.Attributes = v
			return s
		},
		"(*wincred.Credential).Attributes",
	)
	lensTargetAlias := __lens.MakeLensStrictWithName(
		func(s *wincred.Credential) string { return s.TargetAlias },
		func(s *wincred.Credential, v string) *wincred.Credential { s.TargetAlias = v; return s },
		"(*wincred.Credential).TargetAlias",
	)
	lensUserName := __lens.MakeLensStrictWithName(
		func(s *wincred.Credential) string { return s.UserName },
		func(s *wincred.Credential, v string) *wincred.Credential { s.UserName = v; return s },
		"(*wincred.Credential).UserName",
	)
	lensPersist := __lens.MakeLensStrictWithName(
		func(s *wincred.Credential) wincred.CredentialPersistence { return s.Persist },
		func(s *wincred.Credential, v wincred.CredentialPersistence) *wincred.Credential {
			s.Persist = v
			return s
		},
		"(*wincred.Credential).Persist",
	)
	// optional lenses
	lensTargetNameO := __lens_option.FromIso[*wincred.Credential](__iso_option.FromZero[string]())(lensTargetName)
	lensCommentO := __lens_option.FromIso[*wincred.Credential](__iso_option.FromZero[string]())(lensComment)
	lensLastWrittenO := __lens_option.FromIso[*wincred.Credential](__iso_option.FromZero[time.Time]())(lensLastWritten)
	lensTargetAliasO := __lens_option.FromIso[*wincred.Credential](__iso_option.FromZero[string]())(lensTargetAlias)
	lensUserNameO := __lens_option.FromIso[*wincred.Credential](__iso_option.FromZero[string]())(lensUserName)
	lensPersistO := __lens_option.FromIso[*wincred.Credential](__iso_option.FromZero[wincred.CredentialPersistence]())(lensPersist)
	return CredentialRefLenses{
		// mandatory lenses
		TargetName:     lensTargetName,
		Comment:        lensComment,
		LastWritten:    lensLastWritten,
		CredentialBlob: lensCredentialBlob,
		Attributes:     lensAttributes,
		TargetAlias:    lensTargetAlias,
		UserName:       lensUserName,
		Persist:        lensPersist,
		// optional lenses
		TargetNameO:  lensTargetNameO,
		CommentO:     lensCommentO,
		LastWrittenO: lensLastWrittenO,
		TargetAliasO: lensTargetAliasO,
		UserNameO:    lensUserNameO,
		PersistO:     lensPersistO,
	}
}

// MakeCredentialPrisms creates a new [CredentialPrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func MakeCredentialPrisms() CredentialPrisms {
	_fromNonZeroTargetName := __option.FromNonZero[string]()
	_prismTargetName := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[string] { return _fromNonZeroTargetName(s.TargetName) },
		func(v string) wincred.Credential {
			return wincred.Credential{TargetName: v}
		},
		"wincred.Credential.TargetName",
	)
	_fromNonZeroComment := __option.FromNonZero[string]()
	_prismComment := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[string] { return _fromNonZeroComment(s.Comment) },
		func(v string) wincred.Credential {
			return wincred.Credential{Comment: v}
		},
		"wincred.Credential.Comment",
	)
	_fromNonZeroLastWritten := __option.FromNonZero[time.Time]()
	_prismLastWritten := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[time.Time] { return _fromNonZeroLastWritten(s.LastWritten) },
		func(v time.Time) wincred.Credential {
			return wincred.Credential{LastWritten: v}
		},
		"wincred.Credential.LastWritten",
	)
	_prismCredentialBlob := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[[]byte] { return __option.Some(s.CredentialBlob) },
		func(v []byte) wincred.Credential {
			return wincred.Credential{CredentialBlob: v}
		},
		"wincred.Credential.CredentialBlob",
	)
	_prismAttributes := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[[]wincred.CredentialAttribute] {
			return __option.Some(s.Attributes)
		},
		func(v []wincred.CredentialAttribute) wincred.Credential {
			return wincred.Credential{Attributes: v}
		},
		"wincred.Credential.Attributes",
	)
	_fromNonZeroTargetAlias := __option.FromNonZero[string]()
	_prismTargetAlias := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[string] { return _fromNonZeroTargetAlias(s.TargetAlias) },
		func(v string) wincred.Credential {
			return wincred.Credential{TargetAlias: v}
		},
		"wincred.Credential.TargetAlias",
	)
	_fromNonZeroUserName := __option.FromNonZero[string]()
	_prismUserName := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[string] { return _fromNonZeroUserName(s.UserName) },
		func(v string) wincred.Credential {
			return wincred.Credential{UserName: v}
		},
		"wincred.Credential.UserName",
	)
	_fromNonZeroPersist := __option.FromNonZero[wincred.CredentialPersistence]()
	_prismPersist := __prism.MakePrismWithName(
		func(s wincred.Credential) __option.Option[wincred.CredentialPersistence] {
			return _fromNonZeroPersist(s.Persist)
		},
		func(v wincred.CredentialPersistence) wincred.Credential {
			return wincred.Credential{Persist: v}
		},
		"wincred.Credential.Persist",
	)
	return CredentialPrisms{
		TargetName:     _prismTargetName,
		Comment:        _prismComment,
		LastWritten:    _prismLastWritten,
		CredentialBlob: _prismCredentialBlob,
		Attributes:     _prismAttributes,
		TargetAlias:    _prismTargetAlias,
		UserName:       _prismUserName,
		Persist:        _prismPersist,
	}
}

// MakeCredentialRefPrisms creates a new [CredentialRefPrisms] with [prisms] for all fields
//
// [prisms]:__prism.Prism
func MakeCredentialRefPrisms() CredentialRefPrisms {
	_fromNonZeroTargetName := __option.FromNonZero[string]()
	_prismTargetName := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[string] { return _fromNonZeroTargetName(s.TargetName) },
		func(v string) *wincred.Credential {
			return &wincred.Credential{TargetName: v}
		},
		"wincred.Credential.TargetName",
	)
	_fromNonZeroComment := __option.FromNonZero[string]()
	_prismComment := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[string] { return _fromNonZeroComment(s.Comment) },
		func(v string) *wincred.Credential {
			return &wincred.Credential{Comment: v}
		},
		"wincred.Credential.Comment",
	)
	_fromNonZeroLastWritten := __option.FromNonZero[time.Time]()
	_prismLastWritten := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[time.Time] { return _fromNonZeroLastWritten(s.LastWritten) },
		func(v time.Time) *wincred.Credential {
			return &wincred.Credential{LastWritten: v}
		},
		"wincred.Credential.LastWritten",
	)
	_prismCredentialBlob := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[[]byte] { return __option.Some(s.CredentialBlob) },
		func(v []byte) *wincred.Credential {
			return &wincred.Credential{CredentialBlob: v}
		},
		"wincred.Credential.CredentialBlob",
	)
	_prismAttributes := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[[]wincred.CredentialAttribute] {
			return __option.Some(s.Attributes)
		},
		func(v []wincred.CredentialAttribute) *wincred.Credential {
			return &wincred.Credential{Attributes: v}
		},
		"wincred.Credential.Attributes",
	)
	_fromNonZeroTargetAlias := __option.FromNonZero[string]()
	_prismTargetAlias := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[string] { return _fromNonZeroTargetAlias(s.TargetAlias) },
		func(v string) *wincred.Credential {
			return &wincred.Credential{TargetAlias: v}
		},
		"wincred.Credential.TargetAlias",
	)
	_fromNonZeroUserName := __option.FromNonZero[string]()
	_prismUserName := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[string] { return _fromNonZeroUserName(s.UserName) },
		func(v string) *wincred.Credential {
			return &wincred.Credential{UserName: v}
		},
		"wincred.Credential.UserName",
	)
	_fromNonZeroPersist := __option.FromNonZero[wincred.CredentialPersistence]()
	_prismPersist := __prism.MakePrismWithName(
		func(s *wincred.Credential) __option.Option[wincred.CredentialPersistence] {
			return _fromNonZeroPersist(s.Persist)
		},
		func(v wincred.CredentialPersistence) *wincred.Credential {
			return &wincred.Credential{Persist: v}
		},
		"wincred.Credential.Persist",
	)
	return CredentialRefPrisms{
		TargetName:     _prismTargetName,
		Comment:        _prismComment,
		LastWritten:    _prismLastWritten,
		CredentialBlob: _prismCredentialBlob,
		Attributes:     _prismAttributes,
		TargetAlias:    _prismTargetAlias,
		UserName:       _prismUserName,
		Persist:        _prismPersist,
	}
}
