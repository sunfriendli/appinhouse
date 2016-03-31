// dto
package controllers

import (
	. "appinhouse/inhouse_service/constants"
)

type ItemsResponseDto struct {
	Code  ErrCode    `json:"code"`
	Msg   string     `json:"msg"`
	Items []*ItemDto `json:"items"`
}

func NewSuccessItemsResponseDto() *ItemsResponseDto {
	i := &ItemsResponseDto{
		Code:  ErrOk,
		Items: []*ItemDto{},
	}
	return i
}

type ItemsDto []*ItemDto

type ItemDto struct {
	Platform    string `json:"platform"`
	Env         string `json:"environment"`
	Version     string `json:"version"`
	Time        string `json:"time"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Down        string `json:"down"`
	Channel     string `json:"channel"`
}
type ItemsResponsePageDto struct {
	Code      ErrCode   `json:"code"`
	Msg       string    `json:"msg"`
	Items     *ItemsDto `json:"items"`
	Page      int       `json:"page"`
	TotalPage int       `json:"total_page"`
}

func NewSuccessItemsResponsePageDto() *ItemsResponsePageDto {
	i := &ItemsResponsePageDto{
		Code:  ErrOk,
		Items: new(ItemsDto),
	}
	return i
}

type ResponseDto struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"msg"`
}

func NewSuccessResponseDto() *ResponseDto {
	i := &ResponseDto{
		Code: ErrOk,
	}
	return i
}
