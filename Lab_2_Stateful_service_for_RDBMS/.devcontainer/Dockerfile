FROM mcr.microsoft.com/devcontainers/cpp:ubuntu-22.04

RUN sudo apt-get update && sudo apt-get install -y git python3 pip iputils-ping postgresql-client\
    && sudo apt-get clean

RUN ldconfig

WORKDIR /workspace/social_network

ENTRYPOINT ["/usr/bin/bash"]