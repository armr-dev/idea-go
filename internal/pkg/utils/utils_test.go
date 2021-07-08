package utils

import (
	"math"
	"testing"
)

func TestUtils(t *testing.T) {
	// t.Run("Split Function (even)", func(t *testing.T) {
	// 	var k1, k2, k3, k4, k5, k6, k7, k8 []byte
	// 	var err error
	// 	key := []byte("abcdefgh")

	// 	k1, k2, k3, k4, k5, k6, k7, k8, err = SplitKey(key)

	// 	if (err != nil) {
	// 		t.Error(err)
	// 	}

	// 	expected1 := []byte("a")
	// 	expected2 := []byte("b")
	// 	expected3 := []byte("c")
	// 	expected4 := []byte("d")
	// 	expected5 := []byte("e")
	// 	expected6 := []byte("f")
	// 	expected7 := []byte("g")
	// 	expected8 := []byte("h")

	// 	if !reflect.DeepEqual(expected1, k1) {
	// 		t.Errorf("expected '%x' but got '%x'", expected1, k1)
	// 	}

	// 	if !reflect.DeepEqual(expected2, k2) {
	// 		t.Errorf("expected '%x' but got '%x'", expected2, k2)
	// 	}

	// 	if !reflect.DeepEqual(expected3, k3) {
	// 		t.Errorf("expected '%x' but got '%x'", expected3, k3)
	// 	}

	// 	if !reflect.DeepEqual(expected4, k4) {
	// 		t.Errorf("expected '%x' but got '%x'", expected4, k4)
	// 	}

	// 	if !reflect.DeepEqual(expected5, k5) {
	// 		t.Errorf("expected '%x' but got '%x'", expected5, k5)
	// 	}

	// 	if !reflect.DeepEqual(expected6, k6) {
	// 		t.Errorf("expected '%x' but got '%x'", expected6, k6)
	// 	}

	// 	if !reflect.DeepEqual(expected7, k7) {
	// 		t.Errorf("expected '%x' but got '%x'", expected7, k7)
	// 	}

	// 	if !reflect.DeepEqual(expected8, k8) {
	// 		t.Errorf("expected '%x' but got '%x'", expected8, k8)
	// 	}
	// })

	t.Run("Modulo Add", func(t *testing.T) {
		var expectedLt, expectedEq, expectedGt byte;
		expectedLt = 2
		expectedEq = 0
		expectedGt = 1
		
		lt := AddModulo(1, 1)
		eq := AddModulo(byte(math.Pow(2, 16)), 0)
		gt := AddModulo(byte(math.Pow(2, 16)), 1)

		if (lt != expectedLt) {
			t.Errorf("expected '%x' but got '%x'", expectedLt, lt)
		}

		if (eq != expectedEq) {
			t.Errorf("expected '%x' but got '%x'", expectedEq, eq)
		}

		if (gt != expectedGt) {
			t.Errorf("expected '%x' but got '%x'", expectedGt, gt)
		}
	})

	t.Run("Modulo Multiply", func(t *testing.T) {
		var expectedLt, expectedEq, expectedGt byte;
		expectedLt = 10
		expectedEq = 1
		expectedGt = 2
		
		lt := MultiplyModulo(2, 5)
		eq := MultiplyModulo(byte(math.Pow(2, 16)+1), 1)
		gt := MultiplyModulo(byte(math.Pow(2, 16)+2), 1)

		if (lt != expectedLt) {
			t.Errorf("expected '%x' but got '%x'", expectedLt, lt)
		}

		if (eq != expectedEq) {
			t.Errorf("expected '%x' but got '%x'", expectedEq, eq)
		}

		if (gt != expectedGt) {
			t.Errorf("expected '%x' but got '%x'", expectedGt, gt)
		}
	})
}