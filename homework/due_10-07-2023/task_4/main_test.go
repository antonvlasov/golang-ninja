package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	pass           = "adminSecretP@ss1!AndAlsoThisPassIsQuiteLong"
	expectedMasked = "_dm_nS_cr_tP_ss_!A_dA_so_hi_Pa_sI_Qu_te_on_"
)

func MaskPasswordBad(password string) string {
	for i := 0; i < len(password); i += 3 {
		password = password[:i] + "_" + password[i+1:]
	}

	return password
}

func MaskPasswordGood(password string) string {
	s := make([]string, 0, len(password))

	for i := 0; i < len(password); i++ {
		if i == 0 || i%3 == 0 {
			s = append(s, "_")
		} else {
			s = append(s, password[i:i+1])
		}
	}
	return strings.Join(s, "")
}

func BenchmarkMaskPassword(b *testing.B) {
	b.Run("bad", func(b *testing.B) {
		var masked string
		for i := 0; i < b.N; i++ {
			masked = MaskPasswordBad(pass)
		}

		_ = masked
	})

	b.Run("good", func(b *testing.B) {
		var masked string
		for i := 0; i < b.N; i++ {
			masked = MaskPasswordGood(pass)
		}

		_ = masked
	})

	b.Run("validate", func(b *testing.B) {
		masked := MaskPasswordBad(pass)
		require.Equal(b, expectedMasked, masked)

		masked = MaskPasswordGood(pass)
		require.Equal(b, expectedMasked, masked)
	})
}
