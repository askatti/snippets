"""flask_api.py."""
##############################################################################

from flask import Flask, jsonify, request, abort, make_response
from flask_restful import Resource
import time

class number(Resource):
    """Rest api class."""
    def __init__(self):
        self.num=11;
    def get(self):
          time.sleep(.1)
          self.num=self.num
          print ("default number: {}".format(self.num))
          return self.num




class hello(Resource):
    """Rest api class."""
    def __init__(self):
        self.num=12;

    def get(self):
        while (1):
          time.sleep(.1)
          self.num=self.num%2
          print ("hello: {}".format(self.num))


class howru(Resource):
    """Rest api class."""
    def __init__(self):
        self.num=13;

    def get(self):
        while (1):
          time.sleep(.1)
          self.num=self.num%3
          print ("howru: {}".format(self.num))


