# subway

## Problem Statement
Design a service that can provide a secure network tunnel between an end user
and a private server. The tunnel should only be used for traffic going to the
private backend, all other traffic should be going through the public interface.

Example: Think about IDE license servers. The software expects the license servers
would be available with a IP address and port. Hosting the license servers in a
private network will not work. We want to create a tunnel to be able to provide
access to only that license server. 