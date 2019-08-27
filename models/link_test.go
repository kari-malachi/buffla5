package models_test

import (
	"github.com/gobuffalo/uuid"
	"github.com/kari-malachi/buffla5/models"
)

func (ms *ModelSuite) Test_Link_BeforeValidations() {
	uid, _ := uuid.NewV4()
	link := &models.Link{
		UserID: uid,
		Link:   "http://gobuffalo.io",
	}
	err := link.BeforeValidations(ms.DB)
	ms.NoError(err)
	ms.NotZero(link.Code)
	ms.Len(link.Code, 7)
}
