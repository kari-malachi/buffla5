package actions

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
	"github.com/kari-malachi/buffla5/models"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func (as *ActionSuite) CreateUser() *models.User {
	user := &models.User{
		Name:  "Kari",
		Email: "kari@ma.com",
	}
	as.NoError(as.DB.Create(user))
	return user
}

func (as *ActionSuite) Login() *models.User {
	user := as.CreateUser()
	as.Session.Set("current_user_id", user.ID)
	return user
}

func (as *ActionSuite) CreateLink(user *models.User) *models.Link {
	link := &models.Link{
		Link:   "www.vimeo.com",
		Code:   "12345",
		UserID: user.ID,
	}
	verrs, err := as.DB.ValidateAndCreate(link)
	as.NoError(err)
	as.False(verrs.HasAny())
	return link
}
