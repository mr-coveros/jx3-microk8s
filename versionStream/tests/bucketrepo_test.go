package main

import (
	"path/filepath"
	"testing"

	"github.com/jenkins-x-plugins/jx-secret/pkg/cmd/populate/templatertesting"
	config "github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestBucketRepoSecrets(t *testing.T) {
	ns := "jx"

	testSecrets := []runtime.Object{
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "jx-boot",
				Namespace: ns,
			},
			Data: map[string][]byte{
				"username": []byte("gitoperatorUsername"),
				"password": []byte("gitoperatorpassword"),
			},
		},

		// some other secrets used for templating the jenkins-maven-settings Secret
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "nexus",
				Namespace: ns,
			},
			Data: map[string][]byte{
				"password": []byte("my-nexus-password"),
			},
		},
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sonatype",
				Namespace: ns,
			},
			Data: map[string][]byte{
				"username": []byte("my-sonatype-username"),
				"password": []byte("my-sonatype-password"),
			},
		},
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "gpg",
				Namespace: ns,
			},
			Data: map[string][]byte{
				"passphrase": []byte("my-secret-gpg-passphrase"),
			},
		},
	}

	runner := templatertesting.Runner{
		TestCases: []templatertesting.TestCase{
			{
				TestName:   "bucketrepo",
				ObjectName: "bucketrepo-config",
				Property:   "config.yaml",
				Format:     "yaml",
				Requirements: &config.RequirementsConfig{
					Repository: "bucketrepo",
					Cluster: config.ClusterConfig{
						Provider:    "microk8s",
						ProjectID:   "myproject",
						ClusterName: "mycluster",
					},
				},
			},
		},
		SchemaFile:  filepath.Join("..", "charts", "jxgh", "bucketrepo", "secret-schema.yaml"),
		Namespace:   ns,
		KubeObjects: testSecrets,
	}
	runner.Run(t)
}
