def split_links(links, product_matcher, product_list, resting_list):
    for link in links:
        if link not in product_list and product_matcher in link:
            product_list.append(link)
        else:
            resting_list.append(link)
