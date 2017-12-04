import pymongo
import datetime

client = pymongo.MongoClient('localhost', 27017)
db = client.price_monitor
products_collection = db.products
products_price_collection = db.products_price

def save_product(product_details):
    product_details['updated_at'] = datetime.datetime.utcnow()
    save_product_price(product_details)
    return products_collection.update_one(
        { 'id': product_details['id'] },
        { '$set': product_details },
        upsert = True
    )

def save_product_price(product_details):
    product_price = {
        'product_id': product_details['id'],
        'price': product_details['price'],
        'created_at': datetime.datetime.utcnow()
    }
    return products_price_collection.insert_one(product_price)
