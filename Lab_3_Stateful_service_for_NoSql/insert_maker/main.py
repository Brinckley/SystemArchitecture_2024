from faker import Faker

fake = Faker()


def write_data(path_posts, path_msgs, number):
    file_posts = open(path_posts, "w")
    file_msgs = open(path_msgs, "w")

    Faker.seed(15)
    for i in range(number):
        s = "{" + (f"\"account_id\":\"{fake.ean(length=13)}\","
                   f" \"content\" : \"{fake.text(max_nb_chars=80)}\"") + "}\n"
        file_posts.write(s)

    Faker.seed(25)
    for i in range(number):
        s = "{" + (f"\"sender_id\":\"{fake.ean(length=13)}\","
                   f" \"receiver_id\":\"{fake.ean(length=13)}\","
                   f" \"content\":\"{fake.text(max_nb_chars=80)}\"") + "}\n"
        file_msgs.write(s)


def main():
    number_env = 80000
    generated_msgs_file_path = 'generated_msgs.json'
    generated_posts_file_path = 'generated_posts.json'
    write_data(generated_posts_file_path, generated_msgs_file_path, number_env)
    print("data written to file")


main()
