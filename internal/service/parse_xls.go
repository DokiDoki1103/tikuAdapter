package service

import (
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type XLSXOptions struct {
	SheetName     string   `json:"sheet"`
	QuestionIndex string   `json:"q"`
	AnswerIndex   string   `json:"a"`
	OptionIndex   []string `json:"o"`
}

func ParseXls(file *excelize.File, opt XLSXOptions) []model.Question {
	rows, err := file.GetRows(opt.SheetName)
	if err != nil {
		return nil
	}
	var tikus []model.Question
	for i := 0; i < len(rows); i++ {
		index := strconv.Itoa(i)
		question, _ := file.GetCellValue(opt.SheetName, opt.QuestionIndex+index)
		answer, _ := file.GetCellValue(opt.SheetName, opt.AnswerIndex+index)

		var optoins = make([]string, 0)
		for _, v := range opt.OptionIndex {
			option, _ := file.GetCellValue(opt.SheetName, v+index)
			optoins = append(optoins, option)
		}
		tiku := model.Question{
			Question: question,
			Answer:   []string{answer},
			Options:  optoins,
		}
		tikus = append(tikus, tiku)
	}
	return tikus
}
