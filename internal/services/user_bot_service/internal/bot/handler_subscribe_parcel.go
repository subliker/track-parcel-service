package bot

import (
	"context"
	"errors"

	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/pu"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/pupb"
	tele "gopkg.in/telebot.v4"
)

func (b *bot) handleSubscribeParcel(subscribe bool) tele.HandlerFunc {
	return func(ctx tele.Context) error {
		defer ctx.Respond()
		trackNum := ctx.Data()
		// if subscribe action
		if subscribe {
			err := b.parcelsUserClient.AddSubscription(context.Background(), &pupb.AddSubscriptionRequest{
				TrackNumber:    trackNum,
				UserTelegramId: ctx.Sender().ID,
			})
			if errors.Is(err, pu.ErrAlreadyExists) {
				ctx.Send(b.bundle.CheckParcel().Errors().AlreadySubscribed())
				return nil
			}
			if errors.Is(err, pu.ErrParcelNotFound) {
				ctx.Send(b.bundle.CheckParcel().Errors().NotFound())
				return nil
			}
			if errors.Is(err, pu.ErrInternal) {
				ctx.Send(b.bundle.Common().Errors().Internal())
				return err
			}

			// make describe btn in old message
			mk := b.client.NewMarkup()
			dBtn := checkParcelBtnDescribe
			dBtn.Data = trackNum
			mk.Inline(mk.Row(dBtn))
			ctx.Edit(ctx.Text(), mk)

			// send ready msg
			ctx.Send(b.bundle.CheckParcel().SubscribeEvent().Subscribed(ctx.Callback().Data))
		} else {
			err := b.parcelsUserClient.DeleteSubscription(context.Background(), &pupb.DeleteSubscriptionRequest{
				TrackNumber:    ctx.Data(),
				UserTelegramId: ctx.Sender().ID,
			})
			if errors.Is(err, pu.ErrSubscriptionNotFound) {
				ctx.Send(b.bundle.CheckParcel().Errors().AlreadyDescribed())
				return nil
			}
			if errors.Is(err, pu.ErrInternal) {
				ctx.Send(b.bundle.Common().Errors().Internal())
				return err
			}

			// make describe btn in old message
			mk := b.client.NewMarkup()
			dBtn := checkParcelBtnSubscribe
			dBtn.Data = trackNum
			mk.Inline(mk.Row(dBtn))
			ctx.Edit(ctx.Text(), mk)

			// send ready msg
			ctx.Send(b.bundle.CheckParcel().SubscribeEvent().Described(ctx.Callback().Data))
		}
		return nil
	}
}
