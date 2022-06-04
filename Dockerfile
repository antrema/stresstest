FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY stresstest /stresstest

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/stresstest"]
