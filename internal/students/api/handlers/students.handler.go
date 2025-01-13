package handler

import (
	"encoding/json"
	"net/http"

	student_model "github.com/pRakesh15/student-api/internal/students/model"
	response_utils "github.com/pRakesh15/student-api/internal/students/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//create a variable to store the req.body

	var student student_model.Student

	//decode the req.body and store the data

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		response_utils.RespondWithJSON(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	response_utils.RespondWithJSON(w, http.StatusCreated, student)

}
