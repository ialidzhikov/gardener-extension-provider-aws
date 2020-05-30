#!/bin/bash -e
#
# Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Make it able to work on minikube and nodeless

IP_ROUTE=$(ip route get 1)
IP_ADDRESS=$(echo ${IP_ROUTE#*src} | awk '{print $1}')

VALIDATOR_SERVICE_NAME="gardener-extensions-validator-aws"
VALIDATOR_ENDPOINT_NAME="gardener-extensions-validator-aws"
APISERVER_SERVICE_PORT=443

if kubectl -n garden get service "$VALIDATOR_SERVICE_NAME" &> /dev/null; then
  kubectl -n garden delete service $VALIDATOR_SERVICE_NAME
fi
if kubectl -n garden get endpoints "$VALIDATOR_ENDPOINT_NAME" &> /dev/null; then
  kubectl -n garden delete endpoints $VALIDATOR_ENDPOINT_NAME
fi

cat <<EOF | kubectl apply -f -
kind: Service
apiVersion: v1
metadata:
  name: $VALIDATOR_SERVICE_NAME
  namespace: garden
spec:
  ports:
  - protocol: TCP
    port: 443
    targetPort: 9443
EOF

cat <<EOF | kubectl apply -f -
---
kind: Endpoints
apiVersion: v1
metadata:
  name: $VALIDATOR_ENDPOINT_NAME
  namespace: garden
subsets:
- addresses:
  - ip: ${IP_ADDRESS}
  ports:
  - port: 9443
EOF
