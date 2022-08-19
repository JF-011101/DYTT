package code

// 服务: 用户类错误
const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 110001

	// ErrUserAlreadyExist - 400: User already exist.
	ErrUserAlreadyExist
)

// 服务: 密钥类错误
const (
	// ErrEncrypt - 400: Secret reach the max count.
	ErrReachMaxCount int = iota + 110101

	// ErrSecretNotFound - 404: Secret not found.
	ErrSecretNotFound
)

// 服务: 数据库类错误
const (
	// ErrVideoNotFound - 400: Video not found.
	ErrVideoNotFound int = iota + 120001

	// ErrCommentNotFound - 400: Comment not found.
	ErrCommentNotFound

	// ErrRelationNotFound - 400: Relation not found.
	ErrRelationNotFound

	// ErrFavoriteNotFound - 400: Favorite not found.
	ErrFavoriteNotFound
)
