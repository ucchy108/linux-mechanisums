FROM ubuntu:20.04

ENV TZ Asia/Tokyo

WORKDIR /src

RUN apt-get update
RUN DEBIAN_FRONTEND=nointeractive apt-get install -y binutils build-essential golang sysstat strace man python3-matplotlib python3-pil fonts-takao \
  fio qemu-kvm virt-manager libvirt-clients virtinst jq docker.io containerd libvirt-daemon-system 
  
RUN apt-get update && apt-get install -y sudo

RUN adduser `id -un` libvirt
RUN adduser `id -un` libvirt-qemu
RUN adduser `id -un` kvm

RUN groupadd -g 1000 developer && \
  useradd -g developer -G sudo -m -s /bin/bash docker && \
  echjo 'docker:password' | chpasswd

RUN echo 'Defailts visiblepw' >> /etc/sudoers
RUN echo 'docker ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER docker 
