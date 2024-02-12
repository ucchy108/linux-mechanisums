FROM ubuntu:20.04

ENV TZ Asia/Tokyo

COPY . /src
WORKDIR /src

RUN apt-get update
RUN DEBIAN_FRONTEND=nointeractive apt-get install -y binutils build-essential golang sysstat strace man python3-matplotlib python3-pil fonts-takao \
  fio qemu-kvm virt-manager libvirt-clients virtinst jq docker.io containerd libvirt-daemon-system

RUN adduser `id -un` libvirt
RUN adduser `id -un` libvirt-qemu
RUN adduser `id -un` kvm

