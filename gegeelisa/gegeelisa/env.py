import os
from dotenv import load_dotenv

load_dotenv('.env')
env = dict(os.environ)

# change here there is some variables that need type-changing
int_var = [
    'GEGEELISA_REDIS_PORT',
]

for var in int_var:
    env[var] = int(env[var])

env['GEGEELISA_MYSQL_DB_URI'] = 'mysql://{}:{}@{}:{}/{}?charset=utf8mb4'.format(
	env['GEGEELISA_MYSQL_USERNAME'],
	env['GEGEELISA_MYSQL_PASSWORD'],
	env['GEGEELISA_MYSQL_HOST'],
	env['GEGEELISA_MYSQL_PORT'],
	env['GEGEELISA_MYSQL_DB'],
)
