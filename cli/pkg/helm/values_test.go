package helm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var (
	valuesYaml = "./../../tests/testdata/main-values.yaml"
)

func TestChartSpec_GetValuesMap(t *testing.T) {
	type fields struct {
		ValuesYaml  string
		ValuesFiles []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Merge TimescaleDB enable to false",
			fields: fields{
				ValuesYaml: `
timescaledb-single:
  enabled: false`,
				ValuesFiles: []string{valuesYaml},
			},
			want:    "{\"cli\":false,\"grafanaDBJob\":{\"resources\":{}},\"kube-prometheus-stack\":{\"enabled\":true,\"fullnameOverride\":\"tobs-kube-prometheus\",\"grafana\":{\"adminPassword\":\"\",\"enabled\":true,\"envFromSecret\":\"{{ .Release.Name }}-grafana-db\",\"prometheus\":{\"datasource\":{\"enabled\":true,\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201\"}},\"sidecar\":{\"dashboards\":{\"enabled\":true,\"files\":[\"dashboards/k8s-cluster.json\",\"dashboards/k8s-hardware.json\"]},\"datasources\":{\"defaultDatasourceEnabled\":false,\"enabled\":true,\"label\":\"tobs_datasource\",\"labelValue\":\"true\"}},\"timescale\":{\"adminPassSecret\":\"{{ .Release.Name }}-credentials\",\"adminUser\":\"postgres\",\"database\":{\"dbName\":\"postgres\",\"enabled\":true,\"host\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\",\"pass\":\"grafanadb\",\"port\":5432,\"schema\":\"grafanadb\",\"sslMode\":\"require\",\"user\":\"grafanadb\"},\"datasource\":{\"dbName\":\"postgres\",\"enabled\":true,\"host\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\",\"pass\":\"grafana\",\"port\":5432,\"sslMode\":\"require\",\"user\":\"grafana\"}}},\"kube-state-metrics\":{\"prometheusScrape\":false},\"prometheus\":{\"prometheusSpec\":{\"additionalScrapeConfigs\":[{\"job_name\":\"kubernetes-service-endpoints\",\"kubernetes_sd_configs\":[{\"role\":\"endpoints\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scrape\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_service_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_node_name\"],\"target_label\":\"kubernetes_node\"}]},{\"job_name\":\"kubernetes-service-endpoints-slow\",\"kubernetes_sd_configs\":[{\"role\":\"endpoints\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scrape_slow\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_service_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_node_name\"],\"target_label\":\"kubernetes_node\"}],\"scrape_interval\":\"5m\",\"scrape_timeout\":\"30s\"},{\"job_name\":\"kubernetes-services\",\"kubernetes_sd_configs\":[{\"role\":\"service\"}],\"metrics_path\":\"/probe\",\"params\":{\"module\":[\"http_2xx\"]},\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_probe\"]},{\"source_labels\":[\"__address__\"],\"target_label\":\"__param_target\"},{\"replacement\":\"blackbox\",\"target_label\":\"__address__\"},{\"source_labels\":[\"__param_target\"],\"target_label\":\"instance\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"}]},{\"job_name\":\"kubernetes-pods\",\"kubernetes_sd_configs\":[{\"role\":\"pod\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scrape\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_pod_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_pod_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_name\"],\"target_label\":\"kubernetes_pod_name\"},{\"action\":\"drop\",\"regex\":\"Pending|Succeeded|Failed\",\"source_labels\":[\"__meta_kubernetes_pod_phase\"]}]},{\"job_name\":\"kubernetes-pods-slow\",\"kubernetes_sd_configs\":[{\"role\":\"pod\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scrape_slow\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_pod_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_pod_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_name\"],\"target_label\":\"kubernetes_pod_name\"},{\"action\":\"drop\",\"regex\":\"Pending|Succeeded|Failed\",\"source_labels\":[\"__meta_kubernetes_pod_phase\"]}],\"scrape_interval\":\"5m\",\"scrape_timeout\":\"30s\"}],\"evaluationInterval\":\"1m\",\"remoteRead\":[{\"readRecent\":true,\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201/read\"}],\"remoteWrite\":[{\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201/write\"}],\"scrapeInterval\":\"1m\",\"scrapeTimeout\":\"10s\",\"storageSpec\":{\"disableMountSubPath\":true,\"volumeClaimTemplate\":{\"spec\":{\"accessModes\":[\"ReadWriteOnce\"],\"resources\":{\"requests\":{\"storage\":\"8Gi\"}}}}}}},\"prometheus-node-exporter\":{\"service\":{\"annotations\":{\"prometheus.io/scrape\":\"false\"}}},\"prometheusOperator\":{\"configReloaderCpu\":\"10m\",\"configReloaderMemory\":\"20Mi\"}},\"promlens\":{\"defaultPrometheusUrl\":\"http://localhost:9201\",\"enabled\":true,\"image\":\"promlabs/promlens:latest\",\"loadBalancer\":{\"enabled\":false}},\"promscale\":{\"connection\":{\"dbName\":\"postgres\",\"host\":{\"nameTemplate\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\"},\"password\":{\"secretTemplate\":\"{{ .Release.Name }}-credentials\",\"timescaleDBSuperUserKey\":\"PATRONI_SUPERUSER_PASSWORD\"},\"port\":5432,\"user\":\"postgres\"},\"enabled\":true,\"image\":\"timescale/promscale:latest\",\"resources\":{\"requests\":{\"cpu\":\"10m\",\"memory\":\"50Mi\"}},\"service\":{\"loadBalancer\":{\"enabled\":false}}},\"timescaledb-single\":{\"backup\":{\"enabled\":false},\"enabled\":false,\"image\":{\"repository\":\"timescale/timescaledb-ha\",\"tag\":\"pg12-ts2.1-latest\"},\"loadBalancer\":{\"enabled\":false},\"persistentVolumes\":{\"data\":{\"size\":\"150Gi\"},\"wal\":{\"size\":\"20Gi\"}},\"replicaCount\":1},\"timescaledbExternal\":{\"db_uri\":\"\",\"enabled\":false}}",
			wantErr: false,
		},
		{
			name: "Merge Promscale resource values",
			fields: fields{
				ValuesYaml: `
promscale:
  resources:
    requests:
      cpu: 50m
      memory: 500Mi`,
				ValuesFiles: []string{valuesYaml},
			},
			want:    "{\"cli\":false,\"grafanaDBJob\":{\"resources\":{}},\"kube-prometheus-stack\":{\"enabled\":true,\"fullnameOverride\":\"tobs-kube-prometheus\",\"grafana\":{\"adminPassword\":\"\",\"enabled\":true,\"envFromSecret\":\"{{ .Release.Name }}-grafana-db\",\"prometheus\":{\"datasource\":{\"enabled\":true,\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201\"}},\"sidecar\":{\"dashboards\":{\"enabled\":true,\"files\":[\"dashboards/k8s-cluster.json\",\"dashboards/k8s-hardware.json\"]},\"datasources\":{\"defaultDatasourceEnabled\":false,\"enabled\":true,\"label\":\"tobs_datasource\",\"labelValue\":\"true\"}},\"timescale\":{\"adminPassSecret\":\"{{ .Release.Name }}-credentials\",\"adminUser\":\"postgres\",\"database\":{\"dbName\":\"postgres\",\"enabled\":true,\"host\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\",\"pass\":\"grafanadb\",\"port\":5432,\"schema\":\"grafanadb\",\"sslMode\":\"require\",\"user\":\"grafanadb\"},\"datasource\":{\"dbName\":\"postgres\",\"enabled\":true,\"host\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\",\"pass\":\"grafana\",\"port\":5432,\"sslMode\":\"require\",\"user\":\"grafana\"}}},\"kube-state-metrics\":{\"prometheusScrape\":false},\"prometheus\":{\"prometheusSpec\":{\"additionalScrapeConfigs\":[{\"job_name\":\"kubernetes-service-endpoints\",\"kubernetes_sd_configs\":[{\"role\":\"endpoints\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scrape\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_service_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_node_name\"],\"target_label\":\"kubernetes_node\"}]},{\"job_name\":\"kubernetes-service-endpoints-slow\",\"kubernetes_sd_configs\":[{\"role\":\"endpoints\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scrape_slow\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_service_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_node_name\"],\"target_label\":\"kubernetes_node\"}],\"scrape_interval\":\"5m\",\"scrape_timeout\":\"30s\"},{\"job_name\":\"kubernetes-services\",\"kubernetes_sd_configs\":[{\"role\":\"service\"}],\"metrics_path\":\"/probe\",\"params\":{\"module\":[\"http_2xx\"]},\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_probe\"]},{\"source_labels\":[\"__address__\"],\"target_label\":\"__param_target\"},{\"replacement\":\"blackbox\",\"target_label\":\"__address__\"},{\"source_labels\":[\"__param_target\"],\"target_label\":\"instance\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"}]},{\"job_name\":\"kubernetes-pods\",\"kubernetes_sd_configs\":[{\"role\":\"pod\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scrape\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_pod_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_pod_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_name\"],\"target_label\":\"kubernetes_pod_name\"},{\"action\":\"drop\",\"regex\":\"Pending|Succeeded|Failed\",\"source_labels\":[\"__meta_kubernetes_pod_phase\"]}]},{\"job_name\":\"kubernetes-pods-slow\",\"kubernetes_sd_configs\":[{\"role\":\"pod\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scrape_slow\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_pod_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_pod_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_name\"],\"target_label\":\"kubernetes_pod_name\"},{\"action\":\"drop\",\"regex\":\"Pending|Succeeded|Failed\",\"source_labels\":[\"__meta_kubernetes_pod_phase\"]}],\"scrape_interval\":\"5m\",\"scrape_timeout\":\"30s\"}],\"evaluationInterval\":\"1m\",\"remoteRead\":[{\"readRecent\":true,\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201/read\"}],\"remoteWrite\":[{\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201/write\"}],\"scrapeInterval\":\"1m\",\"scrapeTimeout\":\"10s\",\"storageSpec\":{\"disableMountSubPath\":true,\"volumeClaimTemplate\":{\"spec\":{\"accessModes\":[\"ReadWriteOnce\"],\"resources\":{\"requests\":{\"storage\":\"8Gi\"}}}}}}},\"prometheus-node-exporter\":{\"service\":{\"annotations\":{\"prometheus.io/scrape\":\"false\"}}},\"prometheusOperator\":{\"configReloaderCpu\":\"10m\",\"configReloaderMemory\":\"20Mi\"}},\"promlens\":{\"defaultPrometheusUrl\":\"http://localhost:9201\",\"enabled\":true,\"image\":\"promlabs/promlens:latest\",\"loadBalancer\":{\"enabled\":false}},\"promscale\":{\"connection\":{\"dbName\":\"postgres\",\"host\":{\"nameTemplate\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\"},\"password\":{\"secretTemplate\":\"{{ .Release.Name }}-credentials\",\"timescaleDBSuperUserKey\":\"PATRONI_SUPERUSER_PASSWORD\"},\"port\":5432,\"user\":\"postgres\"},\"enabled\":true,\"image\":\"timescale/promscale:latest\",\"resources\":{\"requests\":{\"cpu\":\"50m\",\"memory\":\"500Mi\"}},\"service\":{\"loadBalancer\":{\"enabled\":false}}},\"timescaledb-single\":{\"backup\":{\"enabled\":false},\"enabled\":true,\"image\":{\"repository\":\"timescale/timescaledb-ha\",\"tag\":\"pg12-ts2.1-latest\"},\"loadBalancer\":{\"enabled\":false},\"persistentVolumes\":{\"data\":{\"size\":\"150Gi\"},\"wal\":{\"size\":\"20Gi\"}},\"replicaCount\":1},\"timescaledbExternal\":{\"db_uri\":\"\",\"enabled\":false}}",
			wantErr: false,
		},
		{
			name: "Merging a invalid field",
			fields: fields{
				ValuesYaml:  "",
				ValuesFiles: []string{valuesYaml},
			},
			wantErr: false,
			want:    "{\"cli\":false,\"grafanaDBJob\":{\"resources\":{}},\"kube-prometheus-stack\":{\"enabled\":true,\"fullnameOverride\":\"tobs-kube-prometheus\",\"grafana\":{\"adminPassword\":\"\",\"enabled\":true,\"envFromSecret\":\"{{ .Release.Name }}-grafana-db\",\"prometheus\":{\"datasource\":{\"enabled\":true,\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201\"}},\"sidecar\":{\"dashboards\":{\"enabled\":true,\"files\":[\"dashboards/k8s-cluster.json\",\"dashboards/k8s-hardware.json\"]},\"datasources\":{\"defaultDatasourceEnabled\":false,\"enabled\":true,\"label\":\"tobs_datasource\",\"labelValue\":\"true\"}},\"timescale\":{\"adminPassSecret\":\"{{ .Release.Name }}-credentials\",\"adminUser\":\"postgres\",\"database\":{\"dbName\":\"postgres\",\"enabled\":true,\"host\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\",\"pass\":\"grafanadb\",\"port\":5432,\"schema\":\"grafanadb\",\"sslMode\":\"require\",\"user\":\"grafanadb\"},\"datasource\":{\"dbName\":\"postgres\",\"enabled\":true,\"host\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\",\"pass\":\"grafana\",\"port\":5432,\"sslMode\":\"require\",\"user\":\"grafana\"}}},\"kube-state-metrics\":{\"prometheusScrape\":false},\"prometheus\":{\"prometheusSpec\":{\"additionalScrapeConfigs\":[{\"job_name\":\"kubernetes-service-endpoints\",\"kubernetes_sd_configs\":[{\"role\":\"endpoints\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scrape\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_service_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_node_name\"],\"target_label\":\"kubernetes_node\"}]},{\"job_name\":\"kubernetes-service-endpoints-slow\",\"kubernetes_sd_configs\":[{\"role\":\"endpoints\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scrape_slow\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_service_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_node_name\"],\"target_label\":\"kubernetes_node\"}],\"scrape_interval\":\"5m\",\"scrape_timeout\":\"30s\"},{\"job_name\":\"kubernetes-services\",\"kubernetes_sd_configs\":[{\"role\":\"service\"}],\"metrics_path\":\"/probe\",\"params\":{\"module\":[\"http_2xx\"]},\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_service_annotation_prometheus_io_probe\"]},{\"source_labels\":[\"__address__\"],\"target_label\":\"__param_target\"},{\"replacement\":\"blackbox\",\"target_label\":\"__address__\"},{\"source_labels\":[\"__param_target\"],\"target_label\":\"instance\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_service_label_(.+)\"},{\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"source_labels\":[\"__meta_kubernetes_service_name\"],\"target_label\":\"kubernetes_name\"}]},{\"job_name\":\"kubernetes-pods\",\"kubernetes_sd_configs\":[{\"role\":\"pod\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scrape\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_pod_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_pod_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_name\"],\"target_label\":\"kubernetes_pod_name\"},{\"action\":\"drop\",\"regex\":\"Pending|Succeeded|Failed\",\"source_labels\":[\"__meta_kubernetes_pod_phase\"]}]},{\"job_name\":\"kubernetes-pods-slow\",\"kubernetes_sd_configs\":[{\"role\":\"pod\"}],\"relabel_configs\":[{\"action\":\"keep\",\"regex\":true,\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scrape_slow\"]},{\"action\":\"replace\",\"regex\":\"(https?)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_scheme\"],\"target_label\":\"__scheme__\"},{\"action\":\"replace\",\"regex\":\"(.+)\",\"source_labels\":[\"__meta_kubernetes_pod_annotation_prometheus_io_path\"],\"target_label\":\"__metrics_path__\"},{\"action\":\"replace\",\"regex\":\"([^:]+)(?::\\\\d+)?;(\\\\d+)\",\"replacement\":\"$1:$2\",\"source_labels\":[\"__address__\",\"__meta_kubernetes_pod_annotation_prometheus_io_port\"],\"target_label\":\"__address__\"},{\"action\":\"labelmap\",\"regex\":\"__meta_kubernetes_pod_label_(.+)\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_namespace\"],\"target_label\":\"kubernetes_namespace\"},{\"action\":\"replace\",\"source_labels\":[\"__meta_kubernetes_pod_name\"],\"target_label\":\"kubernetes_pod_name\"},{\"action\":\"drop\",\"regex\":\"Pending|Succeeded|Failed\",\"source_labels\":[\"__meta_kubernetes_pod_phase\"]}],\"scrape_interval\":\"5m\",\"scrape_timeout\":\"30s\"}],\"evaluationInterval\":\"1m\",\"remoteRead\":[{\"readRecent\":true,\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201/read\"}],\"remoteWrite\":[{\"url\":\"http://{{ .Release.Name }}-promscale-connector.{{ .Release.Namespace }}.svc.cluster.local:9201/write\"}],\"scrapeInterval\":\"1m\",\"scrapeTimeout\":\"10s\",\"storageSpec\":{\"disableMountSubPath\":true,\"volumeClaimTemplate\":{\"spec\":{\"accessModes\":[\"ReadWriteOnce\"],\"resources\":{\"requests\":{\"storage\":\"8Gi\"}}}}}}},\"prometheus-node-exporter\":{\"service\":{\"annotations\":{\"prometheus.io/scrape\":\"false\"}}},\"prometheusOperator\":{\"configReloaderCpu\":\"10m\",\"configReloaderMemory\":\"20Mi\"}},\"promlens\":{\"defaultPrometheusUrl\":\"http://localhost:9201\",\"enabled\":true,\"image\":\"promlabs/promlens:latest\",\"loadBalancer\":{\"enabled\":false}},\"promscale\":{\"connection\":{\"dbName\":\"postgres\",\"host\":{\"nameTemplate\":\"{{ .Release.Name }}.{{ .Release.Namespace }}.svc.cluster.local\"},\"password\":{\"secretTemplate\":\"{{ .Release.Name }}-credentials\",\"timescaleDBSuperUserKey\":\"PATRONI_SUPERUSER_PASSWORD\"},\"port\":5432,\"user\":\"postgres\"},\"enabled\":true,\"image\":\"timescale/promscale:latest\",\"resources\":{\"requests\":{\"cpu\":\"10m\",\"memory\":\"50Mi\"}},\"service\":{\"loadBalancer\":{\"enabled\":false}}},\"timescaledb-single\":{\"backup\":{\"enabled\":false},\"enabled\":true,\"image\":{\"repository\":\"timescale/timescaledb-ha\",\"tag\":\"pg12-ts2.1-latest\"},\"loadBalancer\":{\"enabled\":false},\"persistentVolumes\":{\"data\":{\"size\":\"150Gi\"},\"wal\":{\"size\":\"20Gi\"}},\"replicaCount\":1},\"timescaledbExternal\":{\"db_uri\":\"\",\"enabled\":false}}",
		},
		{
			name: "Providing a invalid values file path",
			fields: fields{
				ValuesYaml:  "",
				ValuesFiles: []string{"abc"},
			},
			wantErr: true,
			want:    "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := &ChartSpec{
				ValuesYaml:  tt.fields.ValuesYaml,
				ValuesFiles: tt.fields.ValuesFiles,
			}
			got, err := spec.GetValuesMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValuesMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			jsonData, err := json.Marshal(got)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(string(jsonData))
			if !reflect.DeepEqual(string(jsonData), tt.want) {
				t.Errorf("GetValuesMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
