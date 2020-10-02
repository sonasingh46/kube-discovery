package main

import (

	discoveryClient "k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"

	//rest "k8s.io/client-go/rest"

	"fmt"
)

type Discovery struct {
	Client discoveryClient.DiscoveryInterface
}

func NewDiscovery()(*Discovery,error){
	config,err:=rest.InClusterConfig()
	if err!=nil{
		return nil,err
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
