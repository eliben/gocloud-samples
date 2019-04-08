package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.opencensus.io/trace"
	"gocloud.dev/gcp"
	"gocloud.dev/server"
	"gocloud.dev/server/sdserver"
)

type GlobalMonitoredResource struct {
	projectId string
}

func (g GlobalMonitoredResource) MonitoredResource() (string, map[string]string) {
	return "global", map[string]string{"project_id": g.projectId}
}

func main() {
	ctx := context.Background()
	credentials, err := gcp.DefaultCredentials(ctx)
	fmt.Println(credentials)

	if err != nil {
		log.Fatal(err)
	}
	tokenSource := gcp.CredentialsTokenSource(credentials)
	fmt.Println(tokenSource)

	projectID, err := gcp.DefaultProjectID(credentials)
	fmt.Println(projectID)
	if err != nil {
		log.Fatal(err)
	}
	mr := GlobalMonitoredResource{projectId: "eliben-test1"}
	fmt.Println("mr", mr)

	exporter, _, err := sdserver.NewExporter(projectID, tokenSource, mr)
	fmt.Println("exporter", exporter)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello\n")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})

	options := &server.Options{
		RequestLogger:         sdserver.NewRequestLogger(),
		HealthChecks:          nil,
		TraceExporter:         exporter,
		DefaultSamplingPolicy: trace.AlwaysSample(),
		Driver:                &server.DefaultDriver{},
	}

	s := server.New(options)
	fmt.Println("Server", s)
	s.ListenAndServe("localhost:8080", mux)
}
