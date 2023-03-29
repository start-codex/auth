package user

type User struct {
	id           string
	email        string
	phoneNumber  string
	passwordHash string
	roleID       string
	profileID    string
	facebookID   string
	googleID     string
	microsoftID  string
	verified     bool
	active       bool
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) PhoneNumber() string {
	return u.phoneNumber
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) RoleID() string {
	return u.roleID
}

func (u *User) ProfileID() string {
	return u.profileID
}

func (u *User) FacebookID() string {
	return u.facebookID
}

func (u *User) GoogleID() string {
	return u.googleID
}

func (u *User) MicrosoftID() string {
	return u.microsoftID
}

func (u *User) IsVerified() bool {
	return u.verified
}

func (u *User) IsActive() bool {
	return u.active
}
