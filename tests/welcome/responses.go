package welcome

import "registrationtogames/fmtogram/types"

func JustTrash(responses chan *types.TelegramResponse) {
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "Hello World ma booooooooy",
				},
			},
		},
	}
}

func JustTrash2(responses chan *types.TelegramResponse) {
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "Owww, nooo, you are the best human in the world!",
				},
			},
		},
	}
}

func Start(responses chan *types.TelegramResponse) {
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "/start",
				},
			},
		},
	}
}

func QueryForShowRules(responses chan *types.TelegramResponse) {
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "GoReg",
				},
			},
		},
	}
}

func QueryForWelcomeToMainMenu(responses chan *types.TelegramResponse) {
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "GoNext",
				},
			},
		},
	}
}
