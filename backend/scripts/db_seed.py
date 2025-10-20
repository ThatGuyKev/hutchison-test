import json 
import sqlite3

seed_data_path = "dogs.json"
output_sql_file = "seed.sql"
db_file = "../app.db"


conn = sqlite3.connect(db_file)
cursor = conn.cursor()


with open(seed_data_path, "r") as file:
    data = json.load(file)

for key, value in data.items():
    if value == []:
        sql  = f"INSERT OR IGNORE INTO dogs (breed) VALUES ('{key}')"
        cursor.execute(sql)
        
    else:
        print("', '".join(value))
        sql  = f"INSERT OR IGNORE INTO dogs (breed, variants) VALUES ('{key}', json_array('{"', '".join(value)}'))"
        cursor.execute(sql)
    


conn.commit()

conn.close()