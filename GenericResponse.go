package plugin_lib

import (
	"fmt"
	"bytes"
)

type (
	Metric struct {
		Type      string  `json:"type"`
		Datum     float64 `json:"datum"`
		Qualifier string  `json:"qualifier"`
	}

	Metadata struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}

	GenericResponse struct {
		Error    string     `json:"error"`
		Metrics  []Metric   `json:"metrics"`
		Metadata []Metadata `json:"metadata"`
	}
)

func (g *GenericResponse) AddMetric(m Metric) {
	g.Metrics = append(g.Metrics, m)
}

func (g *GenericResponse) AddMetadata(m Metadata) {
	g.Metadata = append(g.Metadata, m)
}

func (g *GenericResponse) Report() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Error: %s\n", g.Error))
	buffer.WriteString(fmt.Sprintf("Metrics: %d\n", len(g.Metrics)))
	for _, m := range g.Metrics {
		buffer.WriteString(fmt.Sprintf("\t%s %f %s\n", m.Type, m.Datum, m.Qualifier))
	}
	buffer.WriteString(fmt.Sprintf("Metadata: %d\n", len(g.Metadata)))
	for _, m := range g.Metadata {
		buffer.WriteString(fmt.Sprintf("\t%s %s\n", m.Type, m.Data))
	}

	return buffer.String()
}

func (m Metric) String() string {
	return fmt.Sprintf("%s|%f|%s", m.Type, m.Datum, m.Qualifier)
}

func (m Metadata) String() string {
	return fmt.Sprintf("%s|%s|%s", m.Type, m.Data)
}

