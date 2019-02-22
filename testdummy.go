package igdb

//go:generate gomodifytags -file $GOFILE -struct TestDummy -add-tags json -w

type TestDummy struct {
	BoolValue       bool          `json:"bool_value"`
	CreatedAt       int           `json:"created_at"`
	EnumTest        TestDummyEnum `json:"enum_test"`
	FloatValue      float64       `json:"float_value"`
	Game            int           `json:"game"`
	IntegerArray    []int         `json:"integer_array"`
	IntegerValue    int           `json:"integer_value"`
	Name            string        `json:"name"`
	NewIntegerValue int           `json:"new_integer_value"`
	Private         bool          `json:"private"`
	Slug            string        `json:"slug"`
	StringArray     []string      `json:"string_array"`
	TestDummies     []int         `json:"test_dummies"`
	TestDummy       int           `json:"test_dummy"`
	UpdatedAt       int           `json:"updated_at"`
	URL             string        `json:"url"`
	User            int           `json:"user"`
}

//go:generate stringer -type=TestDummyEnum

type TestDummyEnum int

const (
	TestDummyEnum1 TestDummyEnum = iota + 1
	TestDummyEnum2
)
