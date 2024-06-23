package handler

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
	"server/internal/core"
	"server/internal/core/domain"
)

type (
	CotacaoHandler interface {
		ObterCotacaoAtual(w http.ResponseWriter, r *http.Request)
		ObterCotacoesRegistradas(w http.ResponseWriter, r *http.Request)
	}

	cotacaoHandler struct {
		cotacaoCore core.CotacaoCore
	}
)

func NewCotacaoHandler(cotacaoCore core.CotacaoCore) CotacaoHandler {
	return &cotacaoHandler{
		cotacaoCore: cotacaoCore,
	}
}

func (h *cotacaoHandler) ObterCotacaoAtual(w http.ResponseWriter, r *http.Request) {
	cotacao, err := h.cotacaoCore.ObterCotacaoAtual()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		if err != nil {
			panic(err)
		}
		return
	}

	writeOK(w, cotacao)
}

func (h *cotacaoHandler) ObterCotacoesRegistradas(w http.ResponseWriter, r *http.Request) {
	cotacoes, err := h.cotacaoCore.ObterCotacoesRegistradas()
	if err != nil {
		writeError(w, err)
		return
	}

	formatResponse := r.URL.Query().Get("format")
	switch formatResponse {
	case "html":
		writeTemplate(w, cotacoes)
	case "console":
		writeConsole(cotacoes)
	default:
		writeOK(w, cotacoes)
	}
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(err.Error())
	if err != nil {
		panic(err)
	}
}

func writeOK(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		panic(err)
	}
}

func writeConsole(cotacoes []*domain.Cotacao) {
	format := "{{range .}} " +
		"Cotacão {{.ID}} -  " +
		"Code: {{.Code}}, " +
		"Codein: {{.Codein}}, " +
		"Name: {{.Name}}, " +
		"High: {{.High}}, " +
		"Low: {{.Low}}, " +
		"VarBid: {{.VarBid}}, " +
		"PctChange: {{.PctChange}}, " +
		"Bid: {{.Bid}}, " +
		"Ask: {{.Ask}}, " +
		"Timestamp: {{.Timestamp}}, " +
		"CreateDate: {{.CreateDate}}.\n{{end}}"

	tmp, err := template.New("CotacaoTemplate").Parse(format)
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, &cotacoes)
	if err != nil {
		panic(err)
	}
}

func writeTemplate(wr io.Writer, data interface{}) {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}
