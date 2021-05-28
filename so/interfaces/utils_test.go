package interfaces

import "testing"

func TestValidateParamsType(t *testing.T) {
	// Base type
	err := ValidateParamsType(
		[]interface{}{[]byte("123"), 123, "123", '1', 1.1, uint(1), int64(1)},
		[]interface{}{[]byte("321"), 321, "321", '2', 2.2, uint(2), int64(2)},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = ValidateParamsType( // params' length not equal
		[]interface{}{123, "321"},
		[]interface{}{321, "123", '1'},
	)
	if err == nil {
		t.Fatalf("params length not equal, expect err is not nil")
	}
	t.Log(err) // params length not equal, expect 3, actual: 2

	err = ValidateParamsType( // type of params not equal
		[]interface{}{123, "321"},
		[]interface{}{"123", "123"},
	)
	if err == nil {
		t.Fatalf("type of params not equal, expect err is not nil")
	}
	t.Log(err) // params[0] type not equal, expect: string, actual: int

	// complex type
	err = ValidateParamsType(
		[]interface{}{testA{}, &testA{}},
		[]interface{}{testA{}, &testA{}},
	)
	if err != nil {
		t.Fatal(err)
	}
}

type testA struct {
	a string
	b int
	c chan int
	d []interface{}
	e float64
}
