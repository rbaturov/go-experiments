package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	toolswatch "k8s.io/client-go/tools/watch"
)

const (
	defaultNodeWatchTimeout = 10 * time.Second
	nodeName                = "cnfdf07.telco5gran.eng.rdu2.redhat.com"
)

func createClient() (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", "/home/titzhak/tlv_cluster/kubeconfig")
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func main() {
	c, err := createClient()
	if err != nil {
		log.Fatal(err)
	}

	//UsingWatchDirectly(c)

	factory := informers.NewSharedInformerFactory(c, 0)
	nodeInformer := factory.Core().V1().Nodes().Informer()
	nodeLister := factory.Core().V1().Nodes().Lister()
	factory.Start(wait.NeverStop)

	if !cache.WaitForCacheSync(wait.NeverStop, nodeInformer.HasSynced) {
		log.Fatal("Timed out waiting for caches to sync")
	}

	for {
		n, err := nodeLister.Get(nodeName)
		fmt.Printf("we are here\n")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("node rv: %q\n, node capacity: %v\n, node allocatable: %v\n", n.ResourceVersion, n.Status.Capacity, n.Status.Allocatable)
		time.Sleep(20 * time.Second)
	}
}

func UsingWatchDirectly(c *kubernetes.Clientset) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultNodeWatchTimeout)
	defer cancel()

	fieldSelector := fields.OneTermEqualSelector("metadata.name", nodeName).String()
	lw := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (object runtime.Object, e error) {
			options.FieldSelector = fieldSelector
			return c.CoreV1().Nodes().List(ctx, options)
		},
		WatchFunc: func(options metav1.ListOptions) (i watch.Interface, e error) {
			options.FieldSelector = fieldSelector
			return c.CoreV1().Nodes().Watch(ctx, options)
		},
	}

	condF := func(event watch.Event) (bool, error) {
		switch event.Type {
		case watch.Deleted, watch.Error:
			fmt.Printf("we are here: %v", event.Type)
			return false, nil
		}
		switch t := event.Object.(type) {
		case *corev1.Node:
			return t.Status.Capacity != nil, nil
		}
		return false, nil
	}

	evt, err := toolswatch.UntilWithSync(ctx, lw, &corev1.Node{}, nil, condF)
	if err != nil {
		log.Fatal(err)
	}
	if n, ok := evt.Object.(*corev1.Node); ok {
		fmt.Printf("node rv: %q, node capacity: %v", n.ResourceVersion, n.Status.Capacity)
	}
	log.Fatal("event object not of type node")
}
