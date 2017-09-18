package main

import (
	"fmt"

	"github.com/previousnext/m8s/api/k8s/env"
	pb "github.com/previousnext/m8s/pb"
	context "golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (srv server) SSHGet(ctx context.Context, in *pb.SSHGetRequest) (*pb.SSHGetResponse, error) {
	resp := new(pb.SSHGetResponse)

	if in.Credentials.Token != *cliToken {
		return resp, fmt.Errorf("token is incorrect")
	}

	secret, err := srv.client.Secrets(*cliNamespace).Get(env.SecretSSH, metav1.GetOptions{})
	if err != nil {
		return resp, err
	}

	resp.SSH = &pb.SSH{
		PrivateKey: secret.Data[keyPrivateKey],
		KnownHosts: secret.Data[keyKnownHosts],
	}

	return resp, nil
}