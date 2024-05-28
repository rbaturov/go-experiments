package main

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tunedv1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
)

func createClient() (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", "/home/rbaturov/kubeconfigs/cnfdt7-hostedcluster")
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
	ctx := context.TODO()

	cfg, err := clientcmd.BuildConfigFromFlags("", "/home/rbaturov/kubeconfigs/cnfdt7-hostedcluster")
	if err != nil {
		fmt.Printf("Error building config: %v\n", err)
		return
	}

	cli, err := client.New(cfg, client.Options{})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}
	if err := tunedv1.AddToScheme(cli.Scheme()); err != nil {
		fmt.Printf("Error adding to scheme: %v\n", err)

		return 
	}

	err = getTunedConditions(ctx, cli)
	if err != nil {
		fmt.Printf("Error getting Tuned Profiles: %v\n", err)
		return
	}

}

func getTunedConditions(ctx context.Context, cli client.Client) error {

	tunedProfileList := &tunedv1.ProfileList{}
	if err := cli.List(ctx, tunedProfileList); err != nil {
		klog.Errorf("Cannot list Tuned Profiles: %v", err)
		return err
	}
	for _, profile := range tunedProfileList.Items {
		fmt.Printf("Tuned Profile Name: %s\n", profile.Name)
	}
	return nil
}
