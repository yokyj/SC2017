package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoOrmService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	tx := gormDb.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	ulist := make([]UserInfo, 0, 0)
	checkErr(gormDb.Find(&ulist).Error)
	return ulist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := UserInfo{}
	u.UID = id
	checkErr(gormDb.First(&u).Error)
	return &u
}