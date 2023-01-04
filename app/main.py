from flask import Flask
from dotenv import load_dotenv

app = Flask(__name__)

load_dotenv()
  
@app.route('/')
def hello():
    return "It's day 2 of my 100 days challenge."
    