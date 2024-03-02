# Auth service 
*sladkoezhkovo*

## Tech stack
- **User store** - PostgreSQL
- **Refresh JWT store** - Redis

## Enviroment variables

```yaml
# jwt.env
JWT_REFRESH_SECRET  - secret phrase for refresh tokens
JWT_ACCESS_SECRET   - secret phrase for access tokens

# pg.env 
# postgresql credentials
POSTGRES_USER       
POSTGRES_PASSWORD   

# redis.env
REDIS_PASSWORD
```