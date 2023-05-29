package models

import (
	"context"

	entmessage "tabelf/backend/gen/entschema/message"
	"tabelf/backend/service/app"
)

func SendMessage(
	ctx context.Context,
	mtp app.MessageType,
	promoterUID string,
	userUID string,
	description string,
) error {
	return app.EntClient.Message.Create().
		SetCategory(string(mtp)).
		SetPromoterUID(promoterUID).
		SetUserUID(userUID).
		SetDescription(description).
		Exec(ctx)
}

func ReadMessage(ctx context.Context, messageUID string, userUID string) error {
	return app.EntClient.Message.Update().
		SetHasRead(true).
		Where(
			entmessage.UID(messageUID),
			entmessage.UserUID(userUID),
		).Exec(ctx)
}
