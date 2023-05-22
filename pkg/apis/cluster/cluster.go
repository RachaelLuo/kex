package cluster

type ClusterInfo struct {
	ID         string `json:"id"`   // kube-system uid
	Code       string `json:"code"` // cluster alias
	Kubeconfig []byte `json:"kubeconfig"`
}
