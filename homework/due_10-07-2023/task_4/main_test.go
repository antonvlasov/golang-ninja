package main

import (
	"testing"
    "strings"
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
    var builder strings.Builder
    for idx, val := range password{
        if idx % 3 == 0{
            builder.WriteByte('_')
        }else{
            builder.WriteRune(val)
        }
    }
    return builder.String()
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
