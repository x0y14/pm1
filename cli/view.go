package cli

import "fmt"

type View struct {
	Action func(Model) Model
	Render func(Model) string
}

var WaitingForToFinishLoadingStorage = View{
	Action: func(model Model) Model {
		return model
	},
	Render: func(model Model) string {
		return fmt.Sprintf("loading storage...\n")
	},
}

var WaitingForToFinishEnteringMasterPassword = View{
	Action: func(model Model) Model {
		enteredMasterPassword := model.textInput.Value()
		if enteredMasterPassword != "pass" {
			return model
		}
		model.textInput.Blur()
		model.MainView = Success
		return model
	},
	Render: func(model Model) string {
		return fmt.Sprintf("please entering master password (length: 4 < n)\n%s\n", model.textInput.View())
	},
}

var Success = View{
	Action: func(model Model) Model {
		return model
	},
	Render: func(model Model) string {
		return fmt.Sprintf("successful loading")
	},
}
