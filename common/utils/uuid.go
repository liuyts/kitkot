package utils

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	uid := uuid.NewString()
	return strings.ReplaceAll(uid, "-", "")
}
