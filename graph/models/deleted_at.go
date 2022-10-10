package models

import (
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

func MarshalDeletedAt(n gorm.DeletedAt) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if n.Valid {
			w.Write([]byte(n.Time.String()))
		} else {
			w.Write([]byte("null"))
		}
	})
}

func UnmarshalDeletedAt(v interface{}) (gorm.DeletedAt, error) {
	s, ok := v.(string)
	t, err := time.Parse("2006-01-02 15:04:05.999999999-0700MST", s)
	return gorm.DeletedAt{Time: t, Valid: err == nil && ok}, err
}
