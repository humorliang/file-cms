package setting

import "testing"

func TestConfigInit(t *testing.T) {
	//
	mode := "dev"
	ConfigInit(&mode)
}
