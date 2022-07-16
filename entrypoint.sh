#!/bin/sh
if [ -z "$(ls -A /config)" ]; then
   cp /app/config-example.yml /config/config.yml
fi

/app/nuntius --config=/config/config.yml