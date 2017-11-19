package utils

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

// NamespaceCreate is a wrapper which will only create a namespace.
func NamespaceCreate(client *kubernetes.Clientset, new *v1.Namespace) (*v1.Namespace, error) {
	ns, err := client.Namespaces().Create(new)
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	return ns, nil
}
