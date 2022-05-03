# SSHProxy

## Problem Statement
Design an auditable SSH proxy service that can enable users to ssh to remote servers
without having a public ip address or exposed public networks. The service should
be reliable, any connection failure in the pipeline should not affect the
user experiences and no data transmission between the host and client would be
affected. We also care about latency, to provide best user experience, we want to
minimize any latency delay.