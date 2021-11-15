package web

import (
	"app/app/users"
	"encoding/json"
	"io"
	"net/http"
)

func (self *WebServer) ListUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var users, err = self.app.ListUsers(r.Context(), users.ListUsersQuery{})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var raw []byte
	if raw, err = json.Marshal(users); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(raw)
}

func (self *WebServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var b, err = io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var data = struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err = json.Unmarshal(b, &data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var dto users.UserDTO
	if dto, err = self.app.CreateUser(r.Context(), users.CreateUserCmd{
		Name:  data.Name,
		Email: data.Email,
	}); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var raw []byte
	if raw, err = json.Marshal(dto); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(raw)
}
