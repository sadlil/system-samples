package copier

import (
	"reflect"
	"testing"
)

type testStruct struct {
	Field      string
	Otherfield int

	Pointer *testStructInternal
	Value   testStructInternal

	Slice []int
}

type testStructInternal struct {
	Field string
	Slice []string
}

func TestCopy(t *testing.T) {
	testCases := []struct {
		name     string
		src      *testStruct
		dest     any
		expected *testStruct
		wantErr  bool
	}{
		{
			name:    "copy_struct_to_struct",
			src:     &testStruct{Field: "foo"},
			dest:    &testStruct{},
			wantErr: false,
		},
		{
			name:    "copy_struct_to_struct_multivalue",
			src:     &testStruct{Field: "foo", Otherfield: 1},
			dest:    &testStruct{},
			wantErr: false,
		},
		{
			name:    "copy_struct_to_struct_multivalue_dest_set",
			src:     &testStruct{Field: "foo", Otherfield: 1},
			dest:    &testStruct{Field: "bar"},
			wantErr: false,
		},
		{
			name:    "copy_struct_to_struct_with_internal",
			src:     &testStruct{Field: "foo", Otherfield: 1, Pointer: &testStructInternal{Field: "foo!"}},
			dest:    &testStruct{Field: "bar"},
			wantErr: false,
		},
		{
			name:    "copy_struct_to_struct_with_value",
			src:     &testStruct{Field: "foo", Otherfield: 1, Value: testStructInternal{Field: "foo!"}},
			dest:    &testStruct{Field: "bar"},
			wantErr: false,
		},
		{
			name:    "copy_struct_to_struct_with_slice",
			src:     &testStruct{Field: "foo", Otherfield: 1, Slice: []int{1, 2, 3}},
			dest:    &testStruct{Field: "bar"},
			wantErr: false,
		},
		{
			name:     "copy_struct_to_incompatible_type",
			src:      &testStruct{Field: "foo"},
			dest:     1,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(tc.src, tc.dest)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Errorf("got error %v; want error %v", err, tc.wantErr)
				return
			}

			dest, ok := tc.dest.(*testStruct)
			if !ok {
				t.Fatalf("tc.dest is not *testStruct")
			}

			if !reflect.DeepEqual(dest, tc.src) {
				t.Errorf("got %v; want %v", tc.dest, tc.expected)
			}
		})
	}
}
