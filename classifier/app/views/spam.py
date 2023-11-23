from app import spamClassifier
from app.classifier.spam import clean_data
from flask import jsonify, request
from app import app 

@app.route("/spam", methods=["POST", "GET"])
def check_message():
    data = request.json
    print(data)
    message = data.get("text", '')

    cleaned_message = clean_data(message)

    res = spamClassifier.classify_as_spam(message=cleaned_message)

    return jsonify({
        "result":res,
    })