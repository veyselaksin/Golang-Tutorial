package routes

// func TestUserCreate(t *testing.T) {
// 	type wanted struct {
// 		statusCode   int
// 		expectedKeys []string
// 	}

// 	// Create fiber app for testing
// 	// db_conn := database.ConnectDB()
// 	// fmt.Println(db_conn)
// 	app := fiber.New()

// 	app.Post("/api/users", CreateUser)

// 	// Test Struct
// 	tests := []struct {
// 		name        string
// 		description string
// 		endpoint    string
// 		payload     any
// 		want        wanted
// 	}{
// 		{
// 			name:        "Create a fake user",
// 			description: "This test should return 200 status code and created post",
// 			endpoint:    "/api/users",
// 			payload: map[string]string{
// 				"first_name": "Example user",
// 				"last_name":  "Lorem ipsum",
// 			},
// 			want: wanted{
// 				statusCode: 200,
// 				expectedKeys: []string{
// 					"id",
// 					"first_name",
// 					"last_name",
// 				},
// 			},
// 		},
// 		{
// 			name:        "Create a user with blank content",
// 			description: "This test should return 400 status code and error message",
// 			endpoint:    "/",
// 			payload: map[string]any{
// 				"first_name": "Harry",
// 				"last_name":  "",
// 			},
// 			want: wanted{
// 				statusCode: 400,
// 				expectedKeys: []string{
// 					"message",
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			body, err := json.Marshal(tt.payload)
// 			if err != nil {
// 				t.Errorf("Cannot parse body: %v", err)
// 				t.Fail()
// 			}

// 			req := httptest.NewRequest("POST", tt.endpoint, bytes.NewReader(body))
// 			req.Header.Add("Content-Type", "application/json")

// 			// Create request
// 			res, err := app.Test(req)
// 			if err != nil {
// 				t.Errorf("Cannot test Fiber handler: %v", err)
// 				t.Fail()
// 			}

// 			// Assertions
// 			if res.StatusCode != tt.want.statusCode {
// 				t.Errorf("Expected status code %d, got %d", tt.want.statusCode, res.StatusCode)
// 			}

// 			// answer, err := io.ReadAll(res.Body)
// 			// if err != nil {
// 			// 	t.Errorf("Cannot parse body: %v", err)
// 			// }

// 			// var message map[string]any
// 			// err = json.Unmarshal([]byte(answer), &message)
// 			// if err != nil {
// 			// 	t.Errorf("Cannot unmarshal response: %v", err)
// 			// }

// 			// for _, s := range tt.want.expectedKeys {
// 			// 	if _, ok := message[s]; !ok {
// 			// 		t.Errorf("Expected response body to contain key %s but it can not be found", s)
// 			// 	}
// 			// }
// 		})
// 	}
// }
