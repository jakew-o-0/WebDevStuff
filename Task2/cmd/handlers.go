package main

import (
	"encoding/json"
	"net/http"
	"webtest/internal/database"
	"webtest/templates"

	"github.com/labstack/echo/v4"
)

func indexPageGet(c echo.Context) error {
	return render(templates.IndexPage(false), c)
}



func loginPageGet(c echo.Context) error {
	return render(templates.LoginPage(false), c)
}

func loginPagePost(c echo.Context) error {
	// validate
	// create session
	// redirect to home page
	return c.HTML(http.StatusOK, "<p>Login pressed</p>")
}



func signupPageGet(c echo.Context) error {
	return render(templates.SignupPage(false), c)
}




func workoutsPageGet(c echo.Context) error {
	var json_test interface{}
	test_json := `{"workout1": 10, "workout2": 2}`
	err := json.Unmarshal([]byte(test_json), &json_test)
	if err != nil {
		return err
	}


	var workout_list []*database.Workout
	test_workout := database.Workout{Title: "test", Description: "this is test description", Content: json_test}
	workout_list = append(workout_list, &test_workout)

	test_workout2 := database.Workout{Title: "test2", Description: "this is test description", Content: json_test}
	workout_list =  append(workout_list, &test_workout2)


	return render(templates.WorkoutsPage(false, workout_list), c)
}

func dietsPageGet(c echo.Context) error {
	return nil
}
