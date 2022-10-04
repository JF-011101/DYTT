/*
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-02 14:03:25
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:35:45
 * @FilePath: \dytt\cmd\user\command\login.go
 * @Description: Login operation business logic
 */

package command

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"

	"github.com/jf-011101/dytt/dal/db"
	"github.com/jf-011101/dytt/grpc_gen/user"
	"github.com/jf-011101/dytt/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *user.DouyinUserRegisterRequest) (int64, error) {
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.ErrUserNotFound
	}
	user := users[0]

	passWordMatch, err := comparePasswordAndHash(req.Password, user.Password)
	if err != nil {
		return 0, err
	}

	if !passWordMatch {
		return 0, errno.ErrPasswordIncorrect
	}
	return int64(user.ID), nil
}

// comparePasswordAndHash compares the password and hash of the given password.
func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	argon2Params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the input password using the same parameters.
	inputHash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, inputHash) == 1 {
		return true, nil
	}
	return false, nil
}

// decodeHash decode the hash of the password from the database.
//
// returns an error if the password is not valid.
func decodeHash(encodedHash string) (argon2Params *Argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errno.ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errno.ErrIncompatibleVersion
	}

	argon2Params = &Argon2Params{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &argon2Params.Memory, &argon2Params.Iterations, &argon2Params.Parallelism); err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	argon2Params.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	argon2Params.KeyLength = uint32(len(hash))

	return argon2Params, salt, hash, nil
}
