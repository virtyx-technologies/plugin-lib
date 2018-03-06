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
		Qualifier string  `json:"qualifier"`
	}

	GenericResponse struct {
		Error    string     `json:"error"`
		Metrics  []Metric   `json:"metrics"`
		Metadata []Metadata `json:"metadata"`
	}
)

func (m Metric) String() string {
	return fmt.Sprintf("%s|%f|%s", m.Type, m.Datum, m.Qualifier)
}

func (m Metadata) String() string {
	return fmt.Sprintf("%s|%s|%s", m.Type, m.Data, m.Qualifier)
}

