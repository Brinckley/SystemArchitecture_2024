from faker import Faker
import os
import requests

fake = Faker()


def send_dataframe(account_url, post_route, msg_route, number):
    headers = {'content-type': 'application/json'}
    for i in range(1, number):
        data = {"username": str(fake.user_name()), "password": "value", "first_name": str(fake.first_name()),
                "last_name": str(fake.last_name()), "email": str(fake.email())}
        print(data)
        r = requests.post(account_url, json=data, headers=headers)
        print("Account : " + str(r.status_code))

    for i in range(1, number):
        data = {"account_id":  i, "content":  fake.text(max_nb_chars=20)}
        post_url = account_url + '/' + str(i) + post_route
        print(data)
        r = requests.post(post_url, json=data, headers=headers)
        print("Post : " + str(r.status_code))

    for i in range(1, number):
        for j in range(1, number):
            if i == j:
                continue
            data = {"sender_id":  i, "receiver_id": j, "content": fake.text(max_nb_chars=20)}
            msg_url = account_url + '/' + str(i) + msg_route
            print(data)
            r = requests.post(msg_url, json=data, headers=headers)
            print("MSG : " + str(r.status_code))


def main():
    host_env = os.getenv("HOST")
    port_env = os.getenv("PORT")
    number_env = 700
    account_route = os.getenv("ROUTE_ACCOUNT")
    message_route = os.getenv("ROUTE_MESSAGES")
    post_route = os.getenv("ROUTE_POSTS")

    account_url = 'http://' + host_env + ':' + port_env + account_route
    print(account_url)
    send_dataframe(account_url, post_route, message_route, number_env)


main()
