package packet

import (
	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/pharmer/cloud/lib"
	"github.com/appscode/pharmer/contexts"
	"github.com/appscode/pharmer/extpoints"
)

func init() {
	extpoints.KubeProviders.Register(new(kubeProvider), "packet")
}

type kubeProvider struct {
}

var _ extpoints.KubeProvider = &kubeProvider{}

func (cluster *kubeProvider) Create(ctx *contexts.ClusterContext, req *proto.ClusterCreateRequest) error {
	return (&clusterManager{ctx: ctx}).create(req)
}

func (cluster *kubeProvider) Scale(ctx *contexts.ClusterContext, req *proto.ClusterReconfigureRequest) error {
	return lib.UnsupportedOperation
}

func (cluster *kubeProvider) Delete(ctx *contexts.ClusterContext, req *proto.ClusterDeleteRequest) error {
	return (&clusterManager{ctx: ctx}).delete(req)
}

func (cluster *kubeProvider) SetVersion(ctx *contexts.ClusterContext, req *proto.ClusterReconfigureRequest) error {
	return lib.UnsupportedOperation
}

func (cluster *kubeProvider) UploadStartupConfig(ctx *contexts.ClusterContext) error {
	conn, err := NewConnector(ctx)
	if err != nil {
		return err
	}
	cm := &clusterManager{ctx: ctx, conn: conn}
	return cm.UploadStartupConfig()
}

func (cluster *kubeProvider) GetInstance(ctx *contexts.ClusterContext, md *contexts.InstanceMetadata) (*contexts.KubernetesInstance, error) {
	conn, err := NewConnector(ctx)
	if err != nil {
		return nil, err
	}
	im := &instanceManager{conn: conn}
	return im.GetInstance(md)
}

func (cluster *kubeProvider) MatchInstance(i *contexts.KubernetesInstance, md *contexts.InstanceMetadata) bool {
	return i.InternalIP == md.InternalIP
}