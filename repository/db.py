import pymongo
import datetime

client = pymongo.MongoClient('localhost', 27017)
db = client.price_monitor
produts_collection = db.products

def save_product(product_details):
    product_details['created_at'] = datetime.datetime.utcnow()
    return produts_collection.insert_one(product_details)
