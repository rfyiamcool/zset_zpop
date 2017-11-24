# zset_zpop

use redis lua script to make zset's zpop function !

> zpop == ( zrange key 0  0, zrem value ) in redis !!!

## Usage:

```
func Zpop(rc redis.Conn, key string) (result string, err error)

:param rc,  redigo conn
:param key, zset key name

```

end.
