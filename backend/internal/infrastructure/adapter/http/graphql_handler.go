package http

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/josediaz/go-and-reactjs/backend/internal/application"
	"github.com/josediaz/go-and-reactjs/backend/internal/domain"
)

// BuildTaskGraphQLSchema construye el schema GraphQL para Tasks (misma lógica que REST).
func BuildTaskGraphQLSchema(service *application.TaskService) (graphql.Schema, error) {
	taskStatusEnum := graphql.NewEnum(graphql.EnumConfig{
		Name: "TaskStatus",
		Values: graphql.EnumValueConfigMap{
			"PENDING":    {Value: string(domain.TaskStatusPending)},
			"IN_PROGRESS": {Value: string(domain.TaskStatusInProgress)},
			"DONE":       {Value: string(domain.TaskStatusDone)},
		},
	})

	taskType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Task",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.String},
			"title": &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"status": &graphql.Field{Type: taskStatusEnum},
			"createdAt": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if t, ok := p.Source.(*domain.Task); ok {
						return t.CreatedAt.Format(time.RFC3339), nil
					}
					return nil, nil
				},
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if t, ok := p.Source.(*domain.Task); ok {
						return t.UpdatedAt.Format(time.RFC3339), nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"task": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return service.GetByID(p.Args["id"].(string))
				},
			},
			"tasks": &graphql.Field{
				Type: graphql.NewList(taskType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return service.ListAll()
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createTask": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"title":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"description": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					desc := ""
					if d, ok := p.Args["description"].(string); ok {
						desc = d
					}
					return service.Create(p.Args["title"].(string), desc)
				},
			},
			"updateTaskStatus": &graphql.Field{
				Type: taskType,
				Args: graphql.FieldConfigArgument{
					"id":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"status": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return service.UpdateStatus(p.Args["id"].(string), domain.TaskStatus(p.Args["status"].(string)))
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
