package register

import "testing"

func TestGenerateConfirmCode(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		code, err := GenerateConfirmCode()
		if err != nil {
			t.Errorf("got a error %q", err)
		}
		if len(code) != 36 {
			t.Errorf("got length %d, want length %d", len(code), 36)
		}
	})
}
