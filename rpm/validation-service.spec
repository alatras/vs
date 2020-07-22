Name:     validation-service
Version:  %{rev}
Release:  1%{?dist}
Epoch:	  %{epoch}
Vendor:   Verifone
Summary:  Validation Service
License:  Commercial
Source0:  validation-service.tar.gz

BuildRequires:  golang >= 1.13.11

%description
Validation Service for the Greenbox environment

%prep
%setup -n validation-service

%build
mkdir -p ./_build/src/bitbucket.verifone.com/
ln -s $(pwd) ./_build/src/bitbucket.verifone.com/validation-service

export GOPATH=$(pwd)/_build:%{gopath}
go build -o validation-service -ldflags "-X main.version=%{version} -s -w -extldflags -static" .

%install
mkdir -p %{buildroot}/%{_unitdir}

install -d %{buildroot}%{_bindir}
install -p -m 0755 ./validation-service %{buildroot}%{_bindir}/validation-service

# Create service: validation-service
cat << EOF > %{buildroot}%{_unitdir}/validation-service.service
[Unit]
Description=Validation Service
[Service]
ExecStart=/usr/bin/validation-service
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=validation-service
User=nobody
[Install]
WantedBy=multi-user.target
EOF

# Create rsyslog config file that write syslog messages from validation-service to /var/log/validation-service.log
mkdir -p %{buildroot}/etc/rsyslog.d
cat << EOF > %{buildroot}/etc/rsyslog.d/validation-service.conf
# Create a template which logs only the message, and doesn't add a timestamp
template(name="validation_service_template" type="string" string="%msg%\n")
# Log everything from validation-service to /var/log/validation-service.log
:programname, isequal, "validation-service" /var/log/validation-service.log;validation_service_template
& stop
EOF

%files
%defattr(-,root,root,-)
%doc README.md redoc-static.html
%config(noreplace) /etc/rsyslog.d/validation-service.conf
%{_bindir}/validation-service
%{_unitdir}/validation-service.service
