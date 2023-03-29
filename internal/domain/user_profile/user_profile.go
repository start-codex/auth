package userprofile

type UserProfile struct {
	id               string
	userId           string
	firstName        string
	lastName         string
	profilePicture   string
	registrationDate string
	lastLoginDate    string
}

func (u *UserProfile) ID() string {
	return u.id
}

func (u *UserProfile) UserID() string {
	return u.userId
}

func (u *UserProfile) FirstName() string {
	return u.firstName
}

func (u *UserProfile) LastName() string {
	return u.lastName
}

func (u *UserProfile) ProfilePicture() string {
	return u.profilePicture
}

func (u *UserProfile) RegistrationDate() string {
	return u.registrationDate
}

func (u *UserProfile) LastLoginDate() string {
	return u.lastLoginDate
}
