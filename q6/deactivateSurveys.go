package q6

import (
	"simplesurveygo/dao"
	"time"
)

func DeactivateSurvey() {
	res := dao.GetAllActiveSurveys()

	if res == nil {
		return
	}

	for _, survey := range res {
		// fmt.Println(survey.Status, survey.Expiry)

		diff := survey.Expiry.Sub(time.Now())
		if diff < 0 {
			dao.UpdateStatus(survey.SurveyName, false)
		}
	}
}
