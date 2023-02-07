package things

import "testing"

func TestACosa(t *testing.T) {

	if CosaSilla.String() != "Silla" {
		t.Error("fail to name apropiadly")
	}
}
