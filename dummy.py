from locust import *


""" class MyTaskSet(TaskSet):
    @task(1)
    def hello(self):
        pass


class Dummy(HttpUser):
    tasks = [MyTaskSet] """


class MyTaskSet(TaskSet):
    @task(1)
    def hello(self):
        pass


class Dummy(HttpLocust):
    task_set = MyTaskSet

