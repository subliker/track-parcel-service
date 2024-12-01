package bot

import tele "gopkg.in/telebot.v4"

var startBtnRegister tele.Btn

func (b *bot) handleStart() tele.HandlerFunc {
	mk := b.client.NewMarkup()

	startBtnRegister = mk.Data(b.bundle.StartMessage().Markup().Register(), "start-register")

	mk.Inline(mk.Row(startBtnRegister))
	return func(ctx tele.Context) error {
		// set handler name
		ctx.Set("handler", "start")

		// redirect if authorized
		authorized, _ := ctx.Get("authorized").(bool)
		if authorized {
			return b.handleMenu()(ctx)
		}

		ctx.Send(b.bundle.StartMessage().Head(ctx.Sender().FirstName)+"\n"+b.bundle.StartMessage().Main(), mk)
		return nil
	}
}
