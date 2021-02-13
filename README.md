# nasa-sonda


### Build Docker
 - sudo docker image build -t nasa-sonda .
 - sudo docker container run -p 8089:8089 nasa-sonda

### Heroku

- heroku container:push -a nasa-sonda web
- heroku container:release -a nasa-sonda web