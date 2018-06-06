package main

type Pod struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Annotations struct {
			DNSAlphaKubernetesIoInternal          string `json:"dns.alpha.kubernetes.io/internal"`
			KubernetesIoConfigHash                string `json:"kubernetes.io/config.hash"`
			KubernetesIoConfigMirror              string `json:"kubernetes.io/config.mirror"`
			KubernetesIoConfigSeen                string `json:"kubernetes.io/config.seen"`
			KubernetesIoConfigSource              string `json:"kubernetes.io/config.source"`
			SchedulerAlphaKubernetesIoCriticalPod string `json:"scheduler.alpha.kubernetes.io/critical-pod"`
		} `json:"annotations"`
		CreationTimestamp string `json:"creationTimestamp"`
		Labels            struct {
			K8SApp string `json:"k8s-app"`
		} `json:"labels"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
		UID             string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Containers []struct {
			Command         []string `json:"command"`
			Image           string   `json:"image"`
			ImagePullPolicy string   `json:"imagePullPolicy"`
			LivenessProbe   struct {
				FailureThreshold int `json:"failureThreshold"`
				HTTPGet          struct {
					Host   string `json:"host"`
					Path   string `json:"path"`
					Port   int    `json:"port"`
					Scheme string `json:"scheme"`
				} `json:"httpGet"`
				InitialDelaySeconds int `json:"initialDelaySeconds"`
				PeriodSeconds       int `json:"periodSeconds"`
				SuccessThreshold    int `json:"successThreshold"`
				TimeoutSeconds      int `json:"timeoutSeconds"`
			} `json:"livenessProbe"`
			Name  string `json:"name"`
			Ports []struct {
				ContainerPort int    `json:"containerPort"`
				HostPort      int    `json:"hostPort"`
				Name          string `json:"name"`
				Protocol      string `json:"protocol"`
			} `json:"ports"`
			Resources struct {
				Requests struct {
					CPU string `json:"cpu"`
				} `json:"requests"`
			} `json:"resources"`
			TerminationMessagePath   string `json:"terminationMessagePath"`
			TerminationMessagePolicy string `json:"terminationMessagePolicy"`
			VolumeMounts             []struct {
				MountPath string `json:"mountPath"`
				Name      string `json:"name"`
				ReadOnly  bool   `json:"readOnly,omitempty"`
			} `json:"volumeMounts"`
		} `json:"containers"`
		DNSPolicy       string `json:"dnsPolicy"`
		HostNetwork     bool   `json:"hostNetwork"`
		NodeName        string `json:"nodeName"`
		RestartPolicy   string `json:"restartPolicy"`
		SchedulerName   string `json:"schedulerName"`
		SecurityContext struct {
		} `json:"securityContext"`
		TerminationGracePeriodSeconds int `json:"terminationGracePeriodSeconds"`
		Tolerations                   []struct {
			Key      string `json:"key,omitempty"`
			Operator string `json:"operator"`
			Effect   string `json:"effect,omitempty"`
		} `json:"tolerations"`
		Volumes []struct {
			HostPath struct {
				Path string `json:"path"`
				Type string `json:"type"`
			} `json:"hostPath"`
			Name string `json:"name"`
		} `json:"volumes"`
	} `json:"spec"`
	Status struct {
		Conditions []struct {
			LastProbeTime      interface{} `json:"lastProbeTime"`
			LastTransitionTime string      `json:"lastTransitionTime"`
			Status             string      `json:"status"`
			Type               string      `json:"type"`
		} `json:"conditions"`
		ContainerStatuses []struct {
			ContainerID string `json:"containerID"`
			Image       string `json:"image"`
			ImageID     string `json:"imageID"`
			LastState   struct {
			} `json:"lastState"`
			Name         string `json:"name"`
			Ready        bool   `json:"ready"`
			RestartCount int    `json:"restartCount"`
			State        struct {
				Running struct {
					StartedAt string `json:"startedAt"`
				} `json:"running"`
			} `json:"state"`
		} `json:"containerStatuses"`
		HostIP    string `json:"hostIP"`
		Phase     string `json:"phase"`
		PodIP     string `json:"podIP"`
		QosClass  string `json:"qosClass"`
		StartTime string `json:"startTime"`
	} `json:"status"`
}
