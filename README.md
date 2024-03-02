# Auth service 
*sladkoezhkovo*

## Tech stack
- **User store** - PostgreSQL
- **Refresh JWT store** - Redis

## Environment variables

```yaml
# jwt.env
JWT_REFRESH_SECRET
JWT_ACCESS_SECRET

# pg.env 
POSTGRES_USER       
POSTGRES_PASSWORD   

# redis.env
REDIS_PASSWORD
```