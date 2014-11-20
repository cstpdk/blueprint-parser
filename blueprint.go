package main

type Blueprint struct {
	Version        string `json:"_version"`
	Metadata       []Metadata
	Name           string
	Description    string
	ResourceGroups []ResourceGroup
}

type Metadata struct {
	Name  string
	Value string
}
type ResourceGroup struct {
	Name        string
	Description string
	Resources   []Resource
}
type Resource struct {
	Name        string
	Description string
	UriTemplate string
	Model       Payload
	Parameters  []Parameter
	Actions     []Action
}

type Action struct {
	Name        string
	Description string
	Method      string
	Parameters  []Parameter
	Examples    []TransactionExample
}

type Payload struct {
	Name        string
	Reference   *Reference
	Description string
	Headers     []Metadata
	Body        string
	Schema      string
}

type Parameter struct {
	Description string
	Type        string
	Required    bool
	Default     string
	Example     string
	Values      []string
}

type TransactionExample struct {
	Name        string
	Description string
	Requests    []Payload
	Responses   []Payload
}

type Reference struct {
	Id string
}
