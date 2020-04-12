from gegeelisa.env import env

from bs4 import BeautifulSoup
beautiful_soup = BeautifulSoup

import redis as redis_lib
redis = redis_lib.Redis(
    host=env['GEGEELISA_REDIS_HOST'], 
    port=env['GEGEELISA_REDIS_PORT'], 
    db=0
)
