package security

import "testing"

func TestIsValidSignature(t *testing.T) {
	type args struct {
		body      []byte
		signature string
		secret    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]byte("hello"), "sha1=9ade18f3e0ee81a5343f4a005f795dbaf9ceefd8", "hello"}, true},
		{"2", args{[]byte("hello"), "9ade18f3e0ee81a5343f4a005f795dbaf9ceefd8", "hello"}, false},
		{"3", args{[]byte("hello"), "sha1=9ade18f3e0ee81a5343f4a005f795dbaf9ceefd8", "false"}, false},
		{"4", args{[]byte("hello"), "sha1=9ade18f3e0ee81a5343f4a005f7", "false"}, false},
	}
	for _, tt := range tests {
		secret = tt.args.secret
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSignature(tt.args.body, tt.args.signature); got != tt.want {
				t.Errorf("IsValidSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
