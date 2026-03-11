		package main

		import (
		"fmt"
		"log"
		"time"

		"github.com/couchbase/gocb/v2"
		)

		func main() {
			// Update this to your cluster details
			connectionString := "couchbases://cb.ozbazrbqjal-6idq.cloud.couchbase.com"
			username := "GOSDKTEST"
			password := "<<password>>"
			bucketName := "travel-sample"
			scopeName := "<<replace with your scope name>>"
			collectionName := "<<replace with your collection name>>"
		
			// User input ends here
			// Airline struct for sample document
			type Airline struct {
				Type     string `json:"type"`
				Id       string `json:"id"`
				Callsign string `json:"callsign"`
				Iata     string `json:"iata"`
				Icao     string `json:"icao"`
				Name     string `json:"name"`
		}
		
			// Sample airline document
			var couchbaseAirline Airline
			couchbaseAirline.Type = "airline"
			couchbaseAirline.Id = "8091"
			couchbaseAirline.Callsign = "CBS"
			couchbaseAirline.Name = "Couchbase Airways"
		
			// Key will equal: "airline_8091"
			key := couchbaseAirline.Type + "_" + couchbaseAirline.Id
		
			options := gocb.ClusterOptions{
				Authenticator: gocb.PasswordAuthenticator{
					Username: username,
					Password: password,
				},
			}
		
			// Sets a pre-configured profile called "wan-development" to help avoid latency issues
			// when accessing Capella from a different Wide Area Network
			// or Availability Zone (e.g. your laptop).
			if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
				log.Fatal(err)
			}
		
			// Initialize the Connection
			cluster, err := gocb.Connect(connectionString, options)
			if err != nil {
				log.Fatal(err)
			}
		
			// Get a reference to the bucket
			bucket := cluster.Bucket(bucketName)
		
			err = bucket.WaitUntilReady(5*time.Second, nil)
			if err != nil {
				log.Fatal(err)
			}
		
			// Get a reference to the collection
			collection := bucket.Scope(scopeName).Collection(collectionName)
		
			// Simple K-V operation - to create a document with specific ID
			res, err := collection.Insert(key, couchbaseAirline, nil)
			if err != nil {
				log.Fatal("\nError: ", err)
			}
			fmt.Println("\nCreate document success. CAS:", res.Cas())
		
			// Simple K-V operation - to retrieve a document by ID
			result, err := collection.Get(key, nil)
			if err != nil {
				log.Fatal("\nError: ", err)
			}
		
			var fetchedDoc Airline
			err = result.Content(&fetchedDoc)
			if err != nil {
				log.Fatal("\nError: ", err)
			}
			fmt.Printf("\nFetch document success. Result: %+v\n", fetchedDoc)
		
			// Simple K-V operation - to update a document by ID
			couchbaseAirline.Name = "Couchbase Airways!!"
			res, err = collection.Replace(key, couchbaseAirline, nil)
			if err != nil {
				log.Fatal("\nError: ", err)
			}
			fmt.Println("\nUpdate document success. CAS:", res.Cas())
		
			// Simple K-V operation - to delete a document by ID
			res, err = collection.Remove(key, nil)
			if err != nil {
				log.Fatal("\nError: ", err)
			}
			fmt.Println("\nDelete document success. CAS:", res.Cas())
		}
	
	