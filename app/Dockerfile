FROM python:3.7
COPY . /app
WORKDIR /app
RUN pip install -r requirements.txt
EXPOSE 5002
ENTRYPOINT ["flask"]
CMD ["run", "--env-file=.env", "--host=0.0.0.0", "--port=5002"]
