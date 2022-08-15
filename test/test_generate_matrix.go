package main

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"io/ioutil"
	"strings"
)

func main() {
	utgSli := []string{}
	mpSli := []string{}
	coSli := []string{}
	btnSli := []string{}
	sbSli := []string{}

	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			cardStr := models.MatrixToHolecards(i, j)
			utgElem := fmt.Sprintf(`var raise%s = fetchRaiseCallFoldData(utgRange, "%s").raise;
var call%s = fetchRaiseCallFoldData(utgRange, "%s").call;
var fold%s = fetchRaiseCallFoldData(utgRange, "%s").fold;
document.documentElement.style.setProperty('--raise%s',(raise%s/(raise%s+call%s+fold%s)*100)+"%%");
document.documentElement.style.setProperty('--call%s',(call%s/(raise%s+call%s+fold%s)*100)+"%%");`,
				cardStr, cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr)
			mpElem := fmt.Sprintf(`var raise%s = fetchRaiseCallFoldData(mpRange, "%s").raise;
var call%s = fetchRaiseCallFoldData(mpRange, "%s").call;
var fold%s = fetchRaiseCallFoldData(mpRange, "%s").fold;
document.documentElement.style.setProperty('--raise%s',(raise%s/(raise%s+call%s+fold%s)*100)+"%%");
document.documentElement.style.setProperty('--call%s',(call%s/(raise%s+call%s+fold%s)*100)+"%%");`,
				cardStr, cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr)
			coElem := fmt.Sprintf(`var raise%s = fetchRaiseCallFoldData(coRange, "%s").raise;
var call%s = fetchRaiseCallFoldData(coRange, "%s").call;
var fold%s = fetchRaiseCallFoldData(coRange, "%s").fold;
document.documentElement.style.setProperty('--raise%s',(raise%s/(raise%s+call%s+fold%s)*100)+"%%");
document.documentElement.style.setProperty('--call%s',(call%s/(raise%s+call%s+fold%s)*100)+"%%");`,
				cardStr, cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr)
			btnElem := fmt.Sprintf(`var raise%s = fetchRaiseCallFoldData(btnRange, "%s").raise;
var call%s = fetchRaiseCallFoldData(btnRange, "%s").call;
var fold%s = fetchRaiseCallFoldData(btnRange, "%s").fold;
document.documentElement.style.setProperty('--raise%s',(raise%s/(raise%s+call%s+fold%s)*100)+"%%");
document.documentElement.style.setProperty('--call%s',(call%s/(raise%s+call%s+fold%s)*100)+"%%");`,
				cardStr, cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr)
			sbElem := fmt.Sprintf(`var raise%s = fetchRaiseCallFoldData(sbRange, "%s").raise;
var call%s = fetchRaiseCallFoldData(sbRange, "%s").call;
var fold%s = fetchRaiseCallFoldData(sbRange, "%s").fold;
document.documentElement.style.setProperty('--raise%s',(raise%s/(raise%s+call%s+fold%s)*100)+"%%");
document.documentElement.style.setProperty('--call%s',(call%s/(raise%s+call%s+fold%s)*100)+"%%");`,
				cardStr, cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr,cardStr)
			utgSli = append(utgSli, utgElem)
			mpSli = append(mpSli, mpElem)
			coSli = append(coSli, coElem)
			btnSli = append(btnSli, btnElem)
			sbSli = append(sbSli, sbElem)
		}
	}
	ioutil.WriteFile("/Users/luolingxiao/projects/shallwepoker/ggpoker_hands_converter/output/utg", []byte(strings.Join(utgSli, "\n\n")), 0644)
	ioutil.WriteFile("/Users/luolingxiao/projects/shallwepoker/ggpoker_hands_converter/output/mp", []byte(strings.Join(mpSli, "\n\n")), 0644)
	ioutil.WriteFile("/Users/luolingxiao/projects/shallwepoker/ggpoker_hands_converter/output/co", []byte(strings.Join(coSli, "\n\n")), 0644)
	ioutil.WriteFile("/Users/luolingxiao/projects/shallwepoker/ggpoker_hands_converter/output/btn", []byte(strings.Join(btnSli, "\n\n")), 0644)
	ioutil.WriteFile("/Users/luolingxiao/projects/shallwepoker/ggpoker_hands_converter/output/sb", []byte(strings.Join(sbSli, "\n\n")), 0644)

}
