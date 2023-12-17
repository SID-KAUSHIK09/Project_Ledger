package main

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Ptype  string `json:"ptype"`
	Status string `json:"status"`
}

// Projects are taken by ZopSmart with their names, type, and status.
// CRUD functionality is provided by API calls.

func main() {
	// initialise gofr object
	app := gofr.New()

	// GET endpoint to get all projects details
	app.GET("/project", func(ctx *gofr.Context) (interface{}, error) {
		var projects []Project
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM projects")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var project Project
			if err := rows.Scan(&project.ID, &project.Name, &project.Ptype, &project.Status); err != nil {
				return nil, err
			}

			projects = append(projects, project)
		}

		return projects, nil
	})

	// POST endpoint to insert a project using JSON payload
	app.POST("/project", func(ctx *gofr.Context) (interface{}, error) {
		var project Project
		if err := ctx.Bind(&project); err != nil {
			return nil, fmt.Errorf("Invalid Input: %v", err)
		}

		// Validate parameters
		if !isValidProjectType(project.Ptype) {
			return nil, fmt.Errorf("Invalid Input: Project type must be in ecommerce, logistics, retail, supplychain, or others")
		}
		if !isValidStatus(project.Status) {
			return nil, fmt.Errorf("Invalid Input: Status must be 'inprocess' or 'completed'")
		}

		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO projects (name, ptype, status) VALUES (?, ?, ?)", project.Name, project.Ptype, project.Status)

		return nil, err
	})

	// Delete endpoint to delete a project by ID
	app.DELETE("/project/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		_, err := ctx.DB().ExecContext(ctx.Context, "DELETE FROM projects WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		return fmt.Sprintf("Project with ID %s deleted successfully", id), nil
	})

	// PUT endpoint to update a project by ID
	app.PUT("/project/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		// Bind the request body to a Project struct
		var project Project
		if err := ctx.Bind(&project); err != nil {
			return nil, fmt.Errorf("Invalid Input: %v", err)
		}

		// Validate parameters
		if !isValidProjectType(project.Ptype) {
			return nil, fmt.Errorf("Invalid Input: Project type must be in ecommerce, logistics, retail, supplychain, or others")
		}
		if !isValidStatus(project.Status) {
			return nil, fmt.Errorf("Invalid Input: Status must be 'inprocess' or 'completed'")
		}

		_, err := ctx.DB().ExecContext(ctx.Context, "UPDATE projects SET name = ?, ptype = ?, status= ? WHERE id = ?", project.Name, project.Ptype, project.Status, id)
		if err != nil {
			return nil, err
		}

		return fmt.Sprintf("Project with ID %s updated successfully", id), nil
	})

	// starting server
	app.Start()
}

// isValidProjectType checks if the project type is valid
func isValidProjectType(ptype string) bool {
	validTypes := map[string]bool{"ecommerce": true, "logistics": true, "retail": true, "supplychain": true, "others": true}
	return validTypes[ptype]
}

// isValidStatus checks if the status is valid
func isValidStatus(status string) bool {
	return status == "inprocess" || status == "completed"
}
