package user

import "github.com/start-codex/utils/password"
import "github.com/start-codex/utils/id"

type Builder struct {
	instance User
	err      error
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build() (User, error) {
	if b.err != nil {
		return b.instance, b.err
	}
	if b.instance.id == "" {
		uid, err := id.New()
		if err != nil {
			b.err = err
			return b.instance, b.err
		}
		b.instance.id = uid.String()
	}
	return b.instance, b.err
}

func (b *Builder) WithID(uuid string) *Builder {
	if uuid == "" {
		b.err = USER_ID_IS_EMPTY
		return b
	}
	uid, err := id.Parse(uuid)
	if err != nil {
		b.err = err
		return b
	}
	b.instance.id = uid.String()
	return b
}

func (b *Builder) WithEmail(email string) *Builder {
	b.instance.email = email
	return b
}

func (b *Builder) WithPhoneNumer(phoneNumber string) *Builder {
	b.instance.phoneNumber = phoneNumber
	return b
}

func (b *Builder) WithPassword(pwd string) *Builder {
	passwordHash, err := password.HashPassword(pwd, 14)
	if err != nil {
		b.err = err
		return b
	}
	b.instance.passwordHash = passwordHash.String()
	return b
}

func (b *Builder) WithRoleID(roleID string) *Builder {
	b.instance.roleID = roleID
	return b
}

func (b *Builder) WithProfileID(profileID string) *Builder {
	b.instance.profileID = profileID
	return b
}

func (b *Builder) WithFacebookID(facebookID string) *Builder {
	b.instance.facebookID = facebookID
	return b
}

func (b *Builder) WithGoogleID(googleID string) *Builder {
	b.instance.googleID = googleID
	return b
}

func (b *Builder) WithMicrosoftID(microsoftID string) *Builder {
	b.instance.microsoftID = microsoftID
	return b
}

func (b *Builder) WithVerified(verified bool) *Builder {
	b.instance.verified = verified
	return b
}

func (b *Builder) WithActive(active bool) *Builder {
	b.instance.active = active
	return b
}
