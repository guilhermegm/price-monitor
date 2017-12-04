def split_links(
    links,
    product_matcher,
    product_url_list,
    product_checked_url_list,
    resting_url_list,
    resting_checked_url_list
):
    for link in links:
        if link not in product_url_list and link not in product_checked_url_list and product_matcher in link:
            product_url_list.append(link)
            product_checked_url_list.append(link)
        elif link not in resting_checked_url_list:
            resting_url_list.append(link)
            resting_checked_url_list.append(link)
