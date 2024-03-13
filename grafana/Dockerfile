ARG GRAFANA_VERSION
FROM grafana/grafana:${GRAFANA_VERSION}

COPY dashboard.yml /etc/grafana/provisioning/dashboards
COPY datasource.yml /etc/grafana/provisioning/datasources
COPY dashboards /var/lib/grafana/dashboards