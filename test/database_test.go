package test

import (
	"log"
	"tasko/internal/util"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DockerDBConnectionPort = "5432"
)

func TestConnectDatabase(t *testing.T) {
	pool, resource := InitTestDocker(DockerDBConnectionPort)
	defer func() {
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	}()

	var err error
	util.DBCon, err = gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port="+DockerDBConnectionPort+" sslmode=disable"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the test database: %s", err)
	}

	sqlDB, err := util.DBCon.DB()
	if err != nil {
		t.Fatalf("Failed to get database instance: %s", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Failed to ping the test database: %s", err)
	}
}

func InitTestDocker(exposedPort string) (*dockertest.Pool, *dockertest.Resource) {
	var passwordEnv = "POSTGRES_PASSWORD=postgres"
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"listen_addresses = '*'",
			passwordEnv,
		},
		ExposedPorts: []string{exposedPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			DockerDBConnectionPort + "/tcp": {
				{HostIP: "0.0.0.0", HostPort: exposedPort},
			},
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"} // Important option when container crash and you want to debug container
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := resource.Expire(30); err != nil {
		log.Printf("Docker error: %s", err)
	}

	// retry if container is not ready
	pool.MaxWait = 30 * time.Second
	if err = pool.Retry(func() error {
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return pool, resource
}
