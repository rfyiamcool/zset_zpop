# zset_zpop

use redis lua script to make zset's zpop function !

> zpop == ( zrange key 0  0, zrem value ) in redis !!!

## Usage:

*input rc,  redigo/redis*

*input key, zset key name*

```
func Zpop(rc redis.Conn, key string) (result string, err error) {
```

end.
