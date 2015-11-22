package pricecalc
import (
	"sort"
	"strconv"
	"strings"
	"time"

	"../../datastructures"
)

/**
 * Calculates the total advertisement cost
 * @param {[type]} purchaseInfo   datastructures.Purchase    [description]
 * @param {[type]} companyInfo    datastructures.CompanyInfo [description]
 * @param {[type]} purchaseReport datastructures.Report      [description]
 */
func AdCost(purchaseInfo datastructures.Purchase, companyInfo datastructures.CompanyInfo, purchaseReport datastructures.Report) float64 {
	cost := calcDaysCost(purchaseInfo, companyInfo, purchaseReport) * purchaseInfo.Columns * purchaseInfo.Rows
	cost += float64(5000 * len(strings.Split(purchaseInfo.AdColors, " ")))
	if purchaseInfo.DesignForYou == "yes" {
		cost += float64(100 * purchaseInfo.Columns * purchaseInfo.Rows)
	}

	cost += float64(cost * companyInfo.Tax)
	return cost
}

/**
 * Converts string representation of month to a time.Moth object
 * @param  {[type]} m string        [description]
 * @return {[type]}   [description]
 */
func monthToTime(m string) time.Month {
	if m == "January" {
		return time.January
	} else if m == "February" {
		return time.February
	} else if m == "March" {
		return time.March
	} else if m == "April" {
		return time.April
	} else if m == "May" {
		return time.May
	} else if m == "June" {
		return time.June
	} else if m == "July" {
		return time.July
	} else if m == "August" {
		return time.August
	} else if m == "September" {
		return time.September
	} else if m == "October" {
		return time.October
	} else if m == "November" {
		return time.November
	} else if m == "December" {
		return time.December
	} else {
		return time.January
	}
}

/**
 * Populates int slice with the days that advertisement is supposed to be shown
 * @param  {[type]} days string        [description]
 * @return {[type]}      [description]
 */
func days(days string) []int {
	var utilArray []int
	daysArray := strings.Split(days, " ")

	for _, day := range daysArray {
		if day == "Monday" {
			utilArray = append(utilArray, 1)
		} else if day == "Tuesday" {
			utilArray = append(utilArray, 2)
		} else if day == "Wednesday" {
			utilArray = append(utilArray, 3)
		} else if day == "Thursday" {
			utilArray = append(utilArray, 4)
		} else if day == "Friday" {
			utilArray = append(utilArray, 5)
		} else if day == "Saturday" {
			utilArray = append(utilArray, 6)
		} else if day == "Sunday" {
			utilArray = append(utilArray, 0)
		}
	}
	return utilArray
}

/**
 * Calculates the cost of the advertisement over the number of days advertisement has been requested
 * @param  {[type]} purchaseInfo   datastructures.Purchase    [description]
 * @param  {[type]} companyInfo    datastructures.CompanyInfo [description]
 * @param  {[type]} purchaseReport datastructures.Report      [description]
 * @return {[type]}                [description]
 */
func calcDaysCost(purchaseInfo datastructures.Purchase, companyInfo datastructures.CompanyInfo, purchaseReport datastructures.Report) float64 {
	var price float64 = 0.0
	myDays := days(purchaseInfo.ShowDates)
	dayToCost := map[int]float64{0: companyInfo.SundayRate, 1: companyInfo.WeekdayRate, 2: companyInfo.WeekdayRate,
		3: companyInfo.WeekdayRate, 4: companyInfo.WeekdayRate, 5: companyInfo.WeekdayRate, 6: companyInfo.SaturdayRate}

	utilArrayStart := strings.Split(purchaseReport.PurchaseInfo.StartDate, " ")
	utilArrayEnd := strings.Split(purchaseReport.PurchaseInfo.EndDate, " ")

	year, _ := strconv.ParseFloat(utilArrayStart[2], 10)
	year2, _ := strconv.ParseFloat(utilArrayEnd[2], 10)

	day, _ := strconv.ParseFloat(utilArrayStart[0], 10)
	day2, _ := strconv.ParseFloat(utilArrayEnd[0], 10)

	s := time.Date(int(year), monthToTime(utilArrayStart[1]), int(day), 23, 0, 0, 0, time.UTC)
	e := time.Date(int(year2), monthToTime(utilArrayEnd[1]), int(day2), 23, 0, 0, 0, time.UTC)

	modifiedE := e.AddDate(0, 0, 1)

	for !s.Equal(modifiedE) {
		if validDay(myDays, int(s.Weekday())) {
			price += dayToCost[int(s.Weekday())]
		}
		s = s.AddDate(0, 0, 1)
	}
	return price
}

/**
 * Checks if day is a valid day. I.E. a day that the customer has request ad to be shown
 * @param  {[type]} days []int         [description]
 * @param  {[type]} day  int           [description]
 * @return {[type]}      [description]
 */
func validDay(days []int, day int) bool {
	i := sort.Search(len(days), func(i int) bool { return days[i] >= day })
	if i < len(days) && days[i] == day {
		return true
	} else {
		return false
	}
}
