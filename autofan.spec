Name:           autofan
Version:        1.0.0
Release:        1%{?dist}
Summary:        Server Fan AutoTune

License:        GPLv3+
URL:            https://www.sonmihpc.com
Source0:        https://www.sonmihpc.com/releases/%{name}-%{version}.tar.gz

%description
Server Fan AutoTune

%prep
%setup -q

%undefine _missing_build_ids_terminate_build
%global debug_package %{nil}

%build

%install
mkdir -p %{buildroot}/%{_sbindir}
mkdir -p %{buildroot}/%{_sysconfdir}/autofan
mkdir -p %{buildroot}/%{_unitdir}

install -m 0744 autofan %{buildroot}/%{_sbindir}/autofan
install -m 0644 config.yaml %{buildroot}/%{_sysconfdir}/autofan/config.yaml
install -m 0644 autofan.service %{buildroot}/%{_unitdir}/autofan.service

%post
systemctl enable autofan --now

%files
%{_sbindir}/autofan
%{_sysconfdir}/autofan/config.yaml
%{_unitdir}/autofan.service

%changelog
* Tue Dec 12 2024 root
- Server Fan AutoTune