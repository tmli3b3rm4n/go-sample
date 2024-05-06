package user

// User represents a user with a name, a reference to their parent user (if any), and a map of referrals.
type User struct {
	Name      string
	Referrals map[string]*User
	Parent    *User
}

// PointReferrals calculates the total number of referrals for a given user, including their referrals' referrals, recursively.
func PointReferrals(u *User) int {
	if len(u.Referrals) == 0 {
		return 0
	}
	count := 0
	for _, r := range u.Referrals {
		count += 1 + PointReferrals(r)
	}
	return count
}

// NewUser creates a new user with the given name and adds it to the parent user's referrals, if provided.
func NewUser(n string, p *User) *User {
	user := &User{
		Name:      n,
		Referrals: map[string]*User{},
	}
	if p != nil {
		user.Parent = p
		p.Referrals[n] = user
	}

	return user
}
