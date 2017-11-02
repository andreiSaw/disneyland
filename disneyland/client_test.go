package disneyland

import (
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type DisneylandTestsConfig struct {
	ClientCert  string `yaml:"client_cert"`
	ClientKey   string `yaml:"client_key"`
	CACert      string `yaml:"ca_cert"`
	ConnectTo   string `yaml:"connect_to"`
	DatabaseURI string `yaml:"db_uri"`
}

var TestsConfig *DisneylandTestsConfig

func initTestsConfig() {
	TestsConfig = &DisneylandTestsConfig{}
	configPath := os.Getenv("DISNEYLAND_TESTS_CONFIG")
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	err = yaml.Unmarshal([]byte(content), TestsConfig)
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

}

func getTransportCredentials() (*credentials.TransportCredentials, error) {
	peerCert, err := tls.LoadX509KeyPair(TestsConfig.ClientCert, TestsConfig.ClientKey)
	if err != nil {
		return nil, err
	}

	caCert, err := ioutil.ReadFile(TestsConfig.CACert)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tc := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{peerCert},
		RootCAs:      caCertPool,
	})

	return &tc, nil
}

func checkJobsEqual(a *Job, b *Job) bool {
	return (a.Project == b.Project) &&
		(a.Id == b.Id) &&
		(a.Status == b.Status) &&
		(a.Metadata == b.Metadata) &&
		(a.Kind == b.Kind) &&
		(a.Output == b.Output) &&
		(a.Input == b.Input)
}

func TestGRPCJobCRUD(t *testing.T) {
	initTestsConfig()
	tc, err := getTransportCredentials()
	if err != nil {
		t.Fail()
	}

	conn, err := grpc.Dial(TestsConfig.ConnectTo, grpc.WithTransportCredentials(*tc))
	checkTestErr(err, t)
	defer conn.Close()
	c := NewDisneylandClient(conn)

	ctx := context.Background()

	createdJob, err := c.CreateJob(ctx, &Job{Status: Job_PENDING})
	checkTestErr(err, t)

	readJob, err := c.GetJob(ctx, &RequestWithId{Id: createdJob.Id})
	checkTestErr(err, t)

	if !checkJobsEqual(createdJob, readJob) {
		t.Fail()
	}

	createdJob.Status = Job_PENDING
	createdJob.Metadata = "meta_test"
	createdJob.Output = "output_test"
	createdJob.Kind = "docker"

	updatedJob, err := c.ModifyJob(ctx, createdJob)
	checkTestErr(err, t)

	if !checkJobsEqual(createdJob, updatedJob) {
		t.Fail()
	}

	createdJob, err = c.CreateJob(ctx, &Job{Kind: "docker"})
	checkTestErr(err, t)

	allJobs, err := c.ListJobs(ctx, &ListJobsRequest{HowMany: 2})
	checkTestErr(err, t)

	if len(allJobs.Jobs) < 1 {
		t.Fail()
	}

	pulledJobs, err := c.PullPendingJobs(ctx, &ListJobsRequest{HowMany: 2})
	checkTestErr(err, t)

	if len(pulledJobs.Jobs) < 1 {
		t.Fail()
	}

	_, err = c.DeleteJob(ctx, &RequestWithId{Id: updatedJob.Id})
	checkTestErr(err, t)

	mJob1 := &Job{Kind: "test1"}
	mJob2 := &Job{Kind: "test2"}
	createdJobs, err := c.CreateMultipeJobs(ctx, &ListOfJobs{[]*Job{mJob1, mJob2}})
	if len(createdJobs.Jobs) != 2 {
		t.Fail()
	}
	checkTestErr(err, t)

	jobs := []*ListJobsRequest{
		{Kind: "test1", HowMany: 1, Project: ""},
		{Kind: "test2", HowMany: 1, Project: ""},
	}
	stream, err := c.BidiJobs(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a job : %v", err)
			}
			log.Printf("Client job.id =%d", in.Id)
		}
	}()
	for _, job := range jobs {
		if err := stream.Send(job); err != nil {
			log.Fatalf("Failed to send a job: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc

}
