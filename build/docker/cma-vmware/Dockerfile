# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ARG KUBECTL_VERSION=1.11.3

# Initial preparations
FROM golang:1.11-alpine
RUN apk update && apk add git make curl mercurial

# Let's add the code in
ADD . /go/src/github.com/samsung-cnct/cma-vmware

# Build package in container
WORKDIR /go/src/github.com/samsung-cnct/cma-vmware
RUN make -f build/Makefile container-build

# Now that we have build a static binary, we can stuff it in a bare-bones alpine image
FROM alpine

ARG KUBECTL_VERSION

COPY --from=0 /go/src/github.com/samsung-cnct/cma-vmware/cma-vmware /cma-vmware

ADD https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl /usr/bin/kubectl
RUN chmod +x /usr/bin/kubectl

ENTRYPOINT ["/cma-vmware"]
