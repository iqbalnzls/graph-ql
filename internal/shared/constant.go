package shared

type SortingType string

var (
	ASC  = SortingType("ASC")
	DESC = SortingType("DESC")
)

func (s SortingType) ToString() string {
	return string(s)
}
