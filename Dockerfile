FROM bash

RUN ls
COPY . .
RUN ls
RUN ls /service/
RUN ls /bin/

CMD ["/service"]