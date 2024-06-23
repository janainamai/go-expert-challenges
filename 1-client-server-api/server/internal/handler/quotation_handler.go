package handler

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/internal/core"
	"server/internal/core/domain"
)

type (
	QuotationHandler interface {
		GetCurrentQuotation(w http.ResponseWriter, r *http.Request)
		GetRegistereQuotations(w http.ResponseWriter, r *http.Request)
	}

	quotationHandler struct {
		quotationCore core.QuotationCore
	}
)

func NewQuotationHandler(quotationCore core.QuotationCore) QuotationHandler {
	return &quotationHandler{
		quotationCore: quotationCore,
	}
}

func (h *quotationHandler) GetCurrentQuotation(w http.ResponseWriter, r *http.Request) {
	quotation, err := h.quotationCore.GetCurrentQuotation()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err.Error())
		if err != nil {
			panic(err)
		}
		return
	}

	writeOK(w, quotation)
}

func (h *quotationHandler) GetRegistereQuotations(w http.ResponseWriter, r *http.Request) {
	quotations, err := h.quotationCore.GetRegisteredQuotations()
	if err != nil {
		writeError(w, err)
		return
	}

	formatResponse := r.URL.Query().Get("format")
	switch formatResponse {
	case "html":
		writeTemplate(w, quotations)
	case "console":
		writeConsole(quotations)
	default:
		writeOK(w, quotations)
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

func writeConsole(quotations []*domain.Quotation) {
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

	tmp, err := template.New("QuotationTemplate").Parse(format)
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, &quotations)
	if err != nil {
		panic(err)
	}
}

func writeTemplate(wr io.Writer, data interface{}) {
	cwd, _ := os.Getwd()

	relativePaths := []string{
		"../internal/handler/templates/quotation.html",
		"internal/handler/templates/quotation.html",
	}

	var templatePath string
	for _, relPath := range relativePaths {
		absPath := filepath.Join(cwd, relPath)
		if _, err := os.Stat(absPath); err == nil {
			templatePath = absPath
			break
		}
	}

	tmp, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Erro ao carregar o template: %s", err)
	}

	err = tmp.Execute(wr, data)
	if err != nil {
		log.Fatalf("Erro ao executar o template: %s", err)
	}
}
