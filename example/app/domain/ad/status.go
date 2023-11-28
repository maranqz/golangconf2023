package ad

import "slices"

type Status int

func NewStatus(in int) (Status, error) {
	res := Status(in)

	if _, ok := statusMP[res]; !ok {
		return 0, ErrStatus
	}

	return res, nil
}

var statusMP = map[Status]struct{}{
	Draft:     {},
	Published: {},
	Archived:  {},
	Ban:       {},
}

const (
	Draft Status = iota + 1
	Published
	Archived
	Ban
	Deleted
)

func (s Status) Is(compare Status, ss ...Status) bool {
	if s == compare {
		return true
	}

	for _, compare := range ss {
		if s == compare {
			return true
		}
	}

	return false
}

// Can проверяет в какой статус можно перейти из текущего
func (s Status) Can(to Status) bool {
	canBe := s.CanBe()

	return len(canBe) > 0 && slices.Contains(canBe, to)
}

func (s Status) CanBe() []Status {
	switch s {
	case Draft:
		return []Status{
			Published,
		}
	case Published:
		return []Status{
			Archived,
			Ban,
		}
	case Archived:
		return []Status{
			Published,
		}
	case Ban:
		return nil
	case Deleted:
		return []Status{
			Draft,
		}
	}

	return nil
}
