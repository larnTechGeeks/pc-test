import os
from app.classifier.spam import SpamClassifier
from flask import Flask
from dotenv import load_dotenv

load_dotenv()

app = Flask(__name__)

api_key = os.environ.get("API_KEY")
spamClassifier = SpamClassifier(api_key)

from app.views import spam