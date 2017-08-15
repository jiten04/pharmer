package api

import "time"

type CloudCredential struct {
	Id           int64
	PHID         string    `xorm:"text not null 'phid'"`
	Name         string    `xorm:"text not null 'name'"`
	UserName     string    `xorm:"text not null 'userName'"`
	Provider     string    `xorm:"text not null 'provider'"`
	Data         string    `xorm:"text not null 'data'"`
	DateCreated  time.Time `xorm:"bigint created 'dateCreated'"`
	DateModified time.Time `xorm:"bigint updated 'dateModified'"`
}

func (t CloudCredential) TableName() string {
	return `"ac_cluster"."cloud_credential"`
}

type SSHKey struct {
	Id                 int64
	PHID               string    `xorm:"string  not null 'phid'"`
	Name               string    `xorm:"text  not null 'name'"`
	PublicKey          string    `xorm:"string  not null 'publicKey'"`
	PrivateKey         string    `xorm:"string  not null 'privateKey'"`
	AWSFingerprint     string    `xorm:"string  not null 'awsFingerprint'"`
	OpenSSHFingerprint string    `xorm:"string  not null 'opensshFingerprint'"`
	IsDeleted          int32     `xorm:"smallint not null 'isDeleted'"`
	DateCreated        time.Time `xorm:"bigint created 'dateCreated'"`
	DateModified       time.Time `xorm:"bigint updated 'dateModified'"`
}

func (t SSHKey) TableName() string {
	return `"ac_cluster"."ssh_key"`
}

/*
+---------------------------------+
|                                 |
|  +---------+     +---------+    |     +--------+
|  | PENDING +-----> FAILING +----------> FAILED |
|  +----+----+     +---------+    |     +--------+
|       |                         |
|       |                         |
|  +----v----+                    |
|  |  READY  |                    |
|  +----+----+                    |
|       |                         |
|       |                         |
|  +----v-----+                   |
|  | DELETING |                   |
|  +----+-----+                   |
|       |                         |
+---------------------------------+
        |
        |
   +----v----+
   | DELETED |
   +---------+
*/
const (
	KubernetesStatus_Pending  = "PENDING"
	KubernetesStatus_Failing  = "FAILING"
	KubernetesStatus_Failed   = "FAILED"
	KubernetesStatus_Ready    = "READY"
	KubernetesStatus_Deleting = "DELETING"
	KubernetesStatus_Deleted  = "DELETED"

	// ref: https://github.com/liggitt/kubernetes.github.io/blob/1d14da9c42266801c9ac13cb9608b9f8010dda49/docs/admin/authorization/rbac.md#default-clusterroles-and-clusterrolebindings
	KubernetesAccessModeGroupTeamAdmin    = "kubernetes:team-admin"
	KubernetesAccessModeGroupClusterAdmin = "kubernetes:cluster-admin"
	KubernetesAccessModeGroupAdmin        = "kubernetes:admin"
	KubernetesAccessModeGroupEditor       = "kubernetes:editor"
	KubernetesAccessModeGroupViewer       = "kubernetes:viewer"
	KubernetesAccessModeGroupDenyAccess   = "deny-access"
)

type ClusterOP int64

type ClusterEdge struct {
	Src     string    `xorm:"text not null 'src'"`
	Type    ClusterOP `xorm:"bigint not null 'type'"`
	Dst     string    `xorm:"text not null 'dst'"`
	Created time.Time `xorm:"bigint created 'dateCreated'"`
	Seq     int64     `xorm:"bigint not null 'seq'"`
	DataID  int64     `xorm:"bigint 'dataID'"`
}

func (e ClusterEdge) TableName() string {
	return `"ac_cluster"."edge"`
}

type Kubernetes struct {
	ID                     int64     `xorm:"bigint not null 'id'" mapper:"-"`
	PHID                   string    `xorm:"text not null 'phid'" mapper:"target=Phid"`
	Name                   string    `xorm:"text not null 'name'" mapper:"target=Name"`
	Provider               string    `xorm:"text not null 'provider'" mapper:"target=Provider"`
	ProviderCredentialPHID string    `xorm:"text not null 'providerCredentialPHID'" mapper:"-"`
	Region                 string    `xorm:"text not null 'region'" mapper:"target=Region"`
	Zone                   string    `xorm:"text not null 'zone'" mapper:"target=Zone"`
	OS                     string    `xorm:"text not null 'os'" mapper:"target=Os"`
	CACertPHID             string    `xorm:"text not null 'cACertPHID'" mapper:"-"`
	SSHKeyPHID             string    `xorm:"text not null 'sshKeyPHID'" mapper:"-"`
	ApiServerURL           string    `xorm:"text not null 'apiServerURL'" mapper:"target=ApiServerUrl"`
	Status                 string    `xorm:"text not null 'status'" mapper:"target=Status"`
	StatusCause            string    `xorm:"text 'statusCause'" mapper:"target=StatusCause"`
	BucketName             string    `xorm:"text not null 'bucketName'" mapper:"-"`
	ContextVersion         int64     `xorm:"text not null 'contextVersion'" mapper:"-"`
	StartupConfigToken     string    `xorm:"text not null 'startupConfigToken'" mapper:"-"`
	DoNotDelete            int32     `xorm:"smallint not null default 0 'doNotDelete'"`
	DefaultAccessLevel     string    `xorm:"text not null 'defaultAccessLevel'" mapper:"-"`
	DateCreated            time.Time `xorm:"bigint created 'dateCreated'" mapper:"-"`
	DateModified           time.Time `xorm:"bigint updated 'dateModified'" mapper:"-"`
}

func (t Kubernetes) TableName() string {
	return `"ac_cluster"."kubernetes"`
}

type KubernetesVersion struct {
	ID           int64     `xorm:"bigint not null 'id'"`
	ClusterName  string    `xorm:"text 'clusterName'"`
	Context      string    `xorm:"text 'context'" mapper:"-"`
	DateCreated  time.Time `xorm:"bigint created 'dateCreated'" mapper:"-"`
	DateModified time.Time `xorm:"bigint updated 'dateModified'" mapper:"-"`
}

func (t KubernetesVersion) TableName() string {
	return `"ac_cluster"."kubernetes_version"`
}

const (
	KubernetesInstanceStatus_Ready   = "READY"
	KubernetesInstanceStatus_Deleted = "DELETED"
)

type S_KubernetesInstance struct {
	ID             int64     `xorm:"bigint not null 'id'"`
	PHID           string    `xorm:"text not null 'phid'"`
	KubernetesPHID string    `xorm:"text not null 'kubernetesPHID'"`
	ExternalID     string    `xorm:"text not null 'externalID'"`
	ExternalStatus string    `xorm:"text not null 'externalStatus'"`
	Name           string    `xorm:"text not null 'name'"`
	ExternalIP     string    `xorm:"text 'externalIP'"`
	InternalIP     string    `xorm:"text 'internalIP'"`
	SKU            string    `xorm:"text not null 'sku'"`
	Role           string    `xorm:"text not null 'role'"`
	Status         string    `xorm:"text not null 'status'"`
	DateCreated    time.Time `xorm:"bigint created 'dateCreated'"`
	DateModified   time.Time `xorm:"bigint updated 'dateModified'"`
}

func (t S_KubernetesInstance) TableName() string {
	return `"ac_cluster"."kubernetes_instance"`
}