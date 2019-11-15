package slack

import "context"

type Service interface {
	SendMessage(ctx context.Context, msg *Message) error
	HandleMessage(ctx context.Context, input handleMessageInput) error
	HandleReaction(ctx context.Context, input handleReactionInput) error
}

type handleMessageInput struct {
	Message  string
	SenderID string
}

type handleReactionInput struct {
	SenderID    string
	RecipientID string
	Type        string
}
