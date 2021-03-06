package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bcokert/bg-mentor/internal/pkg/database"
	"github.com/bcokert/bg-mentor/internal/pkg/requesterror"
	"github.com/bcokert/bg-mentor/pkg/model"
	"github.com/bcokert/bg-mentor/pkg/request"
)

type MemberHandler struct {
	DB *database.Database
}

func (h *MemberHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = request.ShiftURL(req.URL.Path)

	var output []byte
	var err error

	switch head {
	case "":
		switch req.Method {
		case http.MethodGet:
			output, err = h.GET(req)
			if err != nil {
				logger.Errorw("Failed to load member/", "error", err)
				RespondJSON("Member", http.StatusInternalServerError, requesterror.InternalError("Member", "An error occurred loading members", req), res)
				return
			}
			RespondJSON("Member", http.StatusOK, output, res)
			return
		default:
			RespondJSON("Member", http.StatusNotFound, requesterror.MethodNotFound("Member", head, req), res)
			return
		}
	case "me":
		switch req.Method {
		case http.MethodGet:
			output, err = h.GETMe(req)
			if err != nil {
				logger.Errorw("Failed to load member/me", "error", err)
				RespondJSON("Member", http.StatusInternalServerError, requesterror.InternalError("Member", "An error occurred loading that member", req), res)
				return
			}
			RespondJSON("Member", http.StatusOK, output, res)
			return
		default:
			RespondJSON("Member", http.StatusNotFound, requesterror.MethodNotFound("Member", head, req), res)
			return
		}
	default:
		RespondJSON("Member", http.StatusNotFound, requesterror.PathNotFound("Member", head, req), res)
		return
	}
}

func (h *MemberHandler) GETMe(r *http.Request) ([]byte, error) {
	email, err := GetMemberEmailFromContext(r.Context())
	if err != nil {
		return nil, err
	}

	var member *model.Member
	member, err = h.DB.GetMember(email)
	if err != nil {
		return nil, err
	}

	return json.Marshal(member)
}

func (h *MemberHandler) GET(r *http.Request) ([]byte, error) {
	members, err := h.DB.GetMembers()
	if err != nil {
		return nil, err
	}

	return json.Marshal(members)
}
