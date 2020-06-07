package main

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	v1 "k8s.io/api/core/v1"
	"log"
)

func main() {
	image := model.Image{
		Registry:   "docker.io/soloio",
		Repository: "example-store-grpc",
		Tag:        "v1",
		PullPolicy: v1.PullIfNotPresent,
	}
	cmd := codegen.Command{
		AppName: "online-shopping",
		Chart: &model.Chart{
			Operators: []model.Operator{
				{
					Name: "books",
					Deployment: model.Deployment{
						Image: image,
					},
					Args: []string{"--service=books"},
				},
				{
					Name: "records",
					Deployment: model.Deployment{
						Image: image,
					},
					Args: []string{"--service=records"},
				},
			},
		},
		ManifestRoot: "install/helm/bookstore",
		Builds: []model.Build{{
			MainFile: "cmd/main.go",
			Push:     true,
			Image:    image,
		}},
		BuildRoot: "build",
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("happy")
}
