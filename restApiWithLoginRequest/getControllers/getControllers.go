package getControllers

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	err error
)

type Labels struct {
	IoCriContainerdKind       string `json:"io.cri-containerd.kind"`
	IoKubernetesContainerName string `json:"io.kubernetes.container.name"`
	IoKubernetesPodName       string `json:"io.kubernetes.pod.name"`
	IoKubernetesPodNamespace  string `json:"io.kubernetes.pod.namespace"`
	IoKubernetesPodUID        string `json:"io.kubernetes.pod.uid"`
	Name                      string `json:"name"`
	NeuvectorImage            string `json:"neuvector.image"`
	NeuvectorRev              string `json:"neuvector.rev"`
	NeuvectorRole             string `json:"neuvector.role"`
	Release                   string `json:"release"`
	Vendor                    string `json:"vendor"`
	Version                   string `json:"version"`
}

type Controllers struct {
	ConnectionState string    `json:"connection_state"`
	CreatedAt       time.Time `json:"created_at"`
	DisconnectedAt  string    `json:"disconnected_at"`
	DisplayName     string    `json:"display_name"`
	Domain          string    `json:"domain"`
	HostId          string    `json:"host_id"`
	HostName        string    `json:"host_name"`
	ID              string    `json:"id"`
	JoinedAt        time.Time `json:"joined_at"`
	Labels          Labels    `json:"labels"`
	Leader          bool      `json:"leader"`
	Name            string    `json:"name"`
	ConnLastError   string    `json:"orch_conn_last_error"`
	ConnStatus      string    `json:"orch_conn_status"`
	StartedAt       time.Time `json:"started_at"`
	Version         string    `json:"version"`
}
type Body struct {
	Controllers []Controllers `json:"controllers"`
}

func GetControllers(respBody []byte) ([]Controllers, error) {
	var body Body

	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, fmt.Errorf("unmarshall error: %v", err)
	}

	return body.Controllers, nil
}
