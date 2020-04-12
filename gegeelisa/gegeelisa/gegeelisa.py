from flask import Flask
from .env import env
from controller import wikipedia_scrape_blueprint
from model import db


def Gegeelisa():
	gegeelisa = Flask(__name__)
	gegeelisa.config['SQLALCHEMY_DATABASE_URI'] = env['GEGEELISA_MYSQL_DB_URI']
	gegeelisa.config['GEGEELISA_SQLALCHEMY_TRACK_MODIFICATIONS'] = True
	gegeelisa.config['GEGEELISA_MYSQL_DATABASE_CHARSET'] = 'utf8mb4'
	gegeelisa.register_blueprint(wikipedia_scrape_blueprint)
	db.init_app(gegeelisa)

	return gegeelisa
