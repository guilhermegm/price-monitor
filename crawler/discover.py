from utils import http
from utils import helper
from repository import db

product_url_list = []
product_checked_url_list = []
resting_url_list = []
resting_checked_url_list = []

def start(SITE_URL):
    next_url = SITE_URL

    while True:
        if not next_url:
            break

        site_content = http.get_site_content(next_url)
        links_found = http.get_links(site_content, SITE_URL)
        helper.split_links(
            links_found,
            '/produto/',
            product_url_list,
            product_checked_url_list,
            resting_url_list,
            resting_checked_url_list
        )
        next_url = resting_url_list.pop()
        print(next_url)

        while product_url_list:
            product_url = product_url_list.pop()
            print(product_url, len(product_url_list))

            product_details = http.get_product_details(
                product_url,
                'prodid: [\'"](.*?)[\'"]',
                'pname: [\'"](.*?)[\'"]',
                'value: [\'"](.*?)[\'"]'
            )
            print(product_details)
            if product_details:
                print(db.save_product(product_details))
