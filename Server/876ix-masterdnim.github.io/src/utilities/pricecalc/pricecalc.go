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
	cost := calcDaysCost(purchaseInfo, companyInfo, purchaseReport) * float64(purchaseInfo.Columns) * float64(purchaseInfo.Rows)
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
	switch m {
	case "January":
		return time.January
	case "February":
		return time.February
	case "March":
		return time.March
	case "April":
		return time.April
	case "May":
		return time.May
	case "June":
		return time.June
	case "July":
		return time.July
	case "August":
		return time.August
	case "September":
		return time.September
	case "October":
		return time.October

	case "November":
		return time.November
	case "December":
		return time.December
	default:
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

	start := time.Date(int(year), monthToTime(utilArrayStart[1]), int(day), 23, 0, 0, 0, time.UTC)
	end := time.Date(int(year2), monthToTime(utilArrayEnd[1]), int(day2), 23, 0, 0, 0, time.UTC)

	endPlusOne := end.AddDate(0, 0, 1)

	for !start.Equal(endPlusOne) {
		if validDay(myDays, int(start.Weekday())) {
			price += dayToCost[int(start.Weekday())]
		}
		start = start.AddDate(0, 0, 1)
	}
	return price
}

/**
 * Checks if day is a valid day. I.E. a day that the customer has requested an ad to be shown
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
