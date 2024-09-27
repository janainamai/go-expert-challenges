package dto

type (
	Error struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
	}
)

func InitError() *Error {
	return &Error{}
}

func NewError(title, detail string) *Error {
	return &Error{
		Title:  title,
		Detail: detail,
	}
}

func (e *Error) GetTitle() string {
	return e.Title
}

func (e *Error) GetDetail() string {
	return e.Detail
}

func (e *Error) WithTitle(title string) *Error {
	e.Title = title
	return e
}

func (e *Error) WithDetail(detail string) *Error {
	e.Detail = detail
	return e
}
