//go:build public_field
// +build public_field

package ad

import (
	"time"
)

// AdPF
// В своих рабочих проектах я как раз использую именно
type AdPF struct {
	ID         AdID
	Status     Status
	ArchivedAt *time.Time
}

func (a *AdPF) Archive(now time.Time) error {
	if !a.Status.Can(Archived) {
		return ErrArchive
	}

	a.ArchivedAt = &now

	return nil
}
