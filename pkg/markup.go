package pkg

type InlineButton struct {
	Text     string          `json:"text"`
	Data     *string         `json:"callback_data,omitempty"`
	CopyText *CopyTextButton `json:"copy_text,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type InlineMarkup struct {
	Keyboard [][]InlineButton `json:"inline_keyboard"`
}

func NewInlineMarkup(
	rows ...[]InlineButton,
) *InlineMarkup {
	return &InlineMarkup{
		Keyboard: rows,
	}
}

type ReplyButton struct {
	Text           string `json:"text"`
	RequestContact *bool  `json:"request_contact,omitempty"`
}

type ReplyMarkup struct {
	Keyboard       [][]ReplyButton `json:"keyboard"`
	ResizeKeyboard *bool           `json:"resize_keyboard,omitempty"`
}

type ReplyMarkupOption func(*ReplyMarkup)

func WithRequestContact(resizeKeyboard bool) ReplyMarkupOption {
	return func(rm *ReplyMarkup) {
		rm.ResizeKeyboard = &resizeKeyboard
	}
}

func NewReplyMarkup(
	rows [][]ReplyButton,
	opts ...ReplyMarkupOption,
) *ReplyMarkup {
	replyMarkup := &ReplyMarkup{
		Keyboard: rows,
	}
	for _, opt := range opts {
		opt(replyMarkup)
	}
	return replyMarkup
}
