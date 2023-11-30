package admins

import (
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/adminsrepo"

	"golang.org/x/crypto/bcrypt"
)

func SaveAdmin(admin models.Admin) *core.HttpError {
	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)

	if err != nil {
		core.LogErr(err)
		return core.InternalError
	}

	err = adminsrepo.Save(admin.Name, string(password))

	if err != nil {
		return core.InternalError
	}

	return nil
}

func Login(admin models.Admin) (*core.JWT, *core.HttpError) {
	pwd, err := adminsrepo.Fetch(admin.Name)

	if err != nil {
		return nil, core.InternalError
	}

	if len(pwd) == 0 {
		return nil, core.NotFoundError
	}

	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(admin.Password))

	if err != nil {
		return nil, core.ForbiddenError
	}

	token := core.GenerateJWT(admin.Name)

	if token == nil {
		return nil, core.InternalError
	}

	return token, nil
}
