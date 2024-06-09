package etc

import (
	"go-inventory/internal/inventory"
	"strings"
	"text/template"
)

func GenerateHosts(hosts []inventory.Host) string {
	// Define the template for /etc/hosts entries
	tmpl := `
{{- range .}}
{{.IP}}	{{.Name}} {{.Alias}} 
{{- end}}`

	// Create a new template and parse the template string
	t, err := template.New("hosts").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Create a buffer to store the rendered output
	var output strings.Builder

	// Execute the template with the hosts data and write the output to the buffer
	err = t.Execute(&output, hosts)
	if err != nil {
		panic(err)
	}

	// Return the rendered output as a string
	return output.String()
}