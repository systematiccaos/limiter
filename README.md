This small program listens to a given MQTT topic and publishes to another topic to limit the power of a pv-system so that it does not export power to the public grid. This script will adjust the limit-power every 20 seconds - which makes sense, as the inverter needs some seconds to adjust to the limits change.
You can adjust the topic-routes via environment variables (see my example) to make it fit your needs. Was tested with a emporia powermeter and a Hoymiles HM-600.

To run this:
`cp .env-default .env`
Adjust the values in the .env file, then run:
`docker-compose --env-file .env up -d`