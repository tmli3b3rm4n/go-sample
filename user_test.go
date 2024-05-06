package user

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	root := &User{
		Name:      "Root",
		Referrals: map[string]*User{},
	}

	type args struct {
		n string
		p *User
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "Create user with parent",
			args: args{n: "Child", p: root},
			want: &User{Name: "Child", Parent: root, Referrals: map[string]*User{}},
		},
		{
			name: "Create user without parent",
			args: args{n: "Orphan", p: nil},
			want: &User{Name: "Orphan", Parent: nil, Referrals: map[string]*User{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.n, tt.args.p); !reflect.DeepEqual(got, tt.want) || (tt.args.p != nil && tt.args.p.Referrals[tt.args.n] != got) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointReferrals(t *testing.T) {
	a := NewUser("A", nil)
	b := NewUser("B", a)
	NewUser("C", a)
	NewUser("D", b)
	NewUser("E", b)

	type args struct {
		u *User
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Referral count for A",
			args: args{u: a},
			want: 4, // A has B, C directly and D, E via B
		},
		{
			name: "Referral count for B",
			args: args{u: b},
			want: 2, // B has D, E
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointReferrals(tt.args.u); got != tt.want {
				t.Errorf("PointReferrals() = %v, want %v", got, tt.want)
			}
		})
	}
}
