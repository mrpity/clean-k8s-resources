FROM scratch
EXPOSE 8080
ENTRYPOINT ["/clean-k8s-resources"]
COPY ./bin/ /