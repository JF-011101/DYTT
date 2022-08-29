package code

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ./error_code_generated.md

const (
	// ErrSuccess - 200: OK.
	ErrSuccess int = 0
)

// 通用: 基本错误
// Code must start with 1xxxxx
const (
	// ErrUnknown - 500: Internal server error.
	ErrUnknown int = iota + 100001

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind

	// ErrValidation - 400: Validation failed.
	ErrValidation

	// ErrTokenInvalid - 401: Token invalid.
	ErrTokenInvalid
)

// 通用：数据库类错误
const (
	// ErrDatabase - 500: Database base error.
	ErrDatabase int = iota + 100101

	// ErrRecordNotFound - 500: record not found error
	ErrRecordNotFound

	// ErrInvalidTransaction - 500: invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction

	// ErrNotImplemented - 500: not implemented
	ErrNotImplemented

	// ErrMissingWhereClause - 500: missing where clause
	ErrMissingWhereClause

	// ErrUnsupportedRelation - 500: unsupported relations
	ErrUnsupportedRelation

	// ErrPrimaryKeyRequired - 500: primary keys required
	ErrPrimaryKeyRequired

	// ErrModelValueRequired - 500: model value required
	ErrModelValueRequired

	// ErrInvalidData - 500: unsupported data
	ErrInvalidData

	// ErrUnsupportedDriver - 500: unsupported driver
	ErrUnsupportedDriver

	// ErrRegistered - 500: registered
	ErrRegistered

	// ErrInvalidField - 500: invalid field
	ErrInvalidField

	// ErrEmptySlice - 500: empty slice found
	ErrEmptySlice

	// ErrDryRunModeUnsupported - 500: dry run mode unsupported
	ErrDryRunModeUnsupported

	// ErrInvalidDB - 500: invalid db
	ErrInvalidDB

	// ErrInvalidValue - 500: invalid value
	ErrInvalidValue

	// ErrInvalidValueOfLength - 500: invalid values do not match length
	ErrInvalidValueOfLength

	// ErrPreloadNotAllowed - 500: preload is not allowed when count is used
	ErrPreloadNotAllowed
)

// 通用：认证授权类错误
const (
	// ErrEncrypt - 401: Error occurred while encrypting the user password.
	ErrEncrypt int = iota + 100201

	// ErrSignatureInvalid - 401: Signature is invalid.
	ErrSignatureInvalid

	// ErrExpired - 401: Token expired.
	ErrExpired

	// ErrInvalidAuthHeader - 401: Invalid authorization header.
	ErrInvalidAuthHeader

	// ErrMissingHeader - 401: The `Authorization` header was empty.
	ErrMissingHeader

	// ErrorExpired - 401: Token expired.
	ErrorExpired

	// ErrPasswordIncorrect - 401: Password was incorrect.
	ErrPasswordIncorrect

	// PermissionDenied - 403: Permission denied.
	ErrPermissionDenied
)

// 通用：编解码类错误
const (
	// ErrEncodingFailed - 500: Encoding failed due to an error with the data.
	ErrEncodingFailed int = iota + 100301

	// ErrDecodingFailed - 500: Decoding failed due to an error with the data.
	ErrDecodingFailed

	// ErrInvalidJSON - 500: Data is not valid JSON.
	ErrInvalidJSON

	// ErrEncodingJSON - 500: JSON data could not be encoded.
	ErrEncodingJSON

	// ErrDecodingJSON - 500: JSON data could not be decoded.
	ErrDecodingJSON

	// ErrInvalidYaml - 500: Data is not valid Yaml.
	ErrInvalidYaml

	// ErrEncodingYaml - 500: Yaml data could not be encoded.
	ErrEncodingYaml

	// ErrDecodingYaml - 500: Yaml data could not be decoded.
	ErrDecodingYaml

	// ErrInvalidHash - 500: Encoded hash is not in the correct format.
	ErrInvalidHash

	// ErrIncompatibleVersion - 500: Incompatible version of encryption algorithm.
	ErrIncompatibleVersion
)
