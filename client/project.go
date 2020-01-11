package client

import (
	"fmt"
	"time"
)

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Identifier  string    `json:"identifier"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedOn   time.Time `json:"created_on"`
	Parent      *Entity   `json:"parent"`
}

type Parent struct{}

type ProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type ProjectResponse struct {
	Project Project `json:"project"`
}

type Entity struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetProject(id int64) (*Project, error) {
	req, err := c.getRequest(fmt.Sprintf("/projects/%v.json", id), "")
	if err != nil {
		return nil, err
	}

	var response ProjectResponse
	_, err = c.Do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response.Project, nil
}

func (c *Client) GetProjects() ([]Project, error) {
	req, err := c.getRequest("/projects.json", "")
	if err != nil {
		return nil, err
	}

	var response ProjectsResponse
	_, err = c.Do(req, &response)
	if err != nil {
		return nil, err
	}

	return response.Projects, nil
}