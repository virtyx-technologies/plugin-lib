package plugin_lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddMetric(t *testing.T) {
	a := assert.New(t)
	g := GenericResponse{}

	a.Equal(0, len(g.Metrics))
	metric := Metric{Type: "vrn", Datum: 0, Qualifier: "unqualified"}
	g.AddMetric(metric)
	a.Equal(1, len(g.Metrics))
	a.Equal(metric, g.Metrics[0])
}

func TestAddMetadata(t *testing.T) {
	a := assert.New(t)
	g := GenericResponse{}

	a.Equal(0, len(g.Metadata))
	metadata := Metadata{Type: "vrn", Data: "unqualified"}
	g.AddMetadata(metadata)
	a.Equal(1, len(g.Metadata))
	a.Equal(metadata, g.Metadata[0])
}

func TestReport(t *testing.T) {
	a := assert.New(t)
	g := GenericResponse{}

	metric := Metric{Type: "vrn", Datum: 0, Qualifier: "unqualified"}
	g.AddMetric(metric)
	metadata := Metadata{Type: "vrn", Data: "unqualified"}
	g.AddMetadata(metadata)

	s := g.Report()
	a.NotEmpty(s)
}

