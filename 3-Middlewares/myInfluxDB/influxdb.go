package myinfluxDB

package main

import (
	"context"

	"github.com/influxdata/influxdb-client-go/v2/domain"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2/internal/examples"
)

func main() {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", "my-token")

	ctx := context.Background()
	// Get Buckets API client
	bucketsAPI := client.BucketsAPI()

	// Get organization that will own new bucket
	org, err := client.OrganizationsAPI().FindOrganizationByName(ctx, "my-org")
	if err != nil {
		panic(err)
	}
	// Create  a bucket with 1 day retention policy
	bucket, err := bucketsAPI.CreateBucketWithName(ctx, org, "bucket-sensors", domain.RetentionRule{EverySeconds: 3600 * 24})
	if err != nil {
		panic(err)
	}

	// Update description of the bucket
	desc := "Bucket for sensor data"
	bucket.Description = &desc
	bucket, err = bucketsAPI.UpdateBucket(ctx, bucket)
	if err != nil {
		panic(err)
	}

	// Close the client
	client.Close()
}
