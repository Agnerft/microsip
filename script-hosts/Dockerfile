FROM nginx:mainline-alpine3.17-slim

RUN apk update && \
    apk upgrade && \
    apk add bash curl jq nano git
    #apk add --update nodejs npm



COPY ./update-hosts.sh /home/data/

#CMD [ "curl", "-LO", "https://dl.k8s.io/release/v1.27.2/bin/linux/amd64/kubectl" ]
#CMD [ "install", "-o", "root", "-g", "root", "-m", "0755", "kubectl", "/usr/local/bin/kubectl" ]
#CMD [ "git", "clone", "https://github.com/ahmetb/kubectx", "/opt/kubectx" ]
#CMD [ "ln", "-s", "/opt/kubectx/kubectx", "/usr/local/bin/kubectx" ]
#CMD [ "ln", "-s", "/opt/kubectx/kubens", "/usr/local/bin/kubens" ]
#CMD [ "git", "clone", "https://github.com/Agnerft/kube", "~/.kube/" ]

RUN curl -LO https://dl.k8s.io/release/v1.27.2/bin/linux/amd64/kubectl && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    git clone https://github.com/ahmetb/kubectx /opt/kubectx && \
    ln -s /opt/kubectx/kubectx /usr/local/bin/kubectx && ln -s /opt/kubectx/kubens /usr/local/bin/kubens && \
    git clone https://github.com/Agnerft/kube ~/.kube/
   # kubens crm


#CMD [ "kubens", "crm" ]