# [WIP] WebPhoneCatalog
Don't lose your numbers.
<hr>

Objective for this project:

Backend: Go, RabbitMQ, PostgreSQL<br/>
Front: React<br/>
K8S: Deployment, Istio, Kiali<br/>
Observability: Prometheus<br/>
Authentication: KeyCloack<br/>
CI: Github Actions and FluxCD
<hr>

- [x] Mock Backend
- [x] Save values at PostgreSQL
- [x] Create a RabbitMQ
- [x] Create a Front
- [x] Create a Makefile  
- [ ] Create a Backend module for queue consumers
- [ ] Create a Tests
- [ ] Create KeyCloak authorization
- [ ] Integration Front with KeyCloack
- [ ] Create Deployment k8s with Prometheus
- [ ] Install and Config Istio and Kiali




Complete: 45%

<hr>

## How to use

You need Docker and Docker Compose installed!

`$make local`

This command will construct local environment
