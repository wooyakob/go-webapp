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

	connectionString := requiredEnv("CONN_STRING")
	bucketName := requiredEnv("SERVER_BUCKET")
	username := requiredEnv("SERVER_USERNAME")
	password := requiredEnv("SERVER_PASS")
	certPath := os.Getenv("SERVER_ROOT_CERT")

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

	if !strings.Contains(connectionString, "://") {
		connectionString = "couchbase://" + connectionString
	}
	if strings.HasPrefix(connectionString, "couchbases://") {
		if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
			log.Fatal(err)
		}
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

	// Get a reference to the default collection, required for older Couchbase server versions
	// col := bucket.DefaultCollection()

	col := bucket.Scope("tenant_agent_00").Collection("users")

	type User struct {
		Name      string   `json:"name"`
		Email     string   `json:"email"`
		Interests []string `json:"interests"`
	}

	// Create and store a Document
	_, err = col.Upsert("u:Jake",
		User{
			Name:      "Jake",
			Email:     "Jake@test-email.com",
			Interests: []string{"Surfing", "Programming"},
		}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get the document back
	getResult, err := col.Get("u:Jake", nil)
	if err != nil {
		log.Fatal(err)
	}

	var inUser User
	err = getResult.Content(&inUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %v\n", inUser)
}