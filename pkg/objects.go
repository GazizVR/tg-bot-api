package pkg

type Update struct {
	Id       int64          `json:"update_id"`
	Message  *Message       `json:"message"`
	Callback *CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
	Id      string              `json:"id"`
	Data    string              `json:"data"`
	Message InaccessibleMessage `json:"message"`
}

type InaccessibleMessage struct {
	Id   int64 `json:"message_id"`
	Chat Chat  `json:"chat"`
}

type Message struct {
	Id          int64           `json:"message_id"`
	Text        string          `json:"text"`
	LinkPreview *LinkPreviewOps `json:"link_preview_options"`
	ReplyTo     *Message        `json:"reply_to_message"`
	Chat        Chat            `json:"chat"`
	Video       *Media          `json:"video"`
	Audio       *Media          `json:"audio"`
	Voice       *Media          `json:"voice"`
	VideoNote   *Media          `json:"video_note"`
	Document    *Document       `json:"document"`
}

type LinkPreviewOps struct {
	URL string `json:"url"`
}

type Chat struct {
	Id int64 `json:"id"`
}

type Media struct {
	Id string `json:"file_id"`
}

type Document struct {
	Media
	MimeType string `json:"mime_type,omitempty"`
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

type InlineButton struct {
	Text string `json:"text"`
	Data string `json:"callback_data"`
}

type File struct {
	UniqueId string `json:"file_unique_id"`
	Path     string `json:"file_path,omitempty"`
}
