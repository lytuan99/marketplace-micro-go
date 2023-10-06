package response

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}

	return ""
}

func ResolveError(err error) any {
	return gin.H{"error": err.Error()}
}

func ResponseError(ctx *gin.Context, err error, code int, message string) {
	var msg string

	if ErrorCode(err) == UniqueViolation {
		code = http.StatusForbidden
		msg = "Duplicate value"
	}

	if err == sql.ErrNoRows {
		code = http.StatusNotFound
		msg = "Not found"
	}

	if code == 0 {
		code = http.StatusInternalServerError
	}

	if message != "" {
		msg = message
	}

	ctx.JSON(code, gin.H{"error": msg})
}
