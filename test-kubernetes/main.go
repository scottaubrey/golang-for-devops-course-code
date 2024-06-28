package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	client, err := getClient()
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}
	context := context.TODO()
	deployment, err := deploy(client, context)
	if err != nil {
		log.Fatalf("error creating deployment: %s", err)
	}

	fmt.Println("Changed deployment")

	err = waitForDeployment(client, context, deployment)
	if err != nil {
		log.Fatalf("error watching deployment: %s", err)
	}
}

func getClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func deploy(client *kubernetes.Clientset, context context.Context) (*appsV1.Deployment, error) {
	name := "helloworld"
	deploymentLabels := map[string]string{
		"app": "helloworld",
	}
	deployment := createHelloWorldDeploymentStruct(name, deploymentLabels)

	_, err := client.AppsV1().Deployments("default").Get(context, name, metaV1.GetOptions{})
	if err != nil && k8serrors.IsNotFound(err) {
		createOptions := metaV1.CreateOptions{}
		deploymentResponse, err :=
			client.AppsV1().Deployments("default").Create(context, &deployment, createOptions)
		if err != nil {
			return nil, fmt.Errorf("error creating deployment: %s", err)
		}
		return deploymentResponse, nil
	}

	deploymentResponse, err := client.AppsV1().Deployments("default").Update(context, &deployment, metaV1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("error updating deployment: %s", err)
	}
	return deploymentResponse, nil
}

func waitForDeployment(client *kubernetes.Clientset, context context.Context, deployment *appsV1.Deployment) error {
	watchResponse, err := client.AppsV1().Deployments("default").Watch(context, metaV1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name=%s", deployment.Name)})
	if err != nil {
		return fmt.Errorf("error watching deployments: %s", err)
	}
	for {
		event := <-watchResponse.ResultChan()

		switch event.Type {
		case watch.Deleted:
			return errors.New("deployment deleted")
		}
		deployState := event.Object.(*appsV1.Deployment)

		targetReplicas := deployment.Spec.Replicas
		readyReplicas := deployState.Status.ReadyReplicas
		fmt.Printf("expected replicas: %+v\n", *targetReplicas)
		fmt.Printf("ready replicas: %+v\n", readyReplicas)

		if readyReplicas == *targetReplicas {
			fmt.Println("Depoyment Ready")
			break
		}
	}
	return nil
}

func createHelloWorldDeploymentStruct(name string, labels map[string]string) appsV1.Deployment {
	replicas := int32(3)
	return appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: name,
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "k8s-demo",
							Image: "wardviaene/k8s-demo",
							Ports: []v1.ContainerPort{{Name: "nodejs-port", ContainerPort: 3000}},
						},
					},
				},
			},
		},
	}
}
