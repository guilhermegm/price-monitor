import requests
import re

def get_site_content(site_url):
    site_request = requests.get(site_url)
    return site_request.text if site_request.status_code == requests.codes.ok else None

def get_links(site_content, site_url):
    links_found = re.finditer(r'href=[\'"](.*?)[\'"]'.format(site_url), site_content)

    for link_match in links_found:
        link_url = link_match.group(1)
        if site_url in link_url:
            yield link_url
        elif link_url.startswith('?'):
            yield '{}/{}'.format(site_url, link_url)
 
def get_product_details(product_url, id_regex, name_regex, price_regex):
    def get_first_match(text_regex):
        matches_found = re.finditer(text_regex, site_content)
        for name_found in matches_found:
            return name_found.group(1)
            break

    site_content = get_site_content(product_url)

    if site_content:
        product_detail = {}
        product_detail['id'] = get_first_match(id_regex)
        product_detail['name'] = get_first_match(name_regex)
        product_detail['price'] = get_first_match(price_regex)

        if None in product_detail.values():
            return

        try:
            product_detail['price'] = float(product_detail['price'])
        except:
            return

        return product_detail
        

