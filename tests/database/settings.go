package database

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/routine"
)

func withOutS(userId int) *bottypes.User {
	var (
		err  error
		user *bottypes.User
	)
	user = new(bottypes.User)
	user.Id = userId
	err = routine.DbRetrieveUser(user)
	if err != nil {
		panic(err)
	}
	user.ExMessageId, err = routine.SelectExMessageId(user.Id)
	if err != nil {
		panic(err)
	}
	if user.Id != 899 {
		panic("user.Id != 899")
	}
	if user.ExMessageId != 1111 {
		panic("user.ExMessageId != 1111")
	}
	if user.LaunchPoint != 0 {
		panic("user.LaunchPoint != 0")
	}
	if user.Reg.GameId != 0 {
		panic("user.Reg.GameId != 0")
	}
	if user.Reg.Seats != 0 {
		panic("user.Reg.Seats != 0")
	}
	if user.Reg.Payment != "" {
		panic("user.Reg.Payment != ``")
	}
	if user.Media.Interval != "" {
		panic("user.Media.Interval != ``")
	}
	if user.Media.Direction != "" {
		panic("user.Media.Direction != ``")
	}
	if user.Media.Limit != 7 {
		panic("user.Media.Limit != 7")
	}
	if user.Media.Id != "" {
		panic("user.Media.Id != ``")
	}
	if user.Media.Counter != 0 {
		panic("user.Media.Counter != 0")
	}
	return user
}

func AfterChooseOptionOnlyLang(userId int) {
	user := withOutS(userId)
	if user.Level != 2 {
		panic("user.Level != 2")
	}
	if user.Language != "ru" {
		panic("user.Language != `ru`")
	}
	if user.Act != "settings" {
		panic("user.Act != `settings`")
	}
	if user.UserRec.Changeable != "" {
		panic("user.UserRec.Changeable != ``")
	}
	if user.UserRec.ActGame != "" {
		panic("user.UserRec.ActGame != ``")
	}
	if user.UserRec.WillChangeable != "language" {
		panic("user.UserRec.WillChangeable != `language`")
	}
	if user.UserRec.NewPay != "" {
		panic("user.UserRec.NewPay != ``")
	}
}

func AfterChooseLanguage(userId int) {
	user := withOutS(userId)
	if user.Level != 3 {
		panic("user.Level != ")
	}
	if user.Language != "en" {
		panic("user.Language != `en`")
	}
	if user.Act != "divarication" {
		panic("user.Act != `divarication`")
	}
	if user.UserRec.Changeable != "" {
		panic("user.UserRec.Changeable != ``")
	}
	if user.UserRec.ActGame != "" {
		panic("user.UserRec.ActGame != ``")
	}
	if user.UserRec.WillChangeable != "language" {
		panic("user.UserRec.WillChangeable != `language`")
	}
	if user.UserRec.NewPay != "" {
		panic("user.UserRec.NewPay != ``")
	}
}
