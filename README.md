# Auth service 
*sladkoezhkovo*

## Tech stack
- **User store** - PostgreSQL
- **Refresh JWT store** - Redis

## Environment variables

```yaml
# .env.jwt
JWT_REFRESH_SECRET
JWT_ACCESS_SECRET

# .env.pg
POSTGRES_USER       
POSTGRES_PASSWORD   

# .env.redis
REDIS_PASSWORD
```