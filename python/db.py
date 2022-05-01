import pymongo


def get_connection():
    mongo_client = pymongo.MongoClient("mongodb://127.0.0.1:27017/benchmark")
    mongo_db = mongo_client['benchmark']
    mongo_collection = mongo_db['sample']
    return mongo_collection


db_payload = get_connection()


def find_all():
    data = []
    for el in db_payload.find({}, {'_id': 0}):
        data.append(el)
    return data
