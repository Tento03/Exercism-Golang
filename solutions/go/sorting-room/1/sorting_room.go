package sorting

import (
    "fmt"
    "strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	result:=fmt.Sprintf("This is the number %.1f",f)
    return result
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
    fNb:=float64(nb.Number())
	result:=fmt.Sprintf("This is a box containing the number %.1f",fNb)
    return result
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fnb,ok:=fnb.(FancyNumber)
    if !ok{
        return 0
    }

    num,_ := strconv.Atoi(fnb.Value())
    return num
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fn, ok := fnb.(FancyNumber)
	if !ok {
		return "This is a fancy box containing the number 0.0"
	}

	num, _ := strconv.ParseFloat(fn.Value(), 64)

	return fmt.Sprintf(
		"This is a fancy box containing the number %.1f",
		num,
	)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(input any) string {
	switch v := input.(type) {
	case int:
		return DescribeNumber(float64(v))

	case float64:
		return DescribeNumber(v)

	case NumberBox:
		return DescribeNumberBox(v)

	case FancyNumberBox:
		return DescribeFancyNumberBox(v)

	default:
		return "Return to sender"
	}
}
