package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pichik/webwatcher/src/datacenter"
	"github.com/pichik/webwatcher/src/misc"
)

func ImportTemplate() {
	filePath := misc.TemplateDir + "bait.js"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	baitContent = content
}

func throwBait(w http.ResponseWriter) {
	w.Header().Set("Content-Type", fmt.Sprintf("text/javascript; charset=utf-8"))
	w.Header().Set("Content-Length", strconv.Itoa(len(baitContent)))
	w.WriteHeader(http.StatusOK)

	_, err := w.Write(baitContent)
	if err != nil {
		misc.ErrorLog.Printf("Error throwing bait: %s", err)
	}
}

func extractJson(jsonData []byte) *datacenter.Data {
	var data datacenter.Data

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		misc.ErrorLog.Printf("%s", err)
	}
	return &data
}
