// ssh_config.go
package ssh

import (
	"bytes"
	"go-inventory/internal/inventory"
	"text/template"
)

const sshConfigTemplate = `
{{- range . }}
Host {{ .Name }}
    HostName {{ .IP }}
    User {{ .Ansible.User }}
    {{- if .Alias }}
Host {{ .Alias }}
    HostName {{ .IP }}
    User {{ .Ansible.User }}
{{- end }}
{{- end }}
`

func Generate(hosts []inventory.Host, windows bool) (string, error) {
    t := template.New("ssh_config")
    t, err := t.Parse(sshConfigTemplate)
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    err = t.Execute(&buf, hosts)
    if err != nil {
        return "", err
    }

    return buf.String(), nil
}