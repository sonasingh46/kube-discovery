package main

import (
	"flag"
	discoveryClient "k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"path/filepath"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/tools/clientcmd"

	//rest "k8s.io/client-go/rest"

	"fmt"
)

type Discovery struct {
	Client discoveryClient.DiscoveryInterface
}

func NewDiscovery()(*Discovery,error){
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
		panic(err.Error())
	}
	dc,err:=discoveryClient.NewDiscoveryClientForConfig(config)
	if err!=nil{
		return nil,err
	}
	return &Discovery{dc},nil
}

func main() {
	fmt.Println("Hello Kube World\n")
	dc,err:=NewDiscovery()
	if err!=nil{
		fmt.Println("Could not create discovery client:",err)
	}
	dc.Client.ServerGroups()
}
