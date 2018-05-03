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

func TestFindMetric(t *testing.T) {
	a := assert.New(t)
	g := GenericResponse{}

	a.Equal(0, len(g.Metrics))
	metric1 := Metric{Type: "vrn:metric:one", Datum: 0, Qualifier: "unqualified"}
	g.AddMetric(metric1)
	metric2 := Metric{Type: "vrn:metric:two", Datum: 0, Qualifier: "unqualified"}
	g.AddMetric(metric2)
	actual, err := g.FindMetric("vrn:metric:two")
	a.Nil(err)
	a.Equal(metric2, *actual)
	actual, err = g.FindMetric("vrn:metric:three")
	a.NotNil(err)
	a.Nil(actual)
}

func TestFindMetadata(t *testing.T) {
	a := assert.New(t)
	g := GenericResponse{}

	metric1 := Metadata{Type: "vrn:metric:one", Data: "321" }
	g.AddMetadata(metric1)
	metric2 := Metadata{Type: "vrn:metric:two", Data: "xyzzy"}
	g.AddMetadata(metric2)
	actual, err := g.FindMetadata("vrn:metric:two")
	a.Nil(err)
	a.Equal(metric2, *actual)
	actual, err = g.FindMetadata("vrn:metric:three")
	a.NotNil(err)
	a.Nil(actual)
}

