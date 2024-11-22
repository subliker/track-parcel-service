package style

import tele "gopkg.in/telebot.v4"

var (
	MenuKeyboard     = makeMenuKeyboard()
	MenuBtnRegister  tele.Btn
	MenuBtnAddParcel tele.Btn
)

func makeMenuKeyboard() *tele.ReplyMarkup {
	mk := tele.ReplyMarkup{
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	MenuBtnRegister = mk.Text("Register")
	MenuBtnAddParcel = mk.Text("Add parcel")

	mk.Reply(
		mk.Row(MenuBtnRegister, MenuBtnAddParcel),
	)

	return &mk
}
