```toml

# all certificates
[[tls.certificates]]
  certFile = "foobar"
  keyFile = "foobar"

# default certificate
[tls.stores.Store0.defaultCertificate]
  certFile = "foobar"
  keyFile = "foobar"

# certificate for mtls
[tls.options.Options0.clientAuth]
  caFiles = ["foobar", "foobar"]

# certificates for backend
[http.serversTransports.ServersTransport0]
  rootCAs = ["foobar", "foobar"]
  
  [[http.serversTransports.ServersTransport0.certificates]]
    certFile = "foobar"
    keyFile = "foobar"
```

Goal
    Reload certificates at each modification even if the dynamic conf didn't change.
    The certificates could be updated by another app/process without modifying the path/ref

TODO: check how it works as file reference && in kubernetes Secret if available.

