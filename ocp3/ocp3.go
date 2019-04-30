package ocp3

import (
	"github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"

	configv1 "github.com/openshift/api/legacyconfig/v1"
)

// reference:
// https://docs.openshift.com/container-platform/3.11/install_config/master_node_configuration.html

// TODO: we may want to be OCP3 minor version aware here

// Config represents OCP3 configuration

type Cluster struct {
	Master Master
	Node   Node
}

type Master struct {
	Config configv1.MasterConfig
}

type Node struct {
	Config configv1.NodeConfig
}

func init() {
	configv1.InstallLegacy(scheme.Scheme)
}

// Decode unmarshals OCP3
func (master *Master) Decode(content []byte) {
	var masterConfig configv1.MasterConfig
	serializer := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)

	_, _, err := serializer.Decode(content, nil, &masterConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	master.Config = masterConfig
}
