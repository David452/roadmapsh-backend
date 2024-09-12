package handlers

import (
	"log"
	"html/template"
	"net/http"
	"strconv"

	"github.com/David452/unit-converter/internal/utils"
	"github.com/David452/unit-converter/internal/convertmap"
)

type SumData struct {
	UnitFrom 		string
	UnitTo 			string
	OriginalNum 	int
	ConvertedNum 	float64
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {

	valueStr := r.FormValue("value")
	unitFromStr := r.FormValue("unit_from")
	unitToStr := r.FormValue("unit_to")
	isLengthStr := r.FormValue("is_length")

	log.Print(valueStr, unitFromStr, unitToStr, isLengthStr)

	var conversionMap map[string]float64

	if isLengthStr == "on" {
		conversionMap = convertmap.LengthConversionMap
	} else {
		conversionMap = convertmap.WeightConversionMap
	}
	


	originalValue, err := strconv.Atoi(valueStr)
	
	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convertedValue := utils.ConvertUnits(float64(originalValue), unitFromStr, unitToStr, conversionMap)


	data := SumData{
		UnitFrom: unitFromStr,
		UnitTo: unitToStr,
		OriginalNum: originalValue,
		ConvertedNum: convertedValue,
	}

	tmpl, err := template.ParseFiles("./static/result.html")

	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	

	w.Header().Set("Content-Type", "text/html")

	tmpl.Execute(w, data)
}