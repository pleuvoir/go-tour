package applog

import "testing"

func TestLogger(t *testing.T) {

	Init("default", "pay")

	Logger("default").Info().Str("default", "default").Int("number", 166).Msg("测试输入")
	Logger("pay").Error().Str("pay", "pay").Msg("测试输入")
}
