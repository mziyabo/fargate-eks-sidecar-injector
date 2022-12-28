{{/*
Expand the name of the chart.
*/}}
{{- define "fargate-sidecar-injector.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "fargate-sidecar-injector.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "fargate-sidecar-injector.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "fargate-sidecar-injector.labels" -}}
helm.sh/chart: {{ include "fargate-sidecar-injector.chart" . }}
{{ include "fargate-sidecar-injector.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "fargate-sidecar-injector.selectorLabels" -}}
app.kubernetes.io/name: {{ include "fargate-sidecar-injector.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "fargate-sidecar-injector.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "fargate-sidecar-injector.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create Certificate
*/}}
{{- define "fargate-sidecar-injector.cert" -}}
{{- $tls := dict -}}
{{ $currTls := lookup "v1" "ConfigMap" "default" (include "fargate-sidecar-injector.fullname" .) }}
{{- if and $currTls $currTls.data }}
{{- $_ := set $tls "ca.crt" (index $currTls.data "ca.crt") }}
{{- $_ := set $tls "prv.key" (index $currTls.data "prv.key") }}
{{- $_ := set $tls "cert.crt" (index $currTls.data "cert.crt") }}
{{- else }}
{{ $ca := genCA "fargate-sidecar-injector" 365 }}
{{ $cert := genSignedCert "fargate-sidecar-injector.default.svc" (list "127.0.0.1") (list "fargate-sidecar-injector" "fargate-sidecar-injector.default" "fargate-sidecar-injector.default.svc") 365 $ca }}
{{- $_ := set $tls "ca.crt" ($ca.Cert) }}
{{- $_ := set $tls "prv.key" ($cert.Key) }}
{{- $_ := set $tls "cert.crt" ($cert.Cert) }}
{{- end }}
{{- end }}
