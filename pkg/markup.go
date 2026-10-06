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
