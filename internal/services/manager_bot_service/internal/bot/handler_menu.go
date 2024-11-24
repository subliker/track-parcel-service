package bot

import tele "gopkg.in/telebot.v4"

var menuBtnAddParcel tele.Btn

func (b *bot) handleMenu() tele.HandlerFunc {
	mk := b.client.NewMarkup()
	mk.ResizeKeyboard = true
	mk.OneTimeKeyboard = true

	menuBtnAddParcel = mk.Text(b.bundle.Menu().Markup().AddParcel())

	mk.Reply(mk.Row(menuBtnAddParcel))
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "menu")

		ctx.Send(b.bundle.Menu().Main(), mk)
		return nil
	}
}
