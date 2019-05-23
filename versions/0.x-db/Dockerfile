# Trivy v0.0.16
#   with initialized vulnerability database
#
# docker run -d --rm -p 9000:9000 pottava/trivy:0.0.16-db

FROM pottava/trivy:0.0.16
RUN trivy alpine:3.9 >/dev/null
