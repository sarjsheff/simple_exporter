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