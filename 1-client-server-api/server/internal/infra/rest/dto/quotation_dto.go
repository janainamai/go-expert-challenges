package dto

type (
	QuotationDTO struct {
		ContentDTO ContentDTO `json:"usdbrl"`
	}

	ContentDTO struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varbid"`
		PctChange  string `json:"pctchange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	}
)
