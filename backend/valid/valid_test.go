package valid

import (
	"testing"
)

func TestUsername(t *testing.T) {
	tests := []struct {
		name  string
		valid bool
	}{
		{"", false},
		{"giornetta      ", false},
		{"     ", false},
		{"    giornetta", false},
		{"21giornetta", false},
		{"giornetta_", false},
		{"gio??rnet0ta", false},
		{"giORNEtta", true},
		{"giornetta00", true},
		{"giornetta_00", true},
		{"michele-giornetta", true},
	}
	for _, test := range tests {
		valid := Username(test.name)
		if test.valid == valid {
			continue
		}
		t.Errorf("%q: expected valid=%t; got error %v", test.name, test.valid, valid)
	}
}

func TestPassword(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"", false},
		{"password      ", false},
		{"     ", false},
		{"    password", false},
		{"21password", true},
		{"pas??word__", true},
		{"gio??rnet0ta", true},
		{"giORNEtta", true},
		{"giornetta00", true},
		{"gi??!!e#tta_00", true},
		{"éé°*helloù", false},
	}
	for _, test := range tests {
		valid := Password(test.password)
		if test.valid == valid {
			continue
		}
		t.Errorf("%q: expected valid=%t; got %v", test.password, test.valid, valid)
	}
}
