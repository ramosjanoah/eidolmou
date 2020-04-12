from flask import Blueprint, jsonify
from helper import *
from service import WikipediaScrapeService

wikipedia_scrape_blueprint = Blueprint('wikipedia_scrape_service', 'wikipedia_scrape_blueprint', url_prefix='/gegeelisa/')

@wikipedia_scrape_blueprint.route('/areyouok', methods=['GET'])
def give_ok():
    return 'OK\n'

@wikipedia_scrape_blueprint.route('/wikipedia-page/<page_identifier>/scrape', methods=['POST'])
def scrape_page(page_identifier):
    return 'OK\n'

@wikipedia_scrape_blueprint.route('/wikipedia-page/<page_identifier>', methods=['GET'])
def get_page(page_identifier):
    return '{} OK\n'.format(page_identifier)
