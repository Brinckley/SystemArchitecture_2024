from faker import Faker
import random
from random import randrange
import os
import requests
import time

fake = Faker()


def send_dataframe(account_url, post_route, msg_route, number):
    time.sleep(7)
    headers = {'content-type': 'application/json'}
    for i in range(1, number+1):
        if i % 1000 == 0: 
            print(i)
        random.seed(20)
        username =  str(fake.user_name()) + str(fake.first_name()) + str(fake.last_name())
        email = str(username) + str(fake.email())
        data = {"username": str(username), "password": username, "first_name": str(fake.first_name()),
                "last_name": str(fake.last_name()), "email": str(email)}
        #print(data)
        r = requests.post(account_url, json=data, headers=headers)
        if r.status_code != 200: 
            print(f"Error in query with data : {r}")
        #print("Account : " + str(r.status_code))

    #for i in range(1, number):
    #    data = {"account_id":  i, "content":  fake.text(max_nb_chars=20)}
    #    post_url = account_url + '/' + str(i) + post_route
    #    print(data)
    #    r = requests.post(post_url, json=data, headers=headers)
    #    print("Post : " + str(r.status_code))
#
    #for i in range(1, number):
    #    for j in range(1, number):
    #        if i == j:
    #            continue
    #        data = {"sender_id":  i, "receiver_id": j, "content": fake.text(max_nb_chars=20)}
    #        msg_url = account_url + '/' + str(i) + msg_route
    #        print(data)
    #        r = requests.post(msg_url, json=data, headers=headers)
    #        print("MSG : " + str(r.status_code))


def main():
    host_env = os.getenv("HOST")
    port_env = os.getenv("PORT")
    number_env = 6000
    account_route = os.getenv("ROUTE_ACCOUNT")
    message_route = os.getenv("ROUTE_MESSAGES")
    post_route = os.getenv("ROUTE_POSTS")

    account_url = 'http://' + host_env + ':' + port_env + account_route
    print(account_url)
    send_dataframe(account_url, post_route, message_route, number_env)


main()
