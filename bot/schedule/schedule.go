package schedule

import (
	"RegistrationToGames/bot/bottypes"
	"RegistrationToGames/bot/dictionary"
	"RegistrationToGames/bot/forall"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"fmt"
)

func ShowTheSchedule(user *bottypes.User, fm *formatter.Formatter) {
	var (
		dict     map[string]string
		schedule []*forall.Game
		mes      string
	)
	dict = dictionary.Dictionary[user.Language]
	schedule = selectSchedule(user.Language)
	for i := 0; i < len(schedule); i++ {
		mes += (fmt.Sprintf(dict["Schedule"], i+1, schedule[i].Sport, schedule[i].Date, schedule[i].Time, schedule[i].Seats, schedule[i].Price, schedule[i].Currency))
	}
	fm.SetIkbdDim([]int{1})
	fm.WriteInlineButtonCmd(dict["MainMenu"], "MainMenu")
	fm.WriteParseMode(types.HTML)
	fm.WriteString(mes)
	fm.WriteChatId(user.Id)
}