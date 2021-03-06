package actions

func (as *ActionSuite) Test_LinksResource_List() {
	user := as.Login()
	link := as.CreateLink(user)
	res := as.HTML("/links").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), link.Link)
	//as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_LinksResource_Show() {
	user := as.Login()
	link := as.CreateLink(user)
	res := as.HTML("/links/%s", link.ID).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), link.Link)
}

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
