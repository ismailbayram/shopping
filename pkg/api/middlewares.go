package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ismailbayram/shopping/internal/users/services"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func PanicLoggerMiddleware(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logDetailsOfError(ctx, err)
			ctx.Error(err.(error))
			//ctx.AbortWithStatus(http.StatusInternalServerError)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
			fmt.Println("in recover")
		}
	}()

	ctx.Next()
}

func SecurityMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	ctx.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
	ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

func JSONRequiredMiddleware(ctx *gin.Context) {
	if ctx.GetHeader("Content-Type") != "application/json" {
		ctx.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "JSON required.",
		})
		return
	}
	ctx.Next()
}

func AuthenticationMiddleware(us *services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			ctx.Next()
		}()

		token := ctx.GetHeader("Authorization")
		if token == "" {
			return
		}
		user, err := us.GetByToken(token)
		if err != nil {
			return
		}
		ctx.Set("user", user.ID)
	}
}

func ErrorHandlerMiddleware(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		for _, err := range ctx.Errors {
			// Find out what type of error it is
			switch err.Type {
			case gin.ErrorTypePublic:
				// Only output public errors if nothing has been written yet
				if !ctx.Writer.Written() {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
			case gin.ErrorTypeBind:
				if err.Error() == "EOF" {
					ctx.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
						"error": "JSON parse error.",
					})
					return
				}
				errs := err.Err.(validator.ValidationErrors)
				errorList := make(map[string]string)
				for _, err := range errs {
					errorList[toSnakeCase(err.Field())] = validationErrorToText(err)
				}

				// Make sure we maintain the preset response status
				status := http.StatusBadRequest
				if ctx.Writer.Status() != http.StatusOK {
					status = ctx.Writer.Status()
				}
				ctx.JSON(status, errorList)

			default:
				logDetailsOfError(ctx, err)
			}

		}
		if !ctx.Writer.Written() {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
		}
	}
}

func logDetailsOfError(ctx *gin.Context, err any) {
	hostname, e := os.Hostname()
	if e != nil {
		hostname = "unknown"
	}

	data, _ := io.ReadAll(ctx.Copy().Request.Body)
	logrus.WithFields(logrus.Fields{
		"path":       fmt.Sprintf("%s?%s", ctx.Request.URL.Path, ctx.Request.URL.RawQuery),
		"method":     ctx.Request.Method,
		"clientIP":   ctx.ClientIP(),
		"hostname":   hostname,
		"statusCode": ctx.Writer.Status(),
		"referer":    ctx.Request.Referer(),
		"userAgent":  ctx.Request.UserAgent(),
		"params":     ctx.Params,
		"data":       string(data),
	}).Error(err)
}

func upperCaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func splitErrorMessage(src string) string {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return src
	}
	var entries []string
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}

	for index, word := range entries {
		if index == 0 {
			entries[index] = upperCaseFirst(word)
		} else {
			entries[index] = strings.ToLower(word)
		}
	}
	justString := strings.Join(entries, " ")
	return justString
}

func validationErrorToText(e validator.FieldError) string {
	word := splitErrorMessage(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", word)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", word, e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", word, e.Param())
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("%s must be %s characters long", word, e.Param())
	}
	return fmt.Sprintf("%s is not valid", word)
}
