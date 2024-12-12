VERSION := $(shell git describe --tags --always --match='v*')
version := $(shell echo $(VERSION) |grep -Eo '[0-9]+\.[0-9]+\.[0-9]')

dev:
	dnf install -y golang-bin rpmdevtools
	rpmdev-setuptree

rpm:
	rm -rf ~/rpmbuild/SOURCES/autofan-$(version)
	rm -f ~/rpmbuild/SPECS/autofan.spec
	mkdir -p ~/rpmbuild/SOURCES/autofan-$(version)
	go build -o  ~/rpmbuild/SOURCES/autofan-$(version)
	cp -r autofan.service ~/rpmbuild/SOURCES/autofan-$(version)
	cp -r config.yaml ~/rpmbuild/SOURCES/autofan-$(version)
	cd ~/rpmbuild/SOURCES;tar -cvzf autofan-$(version).tar.gz autofan-$(version)/;rm -rf autofan-$(version)/
	cp -r autofan.spec ~/rpmbuild/SPECS/
	rpmbuild -bb ~/rpmbuild/SPECS/autofan.spec

.PHONY: rpm