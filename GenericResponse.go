package plugin_lib

import (
	"bytes"
	"fmt"
	"errors"
)

type (
	// The struct returned to the agent by a plugin,
	// containing all metrics and metadata and the
	// text of any error which may have occurred
	GenericResponse struct {
		Error    string     `json:"error"`
		Metrics  []Metric   `json:"metrics"`
		Metadata []Metadata `json:"metadata"`
	}

	// Encapsulates a metric captured by a plugin
	Metric struct {
		Type      string  `json:"type"`      // The metric type as a Virtyx Resource Name (vrn)
		Datum     float64 `json:"datum"`     // The value captured
		Qualifier string  `json:"qualifier"` // Optional string to qualify the source of the metric
	}

	// Encapsulates metadata (textual information) captured by a plugin
	Metadata struct {
		Type string `json:"type"` // The metadata type as a Virtyx Resource Name (vrn)
		Data string `json:"data"` // The value captured
	}
)

// Add a metric to the response
func (g *GenericResponse) AddMetric(m Metric) {
	g.Metrics = append(g.Metrics, m)
}

// Add metadata to the response
func (g *GenericResponse) AddMetadata(m Metadata) {
	g.Metadata = append(g.Metadata, m)
}

// Find a metric given its VRN (type)
func (g *GenericResponse) FindMetric(vrn string) (*Metric, error) {
	for _, m := range g.Metrics {
		if m.Type == vrn {
			return &m, nil
		}
	}
	return nil, errors.New("not found")
}

// Find a metadata given its VRN (type)
func (g *GenericResponse) FindMetadata(vrn string) (*Metadata, error) {
	for _, m := range g.Metadata {
		if m.Type == vrn {
			return &m, nil
		}
	}
	return nil, errors.New("not found")
}

// Format the response as a human-readable string
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
	return fmt.Sprintf("%s|%s", m.Type, m.Data)
}

