%define name validation-service

Name:     %{name}
Version:  %{rev}
Release:  1%{?dist}
Epoch:	  %{epoch}
Vendor:   Verifone
Summary:  Validation Service
License:  Commercial
Source:  validation-service.tar.gz

BuildRequires:  golang >= 1.13.11
AutoReqProv:    no

%description
Validation Service for the Greenbox environment

%prep
%setup -n %{name}

%build
mkdir -p ./_build/src/bitbucket.verifone.com/
ln -s $(pwd) ./_build/src/bitbucket.verifone.com/validation-service

export GOPATH=$(pwd)/_build:%{gopath}
go build -o validation-service -ldflags "-X main.version=%{version} -s -w -extldflags -static" .

%pre
echo 'executing preinstall script'
if [ -e /etc/systemd/system/%{name}.service ]; then
	/etc/systemd/system/%{name}.service stop &> /dev/null
fi

%install
mkdir -p %{buildroot}/etc/dimebox/%{name}
mkdir -p %{buildroot}%{_unitdir}

install -d %{buildroot}%{_bindir}
install -p -m 0755 ./%{name} %{buildroot}%{_bindir}/%{name}

# Create service: validation-service
mkdir -p %{buildroot}/etc/systemd/system
cp ./rpm/%{name}.service %{buildroot}/etc/systemd/system/

# Create rsyslog config file that write syslog messages from validation-service to /var/log/dimebox/validation-service.log
mkdir -p %{buildroot}/etc/rsyslog.d
mkdir -p %{buildroot}/var/log/dimebox/%{name}
cp ./rpm/%{name}.conf %{buildroot}/etc/rsyslog.d/

%post
echo 'Setting permission to directories...'
chmod -R ug+x,o+r /etc/dimebox/%{name}
chmod -R ug+x,o+r /var/log/dimebox/%{name}
chmod -R ug+x,o+r /opt/dimebox/%{name}

echo '%{name} installed. Before start check configuration files.'
systemctl enable /etc/systemd/system/%{name}.service

%clean
rm -rf %{buildroot}

%preun
if [ $1 -eq 0 ]; then
    echo 'Stopping the Service %{name} before uninstalling the rpm..'
	systemctl stop %{name}.service
	echo 'Service %{name} stopped'
    systemctl disable %{name}.service
fi

%files
%defattr(-,root,root,-)
%doc README.md
%config(noreplace) /etc/rsyslog.d/%{name}.conf
%{_bindir}/%{name}
%{_unitdir}/%{name}.service
/var/log/dimebox/%{name}
/etc/dimebox/%{name}