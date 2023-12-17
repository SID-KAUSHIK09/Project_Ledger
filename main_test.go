package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectAPI(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/project":
			switch r.Method {
			case http.MethodGet:
				// Test GET /project
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`[{"id":1,"name":"TestProject","ptype":"ecommerce","status":"inprocess"}]`))

			case http.MethodPost:
				// Test POST /project
				var project Project
				err := json.NewDecoder(r.Body).Decode(&project)
				assert.NoError(t, err)
				assert.Equal(t, "TestProject", project.Name)
				assert.Equal(t, "ecommerce", project.Ptype)
				assert.Equal(t, "inprocess", project.Status)

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"Project created successfully"}`))

			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		case "/project/1":
			switch r.Method {
			case http.MethodPut:
				// Test PUT /project/{id}
				var project Project
				err := json.NewDecoder(r.Body).Decode(&project)
				assert.NoError(t, err)
				assert.Equal(t, "UpdatedProject", project.Name)
				assert.Equal(t, "logistics", project.Ptype)
				assert.Equal(t, "completed", project.Status)

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"Project updated successfully"}`))

			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		case "/project/2":
			switch r.Method {
			case http.MethodDelete:
				// Test DELETE /project/{id}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"Project deleted successfully"}`))

			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	t.Run("GetProjects", func(t *testing.T) {
		resp, err := http.Get(ts.URL + "/project")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("CreateProject", func(t *testing.T) {
		project := Project{Name: "TestProject", Ptype: "ecommerce", Status: "inprocess"}
		payload, err := json.Marshal(project)
		assert.NoError(t, err)

		resp, err := http.Post(ts.URL+"/project", "application/json", strings.NewReader(string(payload)))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("UpdateProject", func(t *testing.T) {
		project := Project{Name: "UpdatedProject", Ptype: "logistics", Status: "completed"}
		payload, err := json.Marshal(project)
		assert.NoError(t, err)

		req, err := http.NewRequest("PUT", ts.URL+"/project/1", strings.NewReader(string(payload)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("DeleteProject", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", ts.URL+"/project/2", nil)
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
