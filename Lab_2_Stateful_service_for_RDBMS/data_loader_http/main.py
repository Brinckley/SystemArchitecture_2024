from faker import Faker
import json
import os
import requests

fake = Faker()


def send_dataframe(url, number):
    for i in range(number):
        data = {'username': fake.user_name(), 'password': 'value', 'first_name': fake.first_name(),
                'last_name': fake.last_name(), 'email': fake.email()}
        json_data = json.dumps(data)
        r = requests.post(url, json=json_data)
        print(r.status_code)


def main():
    host_env = os.getenv("HOST")
    port_env = os.getenv("PORT")
    route_env = os.getenv("ROUTE")
    number_env = os.getenv("NUMBER")

    url = 'http://' + host_env + ':' + port_env + route_env
    send_dataframe(url, number_env)


main()
