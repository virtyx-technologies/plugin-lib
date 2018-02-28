package plugin_lib

import "fmt"

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

func (m Metric) AsString() string {
	return fmt.Sprintf("%s|%f|%s", m.Type, m.Datum, m.Qualifier)
}
