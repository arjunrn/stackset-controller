package controller

import (
	"testing"

	fake_cs "github.com/zalando-incubator/stackset-controller/pkg/client/clientset/versioned/fake"
	ss_clientset "github.com/zalando-incubator/stackset-controller/pkg/clientset"
	fake_k8s "k8s.io/client-go/kubernetes/fake"
)

func getFakeClient() *StackSetController {
	fStacksetClient := fake_cs.NewSimpleClientset()
	fk8sClient := fake_k8s.NewSimpleClientset()
	fSSClientSet := ss_clientset.NewClientset(fk8sClient, fStacksetClient)

	return NewStackSetController(fSSClientSet, "test-controller", 0)
}
func testIngressReconcile(t *testing.T) {
	controller := getFakeClient()
	sc := StackSetContainer{}
	controller.ReconcileIngress(sc)
}
