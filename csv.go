package helper

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

func ExportCsv(w http.ResponseWriter, records [][]string, fileName string) {
	if fileName == "" {
		fileName = "file"
	}
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	var buf bytes.Buffer
	buf.WriteString("\xef\xbb\xbf")
	cw := csv.NewWriter(&buf)
	if err := cw.Error(); err != nil {
		return
	}
	cw.WriteAll(records)
	io.Copy(w, &buf)
}
