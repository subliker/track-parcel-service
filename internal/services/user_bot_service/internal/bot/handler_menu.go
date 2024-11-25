package bot

import tele "gopkg.in/telebot.v4"

var menuBtnCheckParcel tele.Btn

func (b *bot) handleMenu() tele.HandlerFunc {
	mk := b.client.NewMarkup()
	mk.ResizeKeyboard = true
	mk.OneTimeKeyboard = true

	menuBtnCheckParcel = mk.Text(b.bundle.Menu().Markup().CheckParcel())

	mk.Reply(mk.Row(menuBtnCheckParcel))
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "menu")

		ctx.Send(b.bundle.Menu().Main(), mk)
		return nil
	}
}
