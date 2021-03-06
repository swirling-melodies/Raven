package userWork

import (
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/common"
	"github.com/swirling-melodies/Raven/models/userModels"
	"xorm.io/xorm"
)

var engine *xorm.Engine

const (
	working = 1
	// Waiting to be used
	wait   = 2
	delect = 3
)

func initDB() {
	engine = common.InitDB()
}

func SetUser() (bool, error) {
	user := new(userModels.UserInfo)
	userAudit := new(userModels.UserInfoAudit)
	token := new(userModels.TokenInfo)
	tokenAudit := new(userModels.TokenInfoAudit)

	initDB()

	flag, err := engine.IsTableExist(user)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(user)
	}

	flag, err = engine.IsTableExist(userAudit)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(userAudit)
	}

	flag, err = engine.IsTableExist(token)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(token)
	}

	flag, err = engine.IsTableExist(tokenAudit)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(tokenAudit)
	}

	return addDefaultUser()
}

func addDefaultUser() (bool, error) {
	user := new(userModels.UserInfo)
	userAudit := new(userModels.UserInfoAudit)
	defaultUser := common.GetDefaultUserString()
	user.UserName = defaultUser.UserName
	user.Password = defaultUser.Password
	user.Status = working
	session := engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return false, err
	}

	_, err := session.Exec("truncate table UserInfo;")
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	_, err = session.Exec("truncate table UserInfoAudit;")
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	_, err = session.Insert(user)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	userAudit = setUserInfoAudit(user)
	userAudit.AuditType = working
	_, err = session.Insert(userAudit)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	session.Commit()

	return true, nil
}

func setUserInfoAudit(data *userModels.UserInfo) *userModels.UserInfoAudit {
	audit := userModels.UserInfoAudit{
		UserID:   data.ID,
		UserName: data.UserName,
		Password: data.Password,
		Status:   data.Status,
	}

	return &audit
}
