Docker wrk command : `docker run --rm -v ${PWD}/scripts:/scripts --net=host williamyeh/wrk -c1 -t1 -d20s --latency -s /scripts/test.lua http://localhost:8080`  
Exec command from directory `wrk`.

## With cache
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.80ms    2.25ms  35.86ms   88.44%
    Req/Sec   710.92    147.76     1.03k    69.50%
  Latency Distribution
     50%  559.00us
     75%    2.81ms
     90%    4.15ms
     99%    6.89ms
  14160 requests in 20.01s, 3.28MB read
  Non-2xx or 3xx responses: 1321
Requests/sec:    707.57
Transfer/sec:    167.78KB
```

## Without cache
```
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.20ms    1.66ms  27.10ms   88.41%
    Req/Sec   477.92     65.87   636.00     68.50%
  Latency Distribution
     50%    1.50ms
     75%    2.75ms
     90%    3.99ms
     99%    7.79ms
  9518 requests in 20.01s, 2.52MB read
  Non-2xx or 3xx responses: 751
Requests/sec:    475.55
Transfer/sec:    129.14KB
```  
