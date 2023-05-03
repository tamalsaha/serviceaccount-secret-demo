package main

import (
	"context"
	"fmt"
	"time"

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	if err := useGeneratedClient(); err != nil {
		panic(err)
	}
}

func useGeneratedClient() error {
	fmt.Println("Using Generated client")
	cfg := ctrl.GetConfigOrDie()
	cfg.QPS = 100
	cfg.Burst = 100

	var err error
	kc, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}

	ns := "default"
	name := "vault"
	ctx := context.TODO()

	sa := &core.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}

	sa, err = kc.CoreV1().ServiceAccounts(ns).Create(ctx, sa, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	secret := &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name + "-token",
			Annotations: map[string]string{
				core.ServiceAccountNameKey: name,
			},
		},
		Type: core.SecretTypeServiceAccountToken,
	}
	secret, err = kc.CoreV1().Secrets(ns).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	err = wait.PollImmediateInfinite(30*time.Second, func() (done bool, err error) {
		secret, err = kc.CoreV1().Secrets(ns).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		return len(secret.Data) > 0, nil
	})
	return err
}
