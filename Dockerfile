FROM bash

COPY . .
RUN ls /bin/

CMD ["/service"]