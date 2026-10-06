package pkg

type ReplyMarkup interface {
	isReplyMarkup()
}

type InlineKeyboardButton struct {
	Text     string          `json:"text"`
	Data     *string         `json:"callback_data,omitempty"`
	CopyText *CopyTextButton `json:"copy_text,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (InlineKeyboardMarkup) isReplyMarkup() {}

func NewInlineMarkup(
	rows ...[]InlineKeyboardButton,
) *InlineKeyboardMarkup {
	return &InlineKeyboardMarkup{
		Keyboard: rows,
	}
}

type ReplyKeyboardButton struct {
	Text           string `json:"text"`
	RequestContact *bool  `json:"request_contact,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard       [][]ReplyKeyboardButton `json:"keyboard"`
	ResizeKeyboard *bool                   `json:"resize_keyboard,omitempty"`
}

func (ReplyKeyboardMarkup) isReplyMarkup() {}

type ReplyMarkupOption func(*ReplyKeyboardMarkup)

func WithRequestContact(resizeKeyboard bool) ReplyMarkupOption {
	return func(rm *ReplyKeyboardMarkup) {
		rm.ResizeKeyboard = &resizeKeyboard
	}
}

func NewReplyMarkup(
	rows [][]ReplyKeyboardButton,
	opts ...ReplyMarkupOption,
) *ReplyKeyboardMarkup {
	replyMarkup := &ReplyKeyboardMarkup{
		Keyboard: rows,
	}
	for _, opt := range opts {
		opt(replyMarkup)
	}
	return replyMarkup
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
}

func (ReplyKeyboardRemove) isReplyMarkup() {}
