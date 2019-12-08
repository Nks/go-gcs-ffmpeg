package models

import (
	"flag"
	"fmt"
	"os"
)

type Parameters struct {
	Storage        string
	Source         string
	Output         string
	ServiceAccount string
}

func ParseParams() (*Parameters, error) {
	defaultSa := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	serviceAccountPtr := flag.String("sa", defaultSa, "Path to Google JSON Service Account file. Default taken from GOOGLE_APPLICATION_CREDENTIALS environment variable")
	storagePtr := flag.String("storage", "", "GCS Storage")
	sourcePtr := flag.String("source", "", "GCS File Source")
	outputPathPtr := flag.String("output", "", "GCS Output Path")

	flag.Parse()

	if *storagePtr == "" || *sourcePtr == "" || *outputPathPtr == "" || *serviceAccountPtr == "" {
		return nil, fmt.Errorf("missing required parameters")
	}

	return &Parameters{
		ServiceAccount: *serviceAccountPtr,
		Storage:        *storagePtr,
		Source:         *sourcePtr,
		Output:         *outputPathPtr,
	}, nil
}
