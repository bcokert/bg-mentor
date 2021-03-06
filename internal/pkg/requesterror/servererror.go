package requesterror

import (
	"fmt"
	"net/http"
	"strings"
)

// InternalError should be used for unexpected internal errors, not client errors.
// It returns an error json response, or an error if something goes wrong
// This is used to handle errors, so we don't want it to fail, thus we manually marshal the json
func InternalError(handler, message string, req *http.Request) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
{
	"message": "%s",
	"handler": "%s",
	"path":    "%s",
	"method":  "%s",
}
	`, message, handler, req.URL.Path, req.Method)))
}

// FileError should be used when a handler cannot process the given path file.
// It returns an error json response, or an error if something goes wrong
// This is used to handle errors, so we don't want it to fail, thus we manually marshal the json
func FileError(handler string, req *http.Request) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
{
	"message": "The '%s' handler could not return the file '%s'",
	"handler": "%s",
	"method":  "%s",
	"tail":    "%s",
}
	`, handler, req.URL.Path, handler, req.Method, req.URL.Path)))
}
