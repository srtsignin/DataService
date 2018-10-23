FROM bash

COPY . .
RUN ls

CMD /service.exe