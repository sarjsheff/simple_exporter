# simple_exporter

Simple config config.yml :
```
listen: ":2112" # listen address
interval: 5     # interval between check
timeout: 5      # connect timeout
urls: 
  - "http://idom.sheff.online"
  - "https://google.com"
```

Run:
```
./simple_exporter -c config.yml
```

| Metric | Description |
| --- | --- |
| ```devmon_alive{url="https://google.com"} 1``` | 1 - url alive, 0 - url down |
| ```devmon_alive{url="https://google.com"} 1``` | 1 - url alive, 0 - url down |
| ```devmon_res_status{code="301",url="https://google.com"} 3``` | Count response by status code |
| ```devmon_tls_notafter{url="https://google.com"} 1.636335421e+09``` | Certificate lifetime not after in unix time |
| ```devmon_tls_notbefore{url="https://google.com"} 1.636335421e+09``` | Certificate lifetime not before in unix time |
| ```devmon_errors{url="https://google.com",error="network"} 1``` | Count errors by url |
