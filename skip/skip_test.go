package skip

import "testing"

//var shouldBeSkipped = false

func TestThing(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Log("this test ran 凸⁠ಠ⁠益⁠ಠ)⁠凸")
}
