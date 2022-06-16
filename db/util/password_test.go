package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPasswprd, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasswprd)

	err = CheckPassword(password, hashedPasswprd)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPasswprd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
