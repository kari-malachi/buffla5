package actions

import "github.com/kari-malachi/buffla5/models"

func (as *ActionSuite) Test_LinksResource_List() {
	user := &models.User{
		Name:  "Kari",
		Email: "kari@ma.com",
	}
	as.NoError(as.DB.Create(user))
	as.Session.Set("current_user_id", user.ID)
	link := &models.Link{
		Link:   "www.vimeo.com",
		Code:   "12345",
		UserID: user.ID,
	}
	as.NoError(as.DB.Create(link))
	res := as.HTML("/links").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), link.Link)
	//as.Fail("Not Implemented!")
}

// func (as *ActionSuite) Test_LinksResource_Show() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_LinksResource_Create() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_LinksResource_Update() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_LinksResource_Destroy() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_LinksResource_New() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_LinksResource_Edit() {
// 	as.Fail("Not Implemented!")
// }
