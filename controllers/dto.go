// dto
package controllers

import (
	. "appinhouse/constants"
)

type ItemsResponseDto struct {
	Code  ErrCode    `json:"code"`
	Items []*ItemDto `json:"items"`
}

func NewSuccessItemsResponseDto() *ItemsResponseDto {
	i := &ItemsResponseDto{
		Code: ErrOk,
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
