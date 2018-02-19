# greenwich
Web service to get timezones for given points.

# usage

start server
```
docker run -p 8080:8080 xydinesh/greenwich
```

send a request using curl or wget
```
curl "http://localhost:8080/timezone/39.623593,-105.141777;47.597893,-122.312020;35.953100,-84.035983;36.186219,-86.281322;33.600525,-112.548158;-33.966503,151.264381;-33.984924,151.307560;-32.014273,115.492316;39.974215,126.413179"

[{"lat":"39.623593","lon":"-105.141777","tz":"America/Denver"},{"lat":"47.597893","lon":"-122.312020","tz":"America/Los_Angeles"},{"lat":"35.953100","lon":"-84.035983","tz":"America/New_York"},{"lat":"36.186219","lon":"-86.281322","tz":"America/Chicago"},{"lat":"33.600525","lon":"-112.548158","tz":"America/Phoenix"},{"lat":"-33.966503","lon":"151.264381","tz":"Australia/Sydney"},{"lat":"-33.984924","lon":"151.307560","tz":"Australia/Sydney"},{"lat":"-32.014273","lon":"115.492316","tz":"Australia/Perth"},{"lat":"39.974215","lon":"126.413179","tz":"Asia/Pyongyang"}]
```