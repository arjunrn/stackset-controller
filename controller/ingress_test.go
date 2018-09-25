package controller

import (
	"k8s.io/apimachinery/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	zv1 "github.com/zalando-incubator/stackset-controller/pkg/apis/zalando/v1"
	fake_cs "github.com/zalando-incubator/stackset-controller/pkg/client/clientset/versioned/fake"
	ss_clientset "github.com/zalando-incubator/stackset-controller/pkg/clientset"
	fake_k8s "k8s.io/client-go/kubernetes/fake"
)

func getFakeController() *StackSetController {
	fStacksetClient := fake_cs.NewSimpleClientset()
	fk8sClient := fake_k8s.NewSimpleClientset()
	fSSClientSet := ss_clientset.NewClientset(fk8sClient, fStacksetClient)
	return NewStackSetController(fSSClientSet, "test-controller", 0)
}

func TestIngressReconcile(t *testing.T) {
	controller := getFakeController()
	sc := StackSetContainer{
		StackContainers:map[types.UID]*StackContainer{
			"test": &StackContainer{

			},
		},
		StackSet: zv1.StackSet{

			Spec: zv1.StackSetSpec{
				Ingress: &zv1.StackSetIngressSpec{
					Hosts: []string{"my-app.example.org"},
				},
			},
		},
	}
	sc.Ingress = &v1beta1.Ingress{}
	err := controller.ReconcileIngress(sc)
	ingressList, err := controller.client.ExtensionsV1beta1().Ingresses("test").List(v1.ListOptions{})
	assert.NoError(t, err)
	ingressLenght := len(ingressList.Items)
	assert.Equal(t, ingressLenght, 1)
	assert.Equal(t, err, nil)
}
