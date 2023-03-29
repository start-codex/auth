package user

import "testing"

func Test_UserBuilder(t *testing.T) {

	user, err := NewBuilder().
		WithID("9e32cb94-01c2-45e8-845e-c91d8adb853e").
		WithEmail("email").
		WithPhoneNumer("phone").
		WithPassword("password").
		WithRoleID("role").
		WithProfileID("profile").
		WithFacebookID("facebook").
		WithGoogleID("google").
		Build()

	if err != nil {
		t.Error(err)
	}

	t.Logf("id: %v", user.ID())

	t.Log("UserBuilder_Test: OK")

}
