package main

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/joho/godotenv"
)

func requiredEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required env var: %s", key)
	}
	return val
}

func main() {
	_ = godotenv.Load()

	// enable logging
	gocb.SetLogger(gocb.VerboseStdioLogger())

	connectionString := requiredEnv("CAPELLA_CONN_STRING")
	bucketName := requiredEnv("CAPELLA_BUCKET")
	username := requiredEnv("CAPELLA_ACCESS")
	password := requiredEnv("CAPELLA_ACCESS_PASS")
	certPath := os.Getenv("CAPELLA_CERT")

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	}

	if certPath != "" {
		certBytes, err := os.ReadFile(certPath)
		if err != nil {
			log.Fatal(err)
		}
		rootCAs := x509.NewCertPool()
		if ok := rootCAs.AppendCertsFromPEM(certBytes); !ok {
			log.Fatalf("failed to parse CERT PEM file: %s", certPath)
		}
		options.SecurityConfig = gocb.SecurityConfig{TLSRootCAs: rootCAs}
	}

	if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(connectionString, "://") {
		connectionString = "couchbases://" + connectionString
	}

	cluster, err := gocb.Connect(connectionString, options)
	if err != nil {
		log.Fatal(err)
	}

	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Perform a N1QL Query (scoped)
	scope := bucket.Scope("tenant_agent_00")
	queryResult, err := scope.Query(
		"SELECT u.* FROM users AS u WHERE META(u).id = $docId",
		&gocb.QueryOptions{
			Adhoc: true,
			NamedParameters: map[string]interface{}{
				"docId": "u:Jake",
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print each found Row
	for queryResult.Next() {
		var result interface{}
		err := queryResult.Row(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	if err := queryResult.Err(); err != nil {
		log.Fatal(err)
	}
}
