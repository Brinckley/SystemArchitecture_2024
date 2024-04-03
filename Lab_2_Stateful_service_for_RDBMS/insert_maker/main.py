import json
import random
from random import randrange

import psycopg2
from faker import Faker

fake = Faker()


def write_data(path, number):
    file = open(path, "w")
    file.writelines("[")
    for i in range(number):
        random.seed(20)
        username = str(fake.password()) + str(fake.user_name()) + str(randrange(10000))
        email = username + fake.email()
        s = "{" + (f"\"username\":\"{username}\","
                   f" \"password\" : \"{str(fake.password())}\","
                   f" \"first_name\" : \"{str(fake.first_name())}\", "
                   f"\"last_name\" : \"{str(fake.last_name())}\","
                   f" \"email\" : \"{email}\"") + "},\n"
        file.writelines(s)
    random.seed(20)
    username = str(fake.password()) + str(fake.user_name()) + str(randrange(10000))
    email = username + fake.email()
    s = "{" + (f"\"username\":\"{username}\","
               f" \"password\" : \"{str(fake.password())}\","
               f" \"first_name\" : \"{str(fake.first_name())}\", "
               f"\"last_name\" : \"{str(fake.last_name())}\","
               f" \"email\" : \"{email}\"") + "}]\n"
    file.writelines(s)


def insert_to_db(path):
    with open(path) as file:
        data = json.load(file)

    con = psycopg2.connect("dbname=sndb user=admin password=admin host=postgres")
    cur = con.cursor()
    insert_sql = """INSERT INTO social_network.account (username, password, first_name, last_name, email) VALUES (%s, %s, %s, %s, %s);"""

    for row in data:
        cur.execute(insert_sql, [row["username"], row["password"],
                                 row["first_name"], row["last_name"], row["email"]])
    con.commit()


def main():
    number_env = 7000
    accounts_file_path = 'insert.json'
    write_data(accounts_file_path, number_env)
    print("data written to file")
    insert_to_db(accounts_file_path)


main()
