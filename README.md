# Python Flask Backend

This is a Flask application, which is containerized with tests.

- Built in Python version 3.9.7
- Uses [Uses Flask](https://flask.palletsprojects.com/en/2.1.x/) as the spam backend backend.
- Uses [Daisy Framework and NextJS](https://daisyui.com/) for UI development.
- Uses [Golang](https://flask-migrate.readthedocs.io/en/latest/) for handinging requests from ui.

# Structure

## classfier
- The `classifier` contains the data processing and open ai logic

## cmd/api
- The `cmd/api` folder contains the entry point for the golang application
  
## internal and web
- The `internal and web` folders contains logic and handlers for golang application

## UI
- The `ui/text-uI` folder contains the nextJS based UI

## Running the project
- Create a python virtual environment, then install the package dependencies using command `pip install -r requirements.txt`
  
- Install Golang dependecies using command `go mod tidy`
- Install Next dependecies using command `npm i or npm install`

- After installing packages run the servers and UI
- To run the `UI` run the script `npm run dev`. This will start the flask app on port `3000` which can be accessed on the browser through `http://127.0.0.1:3000`.
- After running the start the Golang application `go run cmd/api/*.go`.
- To run python backend from `classifier` folder run `python3 run.py` will start the application on `http://127.0.0.1:5000`
- Remember to update the OPEN_AI key to be able to use it.
