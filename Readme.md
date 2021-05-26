

##### Add below to your configuration file
```
rateLimitConfig:
  redis:
    host: 0.0.0.0
    port: 6379
  rateLimit:
    - key: /relative-url-01 #Relative URL which you want to limit
      rate: 100 #Number of request in one unit
      unit: minute #Valid values are minute/second for now
    - key: /relative-url-02
      rate: 10
      unit: second
```

###### Have this line in your Config Go file (for above config)
```
# Import
import rateLimitConfig "github.com/skshukla/sampleRateLimit/config"
..
RateLimitConfig rateLimitConfig.RateLimitConfig `yaml:"rateLimitConfig"`
```


###### Use Below code to invoke the validation and return response if error, otherwise continue processing as usual
```
# Import
sampleRateLimit "github.com/skshukla/sampleRateLimit"
....
err := sampleRateLimit.ValidateRateLimit(&container.AppConfig.RateLimitConfig , r)
if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("Validate Threshold Reached for URL {%s}", r.URL.Path)))
    return
}
```
