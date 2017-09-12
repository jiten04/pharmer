package cloud

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	proto "github.com/appscode/api/ssh/v1beta1"
	"github.com/appscode/pharmer/api"
	"k8s.io/client-go/util/cert"
	kubeadmconst "k8s.io/kubernetes/cmd/kubeadm/app/constants"
)

func CreateCACertificates(ctx context.Context, cluster *api.Cluster) (context.Context, error) {
	Logger(ctx).Infoln("Generating CA certificate for cluster")

	certStore := Store(ctx).Certificates(cluster.Name)

	// -----------------------------------------------
	if cluster.Spec.CACertName == "" {
		cluster.Spec.CACertName = kubeadmconst.CACertAndKeyBaseName

		caKey, err := cert.NewPrivateKey()
		if err != nil {
			return ctx, fmt.Errorf("Failed to generate private key. Reason: %v.", err)
		}
		caCert, err := cert.NewSelfSignedCACert(cert.Config{CommonName: cluster.Spec.CACertName}, caKey)
		if err != nil {
			return ctx, fmt.Errorf("Failed to generate self-signed certificate. Reason: %v.", err)
		}

		ctx = context.WithValue(ctx, paramCACert{}, caCert)
		ctx = context.WithValue(ctx, paramCAKey{}, caKey)
		certStore.Create(cluster.Spec.CACertName, caCert, caKey)
	}

	// -----------------------------------------------
	if cluster.Spec.FrontProxyCACertName == "" {
		cluster.Spec.FrontProxyCACertName = kubeadmconst.FrontProxyCACertAndKeyBaseName
		frontProxyCAKey, err := cert.NewPrivateKey()
		if err != nil {
			return ctx, fmt.Errorf("Failed to generate private key. Reason: %v.", err)
		}
		frontProxyCACert, err := cert.NewSelfSignedCACert(cert.Config{CommonName: cluster.Spec.CACertName}, frontProxyCAKey)
		if err != nil {
			return ctx, fmt.Errorf("Failed to generate self-signed certificate. Reason: %v.", err)
		}

		ctx = context.WithValue(ctx, paramFrontProxyCACert{}, frontProxyCACert)
		ctx = context.WithValue(ctx, paramFrontProxyCAKey{}, frontProxyCAKey)
		certStore.Create(cluster.Spec.FrontProxyCACertName, frontProxyCACert, frontProxyCAKey)
	}

	Logger(ctx).Infoln("CA certificates generated successfully.")
	return ctx, nil
}

func LoadCACertificates(ctx context.Context, cluster *api.Cluster) (context.Context, error) {
	certStore := Store(ctx).Certificates(cluster.Name)

	caCert, caKey, err := certStore.Get(cluster.Spec.CACertName)
	if err != nil {
		return ctx, fmt.Errorf("Failed to get CA certificates. Reason: %v.", err)
	}
	ctx = context.WithValue(ctx, paramCACert{}, caCert)
	ctx = context.WithValue(ctx, paramCAKey{}, caKey)

	frontProxyCACert, frontProxyCAKey, err := certStore.Get(cluster.Spec.FrontProxyCACertName)
	if err != nil {
		return ctx, fmt.Errorf("Failed to get front proxy CA certificates. Reason: %v.", err)
	}
	ctx = context.WithValue(ctx, paramFrontProxyCACert{}, frontProxyCACert)
	ctx = context.WithValue(ctx, paramFrontProxyCAKey{}, frontProxyCAKey)

	return ctx, nil
}

func CreateAdminCertificate(ctx context.Context) (*x509.Certificate, *rsa.PrivateKey, error) {
	cfg := cert.Config{
		CommonName:   "cluster-admin",
		Organization: []string{kubeadmconst.MastersGroup},
		Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	adminKey, err := cert.NewPrivateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate private key. Reason: %v.", err)
	}
	adminCert, err := cert.NewSignedCert(cfg, adminKey, CACert(ctx), CAKey(ctx))
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate server certificate. Reason: %v.", err)
	}
	return adminCert, adminKey, nil
}

func CreateSSHKey(ctx context.Context, cluster *api.Cluster) (context.Context, error) {
	sshKey, err := api.NewSSHKeyPair()
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, paramSSHKey{}, sshKey)
	err = Store(ctx).SSHKeys(cluster.Name).Create(cluster.Status.SSHKeyExternalID, sshKey.PublicKey, sshKey.PrivateKey)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func LoadSSHKey(ctx context.Context, cluster *api.Cluster) (context.Context, error) {
	publicKey, privateKey, err := Store(ctx).SSHKeys(cluster.Name).Get(cluster.Status.SSHKeyExternalID)
	if err != nil {
		return ctx, fmt.Errorf("Failed to get SSH key. Reason: %v.", err)
	}
	protoSSH := &proto.SSHKey{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	ctx = context.WithValue(ctx, paramSSHKey{}, protoSSH)
	return ctx, nil
}
