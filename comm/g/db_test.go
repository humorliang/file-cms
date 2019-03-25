package g

import "testing"

func TestInitDB(t *testing.T) {
	InitDB()
	t.Log(PsgCon)
}
